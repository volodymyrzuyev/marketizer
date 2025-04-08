package server

import (
	"context"
	"net/http"

	"fmt"
	"log"
	"time"

	"github.com/a-h/templ"
	"github.com/coder/websocket"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volodymyrzuyev/marketizer/cmd/web"
)

func (s *Server) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("marketizer")
		if err != nil {
			return c.Redirect(302, "/login")
		}
		if cookie.Value == "" {
			return c.Redirect(302, "/login")
		}
		if cookie.MaxAge < 0 {
			return c.Redirect(302, "/login")
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method matches the one used when creating the token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.Redirect(302, "/login")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Redirect(302, "/login")
		}

		c.Set("user", claims["user"])

		return next(c)
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/", s.home, s.AuthMiddleware)
	e.GET("/login", echo.WrapHandler(templ.Handler(web.LoginPage(nil))))
	e.GET("/register", echo.WrapHandler(templ.Handler(web.Register(nil))))

	e.POST("/login", s.loginHandler)
	e.POST("/register", s.register)

	e.GET("/websocket", s.websocketHandler)

	return e
}

func (s *Server) home(c echo.Context) error {
	items, err := s.db.GetItems()
	if err != nil {
		return templ.Handler(web.InternalError()).Component.Render(context.TODO(), c.Response().Writer)
	}

	return templ.Handler(web.Items(items)).Component.Render(context.TODO(), c.Response().Writer)
}

func (s *Server) loginHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	usr, err := s.db.GetUser(email)
	if err != nil {
		return templ.Handler(web.LoginPage(web.ShouldRegister())).Component.Render(context.TODO(), c.Response().Writer)
	}
	if usr.Password.String != password {
		return templ.Handler(web.LoginPage(web.InvalidPassword())).Component.Render(context.TODO(), c.Response().Writer)
	}

	setCookie(c, usr.Email)

	return c.Redirect(302, "/")
}

func (s *Server) register(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	name := c.FormValue("name")

	_, err := s.db.GetUser(email)
	if err == nil {
		return templ.Handler(web.Register(web.EmailExists())).Component.Render(context.TODO(), c.Response().Writer)
	}

	err = s.db.AddUser(email, password, name)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}

	setCookie(c, email)

	return c.Redirect(302, "/")
}

func (s *Server) websocketHandler(c echo.Context) error {
	w := c.Response().Writer
	r := c.Request()
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

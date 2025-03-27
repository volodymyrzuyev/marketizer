package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func setCookie(c echo.Context, email string) {
	fmt.Println(string(jwtKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(2 * time.Hour),
	})

	tokenStr, _ := token.SignedString(jwtKey)

	cookie := new(http.Cookie)
	cookie.Name = "marketizer"
	cookie.Value = tokenStr
	cookie.Path = "/"

	// Set MaxAge to 1 hour
	cookie.MaxAge = 3600

	// Set Expires to 1 hour from now

	// Optionally, you can set Secure and HttpOnly flags
	cookie.Secure = false
	cookie.HttpOnly = true

	// Set the cookie in the response
	c.SetCookie(cookie)
}

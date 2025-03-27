package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/volodymyrzuyev/marketizer/internal/database"
	"github.com/volodymyrzuyev/marketizer/internal/steam"
	mApi "github.com/volodymyrzuyev/steamCommunityMarket"
)

type Server struct {
	port int

	db database.Service
}

func (s *Server) runSteam() {
	requestInterval := time.Duration(10 * time.Second)
	api := mApi.NewApiController(requestInterval)
	parser := steam.NewParser(s.db.AddItems)

	for {
		params := mApi.MarketRecentParams{
			Country:  mApi.UnitedStates,
			Language: mApi.English,
			Currency: mApi.USD,
		}

		resp, err := api.MarketRecent(params)
		if err != nil {
			fmt.Println(err)
			time.Sleep(12 * time.Second)
			continue
		}

		err = parser.RunParsers(resp)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(12 * time.Second)
	}
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go NewServer.runSteam()

	return server
}

package main

import (
	"log"
	"os"

	"github.com/cameron-wood/go-api-starter/internal/server"
)

func main() {
	log.Println("Starting API")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Println("Variable SERVER_PORT not set, using default")
		port = "80"
	}

	c := &server.Config{
		Port: port,
	}

	s := server.New(c)
	s.SetupRoutes()
	s.Start()
}

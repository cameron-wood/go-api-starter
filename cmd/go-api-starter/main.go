package main

import (
	"log"

	"github.com/cameron-wood/go-api-starter/internal/server"
)

func main() {
	log.Println("Starting API")

	s := server.New("80")
	s.SetupRoutes()
	s.Start()
}

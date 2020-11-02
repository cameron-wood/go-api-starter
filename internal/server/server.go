package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server holds web server components
type Server struct {
	port string
	r    *chi.Mux
}

// New returns a new web server
func New() *Server {
	return &Server{}
}

// Start runs the web server
func (s *Server) Start() {
	log.Printf("Starting HTTP server on port %s", s.port)
	err := http.ListenAndServe(":"+s.port, s.r)
	if err != nil {
		log.Fatal(err)
	}
}

// SetupRoutes configures middleware, routes, and handlers
func (s *Server) SetupRoutes() {
	s.r = chi.NewRouter()

	// register middleware
	s.r.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(time.Minute),
	)

	s.r.Get("/health", healthCheck())

	log.Println("Successfully registered routes")
}

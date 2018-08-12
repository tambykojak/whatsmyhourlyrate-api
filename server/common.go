package server

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/tambykojak/whatsmyhourlyrate-api/server/actions"
)

// Server TODO
type Server struct {
	Port   int
	logger *zap.SugaredLogger
}

// Initialize TODO
func (s *Server) Initialize() {
	s.initializeLogger()
	s.setupRoutes()
	s.startListening()
}

func (s *Server) initializeLogger() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal("The logger could not be setup.")
	}

	logger.Info("Logger has successfully been setup.")
	s.logger = logger.Sugar()
}

func (s *Server) setupRoutes() {
	type route struct {
		path    string
		handler http.HandlerFunc
	}

	routes := []route{
		{path: "/health_check", handler: actions.HandleHealthCheck},
		{path: "/hourly_rate", handler: actions.HandleGetHourlyRate}}

	for _, route := range routes {
		http.HandleFunc(route.path, s.middleware(route.handler))
	}
}

func (s *Server) middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func (s *Server) startListening() {
	port := fmt.Sprintf(":%d", s.Port)
	s.logger.Infof("Server starting on port %d.", s.Port)
	http.ListenAndServe(port, nil)
}

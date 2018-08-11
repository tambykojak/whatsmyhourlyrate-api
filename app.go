package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type app struct {
	port   int
	logger *zap.SugaredLogger
}

func (a *app) initialize() {
	a.initializeLogger()
	a.setupRoutes()
	a.startListening()
}

func (a *app) initializeLogger() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal("The logger could not be setup.")
	}

	logger.Info("Logger has successfully been setup.")
	a.logger = logger.Sugar()
}

func (a *app) setupRoutes() {
	type route struct {
		path    string
		handler http.HandlerFunc
	}

	routes := []route{
		{path: "/health_check", handler: a.handleHealthCheck},
		{path: "/hourly_rate", handler: a.handleGetHourlyRate}}

	for _, route := range routes {
		http.HandleFunc(route.path, a.middleware(route.handler))
	}
}

func (a *app) middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func (a *app) startListening() {
	port := fmt.Sprintf(":%d", a.port)
	a.logger.Infof("Server starting on port %s", port)
	http.ListenAndServe(port, nil)
}

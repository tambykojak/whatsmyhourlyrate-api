package app

import (
	"fmt"
	"log"
	"net/http"
	"whatsmyhourlyrate/actions"

	"go.uber.org/zap"
)

// App TODO
type App struct {
	Port   int
	logger *zap.SugaredLogger
}

// Initialize TODO
func (a *App) Initialize() {
	a.initializeLogger()
	a.setupRoutes()
	a.startListening()
}

func (a *App) initializeLogger() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal("The logger could not be setup.")
	}

	logger.Info("Logger has successfully been setup.")
	a.logger = logger.Sugar()
}

func (a *App) setupRoutes() {
	type route struct {
		path    string
		handler http.HandlerFunc
	}

	routes := []route{
		{path: "/health_check", handler: actions.HandleHealthCheck},
		{path: "/hourly_rate", handler: actions.HandleGetHourlyRate}}

	for _, route := range routes {
		http.HandleFunc(route.path, a.middleware(route.handler))
	}
}

func (a *App) middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func (a *App) startListening() {
	port := fmt.Sprintf(":%d", a.Port)
	a.logger.Infof("Server starting on port %d.", a.Port)
	http.ListenAndServe(port, nil)
}

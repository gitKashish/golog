package server

import (
	"net/http"

	"github.com/gitKashish/golog/internal/handlers/api"
	"github.com/gitKashish/golog/internal/handlers/web"
	"github.com/gitKashish/golog/internal/server/middleware"
	"github.com/gitKashish/golog/pkg/logger"
)

// SetupRoutes sets up all routes for the server
func (s *Server) SetupRoutes() {
	// Apply global middleware
	s.router.Handle("/", middleware.Chain(
		http.NotFoundHandler(),
		middleware.Logging(),
		middleware.Recovery(),
	))

	// API routes
	apiHandler := api.NewAPIHandler()
	s.router.Handle("/api/", http.StripPrefix("/api", middleware.Chain(
		apiHandler,
		middleware.Logging(),
		middleware.Recovery(),
	)))

	// Web app routes
	webHandler := web.NewWebHandler()
	s.router.Handle("/app/", http.StripPrefix("/app", middleware.Chain(
		webHandler,
		middleware.Logging(),
		middleware.Recovery(),
	)))

	// Static files
	fs := http.FileServer(http.Dir("web/static"))
	s.router.Handle("/static/", http.StripPrefix("/static", middleware.Chain(
		fs,
		middleware.Logging(),
		middleware.Recovery(),
	)))

	logger.Info("Routes configured successfully")
}

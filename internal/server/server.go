package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gitKashish/golog/internal/config"
	"github.com/gitKashish/golog/pkg/logger"
)

// Server represents the HTTP server
type Server struct {
	config     *config.Config
	httpServer *http.Server
	router     *http.ServeMux
}

// NewServer creates a new server with the given configuration
func NewServer(cfg *config.Config) *Server {
	router := http.NewServeMux()

	return &Server{
		config: cfg,
		router: router,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// Router returns the server's router
func (s *Server) Router() *http.ServeMux {
	return s.router
}

// Start starts the HTTP server
func (s *Server) Start() error {
	logger.Info("Starting server on port %d", s.config.Server.Port)
	logger.Info("Web App: http://localhost:%d/app/", s.config.Server.Port)
	logger.Info("API: http://localhost:%d/api/", s.config.Server.Port)

	return s.httpServer.ListenAndServe()
}

// Stop gracefully stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	logger.Info("Stopping server...")
	return s.httpServer.Shutdown(ctx)
}

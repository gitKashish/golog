package server

import (
	"fmt"
	"net/http"

	"github.com/gitkashish/golog/internal/handlers"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{port}
}

func (server *Server) Serve() error {
	addr := fmt.Sprintf(":%d", server.port)

	router := http.NewServeMux()

	// API Handler
	apiHandler := handlers.NewAPIHandler()
	router.Handle("/api/", http.StripPrefix("/api", apiHandler))

	// App Handler
	appHandler := handlers.NewAppHandler()
	router.Handle("/app/", http.StripPrefix("/app", appHandler))

	fmt.Printf("starting server on http://localhost%s", addr)
	return http.ListenAndServe(addr, router)
}

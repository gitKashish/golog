package server

import (
	"fmt"
	"net/http"

	"github.com/gitKashish/golog/internal/handlers"
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

	fmt.Printf("GOLOG - Starting Server\nWeb App : http://localhost%s/app/\nAPI : http://localhost%s/api/", addr, addr)
	return http.ListenAndServe(addr, router)
}

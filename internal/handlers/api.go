package handlers

import (
	"net/http"
)

func NewAPIHandler() *http.ServeMux {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /greet", handleGreet)

	return router
}

func handleGreet(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello, World!"))
}

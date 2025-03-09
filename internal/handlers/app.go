package handlers

import (
	"html/template"
	"net/http"

	"github.com/gitKashish/golog/internal/web"
)

var tmpl *template.Template = web.NewTemplate()

func NewAppHandler() *http.ServeMux {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /", handleHome)

	return router
}

// Render Home Page
func handleHome(response http.ResponseWriter, request *http.Request) {
	data := map[string]any{
		"Title": "Test Title",
		"Log":   "",
	}
	tmpl.ExecuteTemplate(response, "index.html", data)
}

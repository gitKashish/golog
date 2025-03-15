package web

import (
	"html/template"
	"net/http"

	"github.com/gitKashish/golog/pkg/logger"
	"github.com/gitKashish/golog/web"
)

var tmpl *template.Template = web.NewTemplate()

// NewWebHandler creates a new web handler
func NewWebHandler() *http.ServeMux {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /", handleHome)

	return router
}

// handleHome renders the home page
func handleHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "GoLog - Log Formatter",
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		logger.Error("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

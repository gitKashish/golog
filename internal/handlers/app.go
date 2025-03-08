package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gitkashish/golog/internal/web"
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
	fmt.Print(tmpl.DefinedTemplates())

	data := map[string]any{
		"Title": "Test Title",
	}

	tmpl.ExecuteTemplate(response, "index.html", data)
}

package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gitKashish/golog/internal/core/parser"
	"github.com/gitKashish/golog/pkg/logger"
	"github.com/gitKashish/golog/web"
)

var tmpl *template.Template = web.NewTemplate()

// NewAPIHandler creates a new API handler
func NewAPIHandler() *http.ServeMux {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /greet", handleGreet)
	router.HandleFunc("POST /format", handleFormat)

	return router
}

// handleGreet handles the greeting endpoint
func handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// handleFormat handles log formatting requests
func handleFormat(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		logger.Error("Error parsing form: %v", err)
		return
	}

	p := parser.NewTemplateParser()
	err = p.SetTemplate(r.FormValue("source_template"), r.FormValue("target_template"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		logger.Error("Error setting template: %v", err)
		return
	}

	logs := []string{}
	for _, line := range strings.Split(r.FormValue("raw_log"), "\n") {
		if line != "" {
			logs = append(logs, p.Parse(line))
		}
	}

	data := map[string]any{
		"Logs": logs,
	}

	tmpl.ExecuteTemplate(w, "form.log.pretty", data)
}

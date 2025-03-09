package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gitkashish/golog/internal/core"
)

func NewAPIHandler() *http.ServeMux {
	router := http.NewServeMux()

	// Register Routes
	router.HandleFunc("GET /greet", handleGreet)
	router.HandleFunc("POST /format", handleFormat)

	return router
}

func handleGreet(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello, World!"))
}

// Handle Log formatting requests and sending textfield with formatted logs back
func handleFormat(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(response, err.Error())
	}

	template, err := core.GetTemplateFromLiterals(request.FormValue("source_template"), request.FormValue("target_template"))
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(response, err.Error())
	}

	logs := []string{}
	for line := range strings.Lines(request.FormValue("raw_log")) {
		logs = append(logs, template.Parse(line))
	}

	data := map[string]any{
		"Logs": logs,
	}

	tmpl.ExecuteTemplate(response, "form.log.pretty", data)
}

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

func handleFormat(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(response, "Invalid Request Data")
	}
	logs := ""
	for line := range strings.Lines(request.FormValue("log")) {
		formattedLog := core.ParseLogLine(line)
		logs = logs + formattedLog
	}

	data := map[string]any{
		"Log": logs,
	}

	tmpl.ExecuteTemplate(response, "log.form.field", data)
}

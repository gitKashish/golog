package web

import (
	"embed"
	"html/template"
)

//go:embed templates/*
var templates embed.FS

func NewTemplate() *template.Template {
	tmpl := template.Must(template.ParseFS(templates, "templates/layouts/*.html", "templates/partials/*.html", "templates/*.html"))
	return tmpl
}

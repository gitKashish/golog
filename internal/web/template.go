package web

import (
	"embed"
	"html/template"
)

//go:embed views/*
var views embed.FS

func NewTemplate() *template.Template {
	tmpl := template.Must(template.ParseFS(views, "views/layouts/*.html", "views/partials/*.html", "views/*.html"))
	return tmpl
}

package web

import "html/template"

func NewTemplate() *template.Template {
	tmpl := template.Must(template.ParseGlob("internal/web/views/layouts/*.html"))
	template.Must(tmpl.ParseGlob("internal/web/views/partials/*.html"))
	template.Must(tmpl.ParseGlob("internal/web/views/*.html"))

	return tmpl
}

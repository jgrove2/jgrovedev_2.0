package handlers

import "html/template"

func HandleTemplates() *template.Template {
	tmpl := template.Must(template.ParseGlob("templates/*"))
	return tmpl
}

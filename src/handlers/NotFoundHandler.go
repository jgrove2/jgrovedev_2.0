package handlers

import (
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := HandleTemplates()
	tmpl.ExecuteTemplate(w, "404", nil)
}

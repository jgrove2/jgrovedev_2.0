package handlers

import (
	"net/http"

	"github.com/jgrove2/jgrovedev_2.0/article"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	PostData := article.GetAll()
	tmpl := HandleTemplates()
	tmpl.ExecuteTemplate(w, "home", PostData)
	tmpl.Execute(w, PostData)
}

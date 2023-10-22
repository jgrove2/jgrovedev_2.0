package handlers

import (
	"html/template"
	"net/http"

	"github.com/jgrove2/jgrovedev_2.0/src/util"
)

func MDToHTMLHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := HandleTemplates()
	tmpl.ExecuteTemplate(w, "markdowntohtml", nil)
}

func MDRenderHandler(w http.ResponseWriter, r *http.Request) {
	markdown := r.PostFormValue("markdown")
	html := string(util.MdToHTML([]byte(markdown)))
	tmpl, _ := template.New("t").Parse(html)
	tmpl.Execute(w, nil)
}

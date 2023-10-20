package handlers

import (
	"net/http"

	"github.com/jgrove2/jgrovedev_2.0/src/article"
	"github.com/jgrove2/jgrovedev_2.0/src/util"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	PostData := article.GetAll()
	for i := 0; i < len(PostData); i++ {
		PostData[i].Date = util.FormatTime(PostData[i].Date)
	}
	tmpl := HandleTemplates()
	tmpl.ExecuteTemplate(w, "home", PostData)
	tmpl.Execute(w, PostData)
}

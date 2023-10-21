package handlers

import (
	"log"
	"net/http"

	"github.com/jgrove2/jgrovedev_2.0/src/article"
	"github.com/jgrove2/jgrovedev_2.0/src/util"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	PostData, err := article.GetAll()
	if err != nil {
		log.Printf("Error getting all posts: %v", err)
	} else {
		for i := 0; i < len(PostData); i++ {
			PostData[i].Date, err = util.FormatTime(PostData[i].Date)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				break
			}
		}
		tmpl := HandleTemplates()
		tmpl.ExecuteTemplate(w, "home", PostData)
	}

}

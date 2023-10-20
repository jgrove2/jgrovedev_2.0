package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jgrove2/jgrovedev_2.0/src/database"
	"github.com/jgrove2/jgrovedev_2.0/src/handlers"
)

var tmpl *template.Template

func init() {
	tmpl = handlers.HandleTemplates()
}



func main() {
	err := database.OpenDatabase()
	if err != nil {
		log.Println("error connecting to postgresql database &v", err)
	}
	defer database.CloseDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/post/{id:[0-9]+}", handlers.PostHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

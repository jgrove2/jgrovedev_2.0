package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jgrove2/jgrovedev_2.0/src/database"
	"github.com/jgrove2/jgrovedev_2.0/src/handlers"
	"github.com/jgrove2/jgrovedev_2.0/src/util"
)

var tmpl *template.Template

func init() {
	tmpl = handlers.HandleTemplates()
}

func main() {
	log.Println("Connecting to database...")
	err := database.OpenDatabase()
	if err != nil {
		log.Println("error connecting to postgresql database &v", err)
	}
	defer database.CloseDatabase()
	log.Println("Connected to the database.")
	port := util.GoDotEnvVariable("PORT")
	ip := util.GoDotEnvVariable("IP")
	addr := fmt.Sprintf("%s:%s", ip, port)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/post/{id:[0-9]+}", handlers.PostHandler)
	// Markdown Rendering Project
	r.HandleFunc("/markdown", handlers.MDToHTMLHandler)
	r.HandleFunc("/markdown/render/", handlers.MDRenderHandler)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	log.Fatal(http.ListenAndServe(addr, r))
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Skill struct {
	Name string
}

type Post struct {
	Type        string
	Title       string
	Date        string
	Skills      []Skill
	Description string
	Link        string
}

type Posts struct {
	Posts []Post
}

func main() {

	posts := Posts{
		Posts: []Post{
			{
				Type:        "Project",
				Title:       "Test 1",
				Date:        "2023-10-10",
				Skills:      []Skill{{Name: "Go"}, {Name: "HTMX"}},
				Description: "This is a test description. With these descriptions I can describe something that is to be described in a description...",
				Link:        "#",
			},
			{
				Type:        "Project",
				Title:       "Test 2",
				Date:        "2023-10-11",
				Skills:      []Skill{{Name: "Go"}, {Name: "HTMX"}},
				Description: "This is a test description. With these descriptions I can describe something that is to be described in a description...",
				Link:        "#",
			},
			{
				Type:        "Project",
				Title:       "Test 3",
				Date:        "2023-10-12",
				Skills:      []Skill{{Name: "Go"}, {Name: "HTMX"}},
				Description: "This is a test description. With these descriptions I can describe something that is to be described in a description...",
				Link:        "#",
			},
			{
				Type:        "Project",
				Title:       "Test 5",
				Date:        "2023-10-13",
				Skills:      []Skill{{Name: "Go"}, {Name: "HTMX"}},
				Description: "This is a test description. With these descriptions I can describe something that is to be described in a description...",
				Link:        "#",
			},
			{
				Type:        "Project",
				Title:       "Test 9",
				Date:        "2023-10-14",
				Skills:      []Skill{{Name: "Go"}, {Name: "HTMX"}},
				Description: "This is a test description. With these descriptions I can describe something that is to be described in a description...",
				Link:        "#",
			},
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../../templates/index.html"))
		tmpl.Execute(w, posts)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

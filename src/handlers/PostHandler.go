package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jgrove2/jgrovedev_2.0/src/article"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Post    article.Article
	Article template.HTML
	Creator article.Creator
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PostData := article.GetPost(vars["id"])
	CreatorData := article.GetCreatorDetails(PostData.Creator)
	articleTemplate := template.HTML(string(mdToHTML([]byte(PostData.Article))))
	PostContent := Post{Post: PostData, Article: articleTemplate, Creator: CreatorData}
	tmpl := HandleTemplates()
	tmpl.ExecuteTemplate(w, "post", PostContent)
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

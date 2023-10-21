package handlers

import (
	"html/template"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/jgrove2/jgrovedev_2.0/src/article"
	"github.com/jgrove2/jgrovedev_2.0/src/util"

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
	PostData, err := article.GetPost(vars["id"])
	if err != nil || reflect.DeepEqual(PostData, article.Article{}) {
		NotFoundHandler(w, r)
		return
	}
	CreatorData, err := article.GetCreatorDetails(PostData.Creator)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	PostData.Date, err = util.FormatTime(PostData.Date)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
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

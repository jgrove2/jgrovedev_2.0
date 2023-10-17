package article

import (
	"fmt"
	"log"

	"github.com/jgrove2/jgrovedev_2.0/database"
	"github.com/lib/pq"
)

type Article struct {
	ID          int      `json:"id"`
	Date        string   `json:"created_at"`
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Skills      []string `json:"skills"`
	Description string   `json:"description"`
	Article     string   `json:"article"`
	Creator     int      `json:"creator"`
}

type Creator struct {
	ID       int    `json:"id"`
	Name     string `json:"full_name"`
	GitHub   string `json:"github"`
	Email    string `json:"email"`
	LinkedIn string `json:"linkedin"`
}

func GetAll() []Article {
	var articles []Article
	rows, err := database.DB.Query(`SELECT id, created_at, type, title, skills, description FROM "Articles";`)
	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Date, &article.Type, &article.Title, pq.Array(&article.Skills), &article.Description); err != nil {
			return nil
		}
		articles = append(articles, article)
	}

	return articles
}

func GetPost(id string) Article {
	var article Article
	var query = `SELECT id, created_at, type, title, skills, article, creator FROM "Articles" WHERE id=` + id
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println(err)
		return article
	}
	for rows.Next() {
		var temp Article
		if err := rows.Scan(&temp.ID, &temp.Date, &temp.Type, &temp.Title, pq.Array(&temp.Skills), &temp.Article, &temp.Creator); err != nil {
			log.Println(err)
			return article
		}
		article = temp
	}

	return article
}

func GetCreatorDetails(id int) Creator {
	var creator Creator
	var query = fmt.Sprintf(`SELECT id, full_name, github, email, linkedin FROM "Creators" WHERE id=%d`, id)
	rows, err := database.DB.Query(query)
	if err != nil {
		return creator
	}
	for rows.Next() {
		var temp Creator
		if err := rows.Scan(&temp.ID, &temp.Name, &temp.GitHub, &temp.Email, &temp.LinkedIn); err != nil {
			return creator
		}
		creator = temp
	}

	return creator
}

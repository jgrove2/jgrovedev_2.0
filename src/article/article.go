package article

import (
	"errors"
	"fmt"
	"log"

	"github.com/jgrove2/jgrovedev_2.0/src/database"
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

func GetAll() ([]Article, error) {
	var articles []Article
	rows, err := database.DB.Query(`SELECT id, created_at, type, title, skills, description FROM prod.posts`)
	if err != nil {
		return articles, errors.New("503")
	}

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Date, &article.Type, &article.Title, pq.Array(&article.Skills), &article.Description); err != nil {
			return articles, errors.New("404")
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func GetPost(id string) (Article, error) {
	var article Article
	var query = `SELECT id, created_at, type, title, skills, article, creator FROM prod.posts WHERE id=` + id
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println("Error getting response from db: &v", err)
		return article, errors.New("503")
	}
	for rows.Next() {
		var temp Article
		if err := rows.Scan(&temp.ID, &temp.Date, &temp.Type, &temp.Title, pq.Array(&temp.Skills), &temp.Article, &temp.Creator); err != nil {
			log.Printf("Error scanning query: %v", err)
			return article, errors.New("404")
		}
		if temp.ID == 0 {
			return article, errors.New("404")
		}
		article = temp
	}

	return article, nil
}

func GetCreatorDetails(id int) (Creator, error) {
	var creator Creator
	var query = fmt.Sprintf(`SELECT id, full_name, github, email, linkedin FROM prod.creators WHERE id=%d`, id)
	rows, err := database.DB.Query(query)
	if err != nil {
		return creator, errors.New("503")
	}
	for rows.Next() {
		var temp Creator
		if err := rows.Scan(&temp.ID, &temp.Name, &temp.GitHub, &temp.Email, &temp.LinkedIn); err != nil {
			return creator, errors.New("404")
		}
		creator = temp
	}

	return creator, nil
}

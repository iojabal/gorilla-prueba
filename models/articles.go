package models

import (
	"prueba/db"
)

type Article struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Userid      User   `json:"userid"`
}

type Articles []Article

//metodos
func (a *Article) Save() {
	if a.Id == 0 {
		a.insert()
	} else {
		a.update()
	}
}

func (a *Article) update() {
	sql := "UPDATE articles SET name=?, category=?, description=?"
	db.Exec(sql, a.Name, a.Category, a.Description)
}

func (a *Article) insert() {
	sql := "INSERT articles SET name=?, category=?, description=?, userid=?"
	result, _ := db.Exec(sql, a.Name, a.Category, a.Description, a.Userid)
	// if err != nil {
	// 	fmt.Println("este es el error", err)
	// }
	a.Id, _ = result.LastInsertId()
}

//fin metodos

//inicio funciones
func NewArticle(Name, category, description string, userid User) *Article {
	article := &Article{
		//Id:          id,
		Name:        Name,
		Category:    category,
		Description: description,
		Userid:      userid,
	}
	return article
}

func CreateArticle(Name, category, description string, userid User) *Article {
	article := NewArticle(Name, category, description, userid)
	article.Save()
	return article
}

func ListArticle() (Articles, error) {
	sql := "SELECT id, name, category, description, userid from articles"
	articles := Articles{}
	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			article := Article{}
			rows.Scan(&article.Id, &article.Name, &article.Category, &article.Description, &article.Userid)
			articles = append(articles, article)
		}
		return articles, nil
	}

}

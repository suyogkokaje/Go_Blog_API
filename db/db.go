package db

import(
	"github.com/suyogkokaje/Go_Blog_API/models"
)

var Articles []models.Article

func InitDB(){
	Articles = []models.Article{
		{Id: "1", Title: "Article1", Desc: "Article Description 1", Content: "Article Content 1"},
		{Id: "2", Title: "Article2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
}


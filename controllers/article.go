package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/suyogkokaje/Go_Blog_API/models"
	"github.com/suyogkokaje/Go_Blog_API/db"
	"github.com/gorilla/mux"
)


func GetArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.Articles)
}

func GetSpecificArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	for _, article := range db.Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	db.Articles = append(db.Articles, article)
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	for index, article := range db.Articles {
		if article.Id == key {
			db.Articles = append(db.Articles[:index], db.Articles[index+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	for index, article := range db.Articles {
		if article.Id == key {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var updatedArticle models.Article
			json.Unmarshal(reqBody, &updatedArticle)
			db.Articles[index].Title = updatedArticle.Title
			db.Articles[index].Desc = updatedArticle.Desc
			db.Articles[index].Content = updatedArticle.Content
			json.NewEncoder(w).Encode(db.Articles[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
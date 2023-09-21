package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/suyogkokaje/Go_Blog_API/models" 
    "github.com/suyogkokaje/Go_Blog_API/db"     
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
    var articles []models.Article
    db.DB.Find(&articles)
    json.NewEncoder(w).Encode(articles)
}

func GetSpecificArticle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    key := params["id"]
    var article models.Article
    if err := db.DB.Where("id = ?", key).First(&article).Error; err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
    var article models.Article
    json.NewDecoder(r.Body).Decode(&article)
    db.DB.Create(&article)
    json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    key := params["id"]
    var article models.Article
    if err := db.DB.Where("id = ?", key).First(&article).Error; err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    db.DB.Delete(&article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    key := params["id"]
    var article models.Article
    if err := db.DB.Where("id = ?", key).First(&article).Error; err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    json.NewDecoder(r.Body).Decode(&article)
    db.DB.Save(&article)
    json.NewEncoder(w).Encode(article)
}

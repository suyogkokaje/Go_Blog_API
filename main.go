package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func getSpecificArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	for index, article := range Articles {
		if article.Id == key {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var updatedArticle Article
			json.Unmarshal(reqBody, &updatedArticle)
			Articles[index].Title = updatedArticle.Title
			Articles[index].Desc = updatedArticle.Desc
			Articles[index].Content = updatedArticle.Content
			json.NewEncoder(w).Encode(Articles[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", getSpecificArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	Articles = []Article{
		{Id: "1", Title: "Article1", Desc: "Article Description 1", Content: "Article Content 1"},
		{Id: "2", Title: "Article2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	handleRequests()
}

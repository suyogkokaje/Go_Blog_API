package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/suyogkokaje/Go_Blog_API/db"
	"github.com/suyogkokaje/Go_Blog_API/controllers"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	router.HandleFunc("/article", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", controllers.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", controllers.GetSpecificArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", controllers.UpdateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	db.InitDB()
	handleRequests()
}

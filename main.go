package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/suyogkokaje/Go_Blog_API/db"
	"github.com/suyogkokaje/Go_Blog_API/controllers"
	"github.com/suyogkokaje/Go_Blog_API/models"
)

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
	loadEnvVariables()

	config := models.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}


    db.InitDB(config)
    defer db.DB.Close()
	handleRequests()
}

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/suyogkokaje/Go_Blog_API/models"
)

var (
    DB *gorm.DB
)

func InitDB(config models.DatabaseConfig) {
    connectionString := "host=" + config.Host +
        " port=" + config.Port +
        " user=" + config.User +
        " dbname=" + config.DBName +
        " password=" + config.Password +
        " sslmode=" + config.SSLMode

    var err error
    DB, err = gorm.Open("postgres", connectionString)
    if err != nil {
        panic("Failed to connect to database")
    }
    DB.AutoMigrate(&models.Article{})
}
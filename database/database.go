package database

import (
	"fmt"
	"log"
	movie "movie/models/movie_model"
	detail "movie/models/specification_model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DataMigration() {
	// here we get the data from enverment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	urlDSN := os.Getenv("urlDSN")
	DB, err := gorm.Open(postgres.Open(urlDSN), &gorm.Config{})
	// DB, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	DB.AutoMigrate(movie.Movie{}, detail.Specification{})

	Database = DB
}

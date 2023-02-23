package database

import (
	"fmt"
	movie "movie/models/movie_model"
	detail "movie/models/specification_model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/movie?parseTime=true"

var err error

func DataMigration() {
	// DB, err := gorm.Open(postgres.Open(urlDSN), &gorm.Config{})
	DB, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	DB.AutoMigrate(movie.Movie{}, detail.Specification{})

	Database = DB
}

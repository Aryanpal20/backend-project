package database

import (
	"fmt"
	emp "gin/model/emp_model"
	user "gin/model/user_model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/microservice"

func DataMigration() {
	DB, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println("Conection Error")
		panic(err.Error())
	}
	DB.AutoMigrate(user.User{}, emp.Employee{})

	Database = DB
}

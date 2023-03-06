package main

import (
	"gin/controller/auth"
	set "gin/controller/set_data"
	"gin/database"
	"gin/middelware"

	"github.com/gin-gonic/gin"
)

func main() {

	database.DataMigration()

	// middelware.CSVfile()

	r := gin.Default()

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/postcsvdata", middelware.AuthRequired(), middelware.AllowOnlyCsvAndXlsx(), set.PostCSVData)
	r.Run()
}

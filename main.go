package main

import (
	movie "movie/controller/movie"
	detail "movie/controller/specification"
	"movie/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DataMigration()

	r := gin.Default()

	r.POST("/postmovie", movie.PostMovie)
	r.GET("/getmoviebyid/:id", movie.GetMovieByID)
	r.GET("/getmoviealldetail", movie.GetMovieAllDetail)
	r.GET("/getallmovie", movie.GetAllMovie)
	r.PATCH("/updatemovie/:id", movie.UpdateMovie)
	r.DELETE("/deletemovie/:id", movie.DeleteMovie)
	r.POST("/postspecification", detail.PostSpecificatiion)
	r.GET("/getspecification/:id", detail.GetSpecification)
	r.PATCH("/updatespecification/:id", detail.UpdateSpecification)
	r.DELETE("/deletespecification/:id", detail.DeleteSpecification)

	r.Run()
}

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

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.POST("/postmovie", movie.PostMovie)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}
		//...
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

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

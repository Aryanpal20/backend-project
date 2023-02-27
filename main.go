package main

import (
	"movie/controller/movie"
	"movie/database"
	"movie/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	database.DataMigration()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a Movie rating app."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "api/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.Default()
	v1 := r.Group(docs.SwaggerInfo.BasePath)
	{
		movies := v1.Group("/getmoviebyid")
		{
			movies.POST("/getmoviebyid", movie.GetMovieByID)
			// accounts.GET("", c.ListAccounts)
			// accounts.POST("", c.AddAccount)
			// accounts.DELETE(":id", c.DeleteAccount)
			// accounts.PATCH(":id", c.UpdateAccount)
			// accounts.POST(":id/images", c.UploadAccountImage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/postmovie", movie.PostMovie)
	r.GET("/getmoviebyid", movie.GetMovieByID)
	r.GET("/getmoviealldetail", movie.GetMovieAllDetail)
	r.GET("/getallmovie", movie.GetAllMovie)
	r.PATCH("/updatemovie/:id", movie.UpdateMovie)
	r.DELETE("/deletemovie/:id", movie.DeleteMovie)
	// r.POST("/postspecification", detail.PostSpecificatiion)
	// r.GET("/getspecification/:id", detail.GetSpecification)
	// r.PATCH("/updatespecification/:id", detail.UpdateSpecification)
	// r.DELETE("/deletespecification/:id", detail.DeleteSpecification)

	r.Run()
}

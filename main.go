package main

import (
	"movie/controller/movie"
	detail "movie/controller/specification"
	"movie/database"
	"movie/docs"
	"net/http"

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
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/v1")
	{
		api.POST("/postmovie", movie.PostMovie)
		api.GET("/getmoviebyid/:id", movie.GetMovieByID)
		api.GET("/getmoviealldetail", movie.GetMovieAllDetail)
		api.GET("/getallmovie", movie.GetAllMovie)
		api.PATCH("/updatemovie/:id", movie.UpdateMovie)
		api.DELETE("/deletemovie/:id", movie.DeleteMovie)
		api.POST("/postspecification", detail.PostSpecificatiion)
		api.GET("/getspecification/:id", detail.GetSpecification)
		api.PATCH("/updatespecification/:id", detail.UpdateSpecification)
		api.DELETE("/deletespecification/:id", detail.DeleteSpecification)

		// Handle GET users request
	}
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

package movie

import (
	"fmt"
	"movie/database"
	movie "movie/models/movie_model"
	detail "movie/models/specification_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Post a New Movie
// @Description Create a new movie with the input paylod
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {object} movie
// @Router /postmovie [post]

func PostMovie(c *gin.Context) {
	var input movie.Movie
	var movies movie.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := movie.Movie{Movie_Name: input.Movie_Name}
	database.Database.Where("movie_name = ?", input.Movie_Name).Find(&movies)
	if movies.Movie_Name != input.Movie_Name {
		database.Database.Create(&movie)
		c.JSON(http.StatusAccepted, gin.H{"Data": movie})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "This movie already present"})
	}

}

func GetMovieByID(c *gin.Context) {
	var movies movie.Movie
	if err := database.Database.Where("id = ?", c.Param("id")).First(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": movies})
}

// here we get both table data.
func GetMovieAllDetail(c *gin.Context) {
	var movies movie.Movie
	var details []detail.Specification
	detail := c.PostForm("detail")
	database.Database.Where("movie_name = ?", detail).Find(&movies)
	database.Database.Where("movie_id = ?", movies.ID).Preload("Specifications").Find(&details)
	movies.Specification = append(movies.Specification, details...)
	c.JSON(http.StatusOK, gin.H{"Data": movies})
}

func GetAllMovie(c *gin.Context) {
	var movies []movie.Movie
	if err := database.Database.Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	fmt.Println(movies)
	c.JSON(http.StatusOK, gin.H{"Data": movies})
}

func UpdateMovie(c *gin.Context) {
	var movies movie.Movie
	var input movie.Movie
	if err := database.Database.Where("id = ?", c.Param("id")).First(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := movie.Movie{Movie_Name: input.Movie_Name}
	database.Database.Model(&movies).Updates(&movie)
	c.JSON(http.StatusOK, gin.H{"Data": movies})
}

func DeleteMovie(c *gin.Context) {
	var movies movie.Movie
	if err := database.Database.Where("id = ?", c.Param("id")).First(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.Database.Delete(&movies)
	c.JSON(http.StatusOK, gin.H{"Data": "your data deleted successfully !!!"})
}

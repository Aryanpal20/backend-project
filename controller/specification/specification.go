package specification

import (
	"movie/database"
	detail "movie/models/specification_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSpecificatiion(c *gin.Context) {
	var input detail.Specification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	detail := detail.Specification{Length: input.Length, Original_Language: input.Original_Language, Year_of_production: input.Year_of_production,
		Director_Name: input.Director_Name, Rating: input.Rating, Genres: input.Genres, Cast: input.Cast, Movie_ID: input.Movie_ID}
	database.Database.Create(&detail)
	c.JSON(http.StatusAccepted, gin.H{"Data": detail})
}

func GetSpecification(c *gin.Context) {
	var details detail.Specification
	if err := database.Database.Where("id = ?", c.Param("id")).First(&details).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": details})
}

func UpdateSpecification(c *gin.Context) {
	var details detail.Specification
	var input detail.Specification
	if err := database.Database.Where("movie_id = ?", c.Param("id")).First(&details).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	detail := detail.Specification{Rating: input.Rating}
	database.Database.Model(&details).Updates(&detail)
	c.JSON(http.StatusOK, gin.H{"Data": details})
}

func DeleteSpecification(c *gin.Context) {
	var detail detail.Specification
	if err := database.Database.Where("id = ?", c.Param("id")).First(&detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.Database.Delete(&detail)
	c.JSON(http.StatusOK, gin.H{"Data": "your data deleted successfully !!!"})
}

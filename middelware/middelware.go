package middelware

import (
	"fmt"
	"gin/controller/auth"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(auth.JwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Token is expired"})
			panic("invalid token")

		} else {
			// c.JSON(http.StatusAccepted, gin.H{"message": "Token is valid"})
			claims := token.Claims.(jwt.MapClaims)["id"]
			fmt.Println("token vaild")
			c.Set("id", claims)
			c.Next()

		}
	}
}

func AllowOnlyCsvAndXlsx() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please upload a file"})
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".csv" && ext != ".xlsx" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Only CSV and XLSX files are allowed"})
			return
		}

		c.Next()
	}
}

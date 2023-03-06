package auth

import (
	"fmt"
	"gin/database"
	user "gin/model/user_model"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// type jwtToken struct {
// 	Token string `json:"token"`
// }

var JwtKey = []byte("Jwt_Key")

func Login(c *gin.Context) {

	var input user.User
	var users user.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := user.User{Email: input.Email, Password: input.Password}
	database.Database.Where("email = ?", user.Email).First(&users)

	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password))

	if err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    users.ID,
			"email": users.Email,
			"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		})

		tokenstring, error := token.SignedString(JwtKey)
		c.JSON(http.StatusAccepted, gin.H{"Token": tokenstring})
		if error != nil {
			fmt.Println(err)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaild email and password"})
	}
}

func Register(c *gin.Context) {

	var input user.User
	var users user.User

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := user.User{Email: input.Email, Password: input.Password}

	database.Database.Where("email = ?", user.Email).Find(&users)
	if users.Email != user.Email {
		Password := []byte(string(user.Password))

		hashedPassword, err := bcrypt.GenerateFromPassword(Password, 10)

		if err != nil {
			panic(err)
		}

		//err = bcrypt.CompareHashAndPassword(hashedPassword, Password)
		user.Password = string(hashedPassword)

		database.Database.Create(&user)
		c.JSON(http.StatusCreated, gin.H{"Message": "Registration Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email already exist !!!"})
	}

}

package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	postgres "github.com/saracha-06422/go-authen/databases"
	"github.com/saracha-06422/go-authen/entity"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type RegisterBody struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Avatar   string `json:"avatar"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Check User Exists
	var userExist entity.User
	postgres.Db.Where("username = ?", json.UserName).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "User Exists",
		})
		return
	}

	//Create User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := entity.User{Username: json.UserName, Password: string(encryptedPassword), Fullname: json.FullName, Avatar: json.Avatar}
	postgres.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "User Create Success",
			"userId":  user.ID,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "User Create Failed",
		})
	}

}

type LoginBody struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check User Exists
	var userExist entity.User
	postgres.Db.Where("username = ?", json.UserName).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "User Dose Not Exists",
		})
		return
	}

	//Check Password encrypt and not encrypt
	errBcrypt := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if errBcrypt == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 1).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "OK",
			"message": "Login Success",
			"token":   tokenString,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Login Failed",
		})
	}
}

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saracha-06422/go-authen/databases"
	"github.com/saracha-06422/go-authen/entity"
)

func ReadAll(c *gin.Context) {
	var users []entity.User
	databases.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Users Read Success",
		"users":   users,
	})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId")
	var user entity.User
	databases.Db.First(&user, userId)

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Users Read Success",
		"users":   user,
		"userId":  userId,
	})
}

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//เป็นการ set ค่่าเข้าไปใน context
			c.Set("userId", claims["userId"])
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "Forbidden",
				"message": err.Error(),
			})
		}
		// before request
		c.Next()
	}
}

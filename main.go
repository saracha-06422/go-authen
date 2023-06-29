package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	AuthController "github.com/saracha-06422/go-authen/controller/auth"
	UsersController "github.com/saracha-06422/go-authen/controller/users"
	Postgres "github.com/saracha-06422/go-authen/databases"
	"github.com/saracha-06422/go-authen/middleware"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	//connect to database
	Postgres.ConnectPostgre()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	//จะเข้าถึง readall ต้องผ่าน middleware ก่อน
	authorized := r.Group(("/users"), middleware.JWTAuthen())
	authorized.GET("/readall", UsersController.ReadAll)
	authorized.GET("/profile", UsersController.Profile)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

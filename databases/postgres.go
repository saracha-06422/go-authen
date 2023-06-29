package databases

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/saracha-06422/go-authen/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectPostgre() {
	sqlInfo := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", os.Getenv("HOST"), 5432, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	Db, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&entity.User{})

	// db, err := sql.Open("postgres", sqlInfo)
	fmt.Println("Successfully database connected!")
}

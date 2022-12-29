package db

import (
	"fmt"
	"os"

	"github.com/HungOnBlog/libma/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDbDns() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
}

func InitDb() {
	db, err := gorm.Open(postgres.Open(getDbDns()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	err = db.AutoMigrate(
		&models.BookTitle{},
		&models.Book{},
		&models.Author{},
		&models.Borrower{},
		&models.Librarian{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated")
}

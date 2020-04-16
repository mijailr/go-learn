package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/mijailr/go-learn/pkg/model"
	"log"
	"os"
)

var db *gorm.DB

func init() {

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
	db.AutoMigrate(
		&Alert{},
	)
}

func Connect() *gorm.DB {
	return db
}

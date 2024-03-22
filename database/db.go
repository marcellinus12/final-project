package database

import (
	"fmt"
	"go-gorm/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database:", err)
	}

	fmt.Println("successfully connected to the database")

	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}

package initializers

import (
	"log"
	"os"
	"url-shortener/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("Error connecting to database")

	}

	DB.AutoMigrate(&models.ResponseURL{})

}

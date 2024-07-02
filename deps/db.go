package deps

import (
	"log"
	"qbc/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateNewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	db.AutoMigrate(&models.EmailLog{})
	db.AutoMigrate(&models.User{})

	return db
}

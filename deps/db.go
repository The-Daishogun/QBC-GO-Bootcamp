package deps

import (
	"log"
	"os"
	"qbc/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateNewDB(dbFileName string) (*gorm.DB, func() error) {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	db.AutoMigrate(&models.EmailLog{})
	db.AutoMigrate(&models.User{})

	return db, func() error {
		return os.Remove(dbFileName)
	}
}

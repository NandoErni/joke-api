package initializers

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database...")
	}

	return db
}

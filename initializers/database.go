package initializers

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "jokes.db" // Local fallback for development
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database...")
	}

	return db
}

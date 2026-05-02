package initializers

import (
	"log"
	"os"

	"github.com/NandoErni/joke-api/models"
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

func SyncDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.Joke{})

	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
}

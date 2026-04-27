package main

import (
	"github.com/NandoErni/joke-api/initializers"
	"github.com/NandoErni/joke-api/models"
)

func main() {
	db := initializers.ConnectToDB()
	db.AutoMigrate(&models.Joke{})
}

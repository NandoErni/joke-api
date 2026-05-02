package repository

import (
	"errors"

	"github.com/NandoErni/joke-api/initializers"
	"github.com/NandoErni/joke-api/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = initializers.ConnectToDB()
	initializers.SyncDatabase(db)
}

func GetJokes() ([]models.Joke, error) {
	var jokes []models.Joke

	result := db.Find(&jokes)

	if result.Error != nil {
		return nil, result.Error
	}

	return jokes, nil
}

func GetJoke(id int) (*models.Joke, error) {
	var joke models.Joke

	result := db.First(&joke, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &joke, nil
}

func CreateJoke(text string) (*models.Joke, error) {
	newJoke := models.Joke{Text: text}

	result := db.Create(&newJoke)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newJoke, nil
}

func GetJokeCount() (int64, error) {
	var count int64

	err := db.Model(&models.Joke{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func DeleteJoke(id int) error {
	result := db.Delete(&models.Joke{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no joke found with that ID")
	}

	return nil
}

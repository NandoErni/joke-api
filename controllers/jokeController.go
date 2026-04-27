package controllers

import (
	"errors"
	"math/rand/v2"
	"net/http"
	"strconv"

	repository "github.com/NandoErni/joke-api/repositories"
	"github.com/gin-gonic/gin"
)

type joke struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var jokessss = []joke{
	{
		ID: 1,
		Text: `Zwei Fisch schwümmed im See. Seit de eint:
„Hesch du en Plan, wie mer do usechömed?“
Seit de ander: „Klar - mir mached eifach d’Flosse ab!“`,
	},
	{
		ID: 2,
		Text: `Lehrer: „Was isch schnäller - Liecht oder Ton?“
Schüler: „Liecht!“
Lehrer: „Wieso?“
Schüler: „Wänn ich de Fernseher a mache, gseh ich s’Bild bevor ich de Lärm ghöre!“`,
	},
	{
		ID: 3,
		Text: `Zwei Küeh stönd uf de Weid. Seit d’eint:
„Du, hesch scho ghört vom Rinderwahnsinn?“
Seit d’ander: „Ja zum Glück bi ich es Pferd!“`,
	},
}

func GetJokes(context *gin.Context) {
	jokes, err := repository.GetJokes()

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Jokes not found"})
	}

	context.IndentedJSON(http.StatusOK, jokes)
}

func getJokeById(id int) (*joke, error) {

	dbJoke, err := repository.GetJoke(id)

	if err != nil {
		return nil, errors.New("Joke not found")
	}

	j := joke{ID: int(dbJoke.ID), Text: dbJoke.Text}

	return &j, nil
}

func GetJoke(context *gin.Context) {
	idStr := context.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	joke, err := getJokeById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Joke not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, joke)
}

func GetRandomJoke(context *gin.Context) {
	count, err := repository.GetJokeCount()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not retireve a random joke"})
		return
	}

	joke, err := getJokeById(rand.IntN(int(count)))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Joke not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, joke)
}

func CreateJoke(context *gin.Context) {
	var newJoke joke

	if err := context.BindJSON(&newJoke); err != nil {
		return
	}

	modelJoke, err := repository.CreateJoke(newJoke.Text)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Joke could not be created"})
		return
	}

	newJoke.ID = int(modelJoke.ID)

	context.IndentedJSON(http.StatusCreated, newJoke)
}

func DeleteJoke(context *gin.Context) {
	idStr := context.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = repository.DeleteJoke(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Joke not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, nil)
}

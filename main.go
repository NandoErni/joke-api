package main

import (
	"github.com/NandoErni/joke-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/jokes", controllers.GetJokes)
	router.GET("/jokes/:id", controllers.GetJoke)
	router.DELETE("/jokes/:id", controllers.DeleteJoke)
	router.GET("/jokes/random", controllers.GetRandomJoke)
	router.POST("/jokes", controllers.CreateJoke)
	router.Run("localhost:8080")
}

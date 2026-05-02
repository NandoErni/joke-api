package main

import (
	"crypto/subtle"
	"fmt"
	"os"

	"github.com/NandoErni/joke-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Starting Server...")
	router := gin.Default()

	// Public routes
	router.GET("/jokes", controllers.GetJokes)
	router.GET("/jokes/:id", controllers.GetJoke)
	router.GET("/jokes/random", controllers.GetRandomJoke)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(APIKeyAuth())
	{
		authorized.POST("/jokes", controllers.CreateJoke)
		authorized.DELETE("/jokes/:id", controllers.DeleteJoke)
	}

	router.Run(":8080")
	fmt.Println("Server running...")
}

func APIKeyAuth() gin.HandlerFunc {
	key := os.Getenv("API_KEY")
	if key == "" {
		return func(c *gin.Context) {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}

	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")

		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(key)) != 1 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

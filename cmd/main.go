package main

import (
	"goingEter/config"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvs()
}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!!",
		})
	})

	router.Run(os.Getenv("SERVER_PORT"))
}

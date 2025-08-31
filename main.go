package main

import (
	"url_shortener/database"
	"url_shortener/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/signup", handlers.SignupHandler)
	r.POST("/signin", handlers.SigninHandler)
	r.Run(":8080")
}

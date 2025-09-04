package main

import (
	"url_shortener/database"
)

func main() {
	database.Connect()
	database.SeedAdmin(database.DB)
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })

	// r.POST("/signup", func(c *gin.Context) {
	// 	handlers.SignupHandler(c, database.DB)
	// })
	// r.POST("/signin", func(c *gin.Context) {
	// 	handlers.SigninHandler(c, database.DB)
	// })
	// r.Run(":8080")
}

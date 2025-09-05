package main

import (
	"url_shortener/database"
	"url_shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.SeedAdmin(database.DB)
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}

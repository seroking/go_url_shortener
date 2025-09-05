package routes

import (
	"url_shortener/database"
	"url_shortener/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	public := router.Group("/api/v1")
	{

		router.POST("/signup", func(c *gin.Context) {
			handlers.SignupHandler(c, database.DB)
		})

		router.POST("/signin", func(c *gin.Context) {
			handlers.SigninHandler(c, database.DB)
		})

	}

	router.PUT("/user/:id", func(c *gin.Context) {
		handlers.UpdateUser(c, database.DB)
	})

	router.DELETE("/user/:id", func(c *gin.Context) {
		handlers.DeleteUser(c, database.DB)
	})

}

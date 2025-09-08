package routes

import (
	"url_shortener/database"
	"url_shortener/handlers"
	"url_shortener/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.GET("/:shortcode", func(c *gin.Context) {
		handlers.RedirectOriginalUrl(c, database.DB)
	})
	public := router.Group("/api/v1")
	{

		public.POST("/signup", func(c *gin.Context) {
			handlers.SignupHandler(c, database.DB)
		})

		public.POST("/signin", func(c *gin.Context) {
			handlers.SigninHandler(c, database.DB)
		})

	}
	protected := router.Group("/api/v1")
	protected.Use(middlewares.AuthenticationMiddleware())
	{
		protected.GET("/links", func(c *gin.Context) {
			handlers.ListUserLinks(c, database.DB)
		})
		protected.POST("/links", func(c *gin.Context) {
			handlers.CreateLink(c, database.DB)
		})
		protected.DELETE("/links/:id", func(c *gin.Context) {
			handlers.DeleteLink(c, database.DB)
		})
	}
	{
		//User profile routes (regular authenticated users)
		protected.GET("/profile", func(c *gin.Context) {
			handlers.GetUserProfile(c, database.DB)
		})
		protected.PUT("/profile", func(c *gin.Context) {
			handlers.UpdateUserProfile(c, database.DB)
		})

		adminRoutes := protected.Group("")
		adminRoutes.Use(middlewares.AdminOnly(database.DB))
		{
			adminRoutes.GET("/users", func(c *gin.Context) {
				handlers.ListUsers(c, database.DB)
			})

			adminRoutes.GET("/users/:id", func(c *gin.Context) {
				handlers.GetUser(c, database.DB)
			})

			adminRoutes.PUT("/users/:id", func(c *gin.Context) {
				handlers.UpdateUser(c, database.DB)
			})

			adminRoutes.DELETE("/users/:id", func(c *gin.Context) {
				handlers.DeleteUser(c, database.DB)
			})
		}

	}
}

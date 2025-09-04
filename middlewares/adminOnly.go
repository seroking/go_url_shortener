package middlewares

import (
	"url_shortener/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func adminOnly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		user_id, exist := c.Get("user_id")
		if !exist {
			c.JSON(400, gin.H{"error": "Invalid user_id"})
			return
		}

		db.First(&user, user_id)

		if user.Role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}

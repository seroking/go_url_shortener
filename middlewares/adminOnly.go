package middlewares

import "github.com/gin-gonic/gin"

func adminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")
		if !exist || role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}

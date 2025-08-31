package handlers

import (
	"url_shortener/models"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

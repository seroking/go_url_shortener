package handlers

import "github.com/gin-gonic/gin"

func SigninHandler(c *gin.Context) {

	c.JSON(200, gin.H{"message": "sign in Page working..."})

}

package handlers

import (
	"url_shortener/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignupInput struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignupHandler(c *gin.Context, db *gorm.DB) {
	var input SignupInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if input.Name == "" || input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}
	user := models.User{
		Username:     input.Name,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}
	if db.Create(&user).Error != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

package handlers

import (
	"url_shortener/models"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type signinInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SigninHandler(c *gin.Context, db *gorm.DB) {
	var input signinInput
	var user models.User

	secretKey := os.Getenv("JWT_SECRET")
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := db.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful!", "token": signedToken})

}

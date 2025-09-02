package handlers

// i still should treat the problems of signup with a mail that already exist in the database
import (
	"errors"
	"log"
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

var result models.User

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

	err = db.Where("username = ?", input.Name).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
	} else if err != nil {
		log.Println("internal server error: ", err)
		return
	} else {
		c.JSON(409, gin.H{"error": "this username already exist"})
		return
	}

	err = db.Where("email = ?", input.Email).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

	} else if err != nil {
		log.Println("internal server error: ", err)
		return
	} else {
		c.JSON(409, gin.H{"error": "this email already exist"})
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

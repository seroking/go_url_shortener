package database

import (
	"log"
	"os"
	"url_shortener/models"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var admin models.User
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found.")
	}
	result := db.First(&admin, "role = ?", "admin")
	if result.Error == gorm.ErrRecordNotFound {
		username := os.Getenv("ADMIN_NAME")
		email := os.Getenv("ADMIN_EMAIL")
		password := os.Getenv("ADMIN_PASSWORD")
		role := "admin"
		if username == "" || email == "" || password == "" {
			log.Fatal("missing admin credentials")
		}

		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("password check failed:", err)
		}
		admin = models.User{
			Username:     username,
			Email:        email,
			PasswordHash: string(hashedpassword),
			Role:         role,
		}
		if db.Create(&admin).Error != nil {
			log.Fatal("failed to create admin user:", err)
		}

	}
}

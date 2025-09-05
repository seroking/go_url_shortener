package handlers

import (
	"url_shortener/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, db *gorm.DB) {
	var Input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		Username:     Input.Username,
		Email:        Input.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})

}

func ListUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	c.JSON(200, gin.H{"data": users})
}

func GetUser(c *gin.Context, db *gorm.DB) {
	var user models.User

	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func UpdateUser(c *gin.Context, db *gorm.DB) {

	var user models.User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var Input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{
		"username": Input.Username,
		"email":    Input.Email,
	}

	if Input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(Input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{"error": "failed to hash the new password"})
			return
		}
		updates["password_hash"] = string(passwordHash)
	}

	db.Model(&user).Updates(updates)
	db.First(&user, id)
	c.JSON(200, gin.H{"message": "user updated successfully", "data": user})
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	id := c.Param("id")

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"errors": "User not found"})
		return
	}

	db.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func GetUserProfile(c *gin.Context, db *gorm.DB) {
	var user models.User
	id, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	response := gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	}

	c.JSON(200, gin.H{"message": "sucess", "data": response})
}

func UpdateUserProfile(c *gin.Context, db *gorm.DB) {
	var user models.User
	var updatedUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	id, exist := c.Get("user_id")

	if !exist {
		c.JSON(401, gin.H{"errors": "Unauthorized"})
		return
	}

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"errors": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if updatedUser.Username != "" {

		user.Username = updatedUser.Username
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
	}

	response := gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"message":  "Profile updated successfully",
	}

	c.JSON(200, response)
}

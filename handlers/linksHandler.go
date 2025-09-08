package handlers

import (
	"strconv"
	"strings"
	"url_shortener/helpers"
	"url_shortener/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLink(c *gin.Context, db *gorm.DB) {
	var Input struct {
		OriginalUrl string `json:"url"`
	}
	var link models.Link

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unothaurized"})
		return
	}
	userIDFloat := userID.(float64)

	userIDVal := uint(userIDFloat)

	if !strings.HasPrefix(Input.OriginalUrl, "http://") && !strings.HasPrefix(Input.OriginalUrl, "https://") {
		c.JSON(400, gin.H{"error": "invalid link format"})
		return
	}

	link = models.Link{
		OriginalUrl: Input.OriginalUrl,
		UserID:      userIDVal,
	}
	if err := db.Create(&link).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to created Link"})
		return
	}

	link.ShortCode = helpers.GenerateCode(strconv.Itoa(int(link.ID)))

	if err := db.Save(&link).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to save database"})
		return
	}
	c.JSON(200, "Link created successfully")
}

func DeleteLink(c *gin.Context, db *gorm.DB) {
	var link models.Link
	id := c.Param("id")
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// if err := db.First(&link, id).Error; err != nil {
	// 	c.JSON(404, gin.H{"error": "Link not found"})
	// 	return
	// }

	if err := db.Preload("User").First(&link, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	userIDFloat := userID.(float64)
	userIDVal := uint(userIDFloat)

	if link.User.Role != "admin" && link.UserID != userIDVal {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	db.Delete(&link)
	c.JSON(200, gin.H{"message": "Link deleted successfully"})

}

func ListUserLinks(c *gin.Context, db *gorm.DB) {
	var links []models.Link

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unothaurized"})
		return
	}

	userIDFloat := userID.(float64)

	userIDVal := uint(userIDFloat)

	db.Where("user_id = ?", userIDVal).Find(&links)

	c.JSON(200, gin.H{"message": "links retrieved successfully", "data": links})
}

func RedirectOriginalUrl(c *gin.Context, db *gorm.DB) {
	shortcode := c.Param("shortcode")
	var link models.Link

	if err := db.Where("short_code = ?", shortcode).First(&link).Error; err != nil {
		c.JSON(404, gin.H{"error": "Link Not Found"})
		return
	}
	link.Clicks += 1
	if err := db.Save(&link).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to update clicks"})
		return
	}
	OriginalUrl := link.OriginalUrl
	c.Redirect(302, OriginalUrl)
}

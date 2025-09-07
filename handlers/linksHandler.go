package handlers

import (
	"strconv"
	"url_shortener/helpers"
	"url_shortener/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLink(c *gin.Context, db gorm.DB) {
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

	link = models.Link{
		OriginalUrl: Input.OriginalUrl,
		UserID:      userIDVal,
	}
	link.ShortCode = helpers.GenerateCode(strconv.Itoa(int(link.ID)))

	if err := db.Create(&link).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to created Link"})
		return
	}

}

func DeleteLink(c *gin.Context, db *gorm.DB) {
	var link models.Link
	id := c.Param("id")

	if err := db.First(&link, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Link not found"})
		return
	}

	db.Delete(&link)
	c.JSON(200, gin.H{"message": "Link deleted successfully"})

}

func ListUserLinks(c *gin.Context, db *gorm.DB) {

}

package controllers

import (
	"go-gorm/database"
	"go-gorm/helpers"
	"go-gorm/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()

	var socialMedia []models.SocialMedia
	err := db.Preload("User").Find(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"socialMedia": socialMedia,
	})
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var socialMedia models.SocialMedia

	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	userID := uint(userData["id"].(float64))
	socialMedia.UserID = userID

	err := db.Create(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var socialMedia models.SocialMedia

	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid parameter",
		})
		return
	}

	userID := uint(userData["id"].(float64))
	socialMedia.UserID = userID
	socialMedia.ID = uint(socialMediaID)

	err = db.Model(&models.SocialMedia{}).Where("id = ?", socialMediaID).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaUrl: socialMedia.SocialMediaUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	var socialMedia models.SocialMedia
	result := db.First(&socialMedia, socialMediaID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Social Media not found",
		})
		return
	}

	err = db.Delete(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfuly deleted",
	})
}

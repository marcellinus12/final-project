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

func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()

	var photos []models.Photo
	query := `SELECT photos.*, users.*, comments.*
               FROM photos
               LEFT JOIN users ON photos.user_id = users.id
               LEFT JOIN comments ON photos.id = comments.photo_id`

	if err := db.Raw(query).Scan(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": photos,
	})
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var photo models.Photo

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	userID := uint(userData["id"].(float64))
	photo.UserID = userID

	err := db.Create(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  "",
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var photo models.Photo

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "invalid parameter",
		})
		return
	}

	userID := uint(userData["id"].(float64))
	photo.UserID = userID
	photo.ID = uint(photoID)

	err = db.Model(&models.Photo{}).Where("id = ?", photoID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  "",
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()

	photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	var photo models.Photo
	result := db.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Photo not found",
		})
		return
	}

	err = db.Delete(&photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfuly deleted",
	})
}

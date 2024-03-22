package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"unique;type:varchar(191)" json:"name" valid:"required~Your name of social media is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Your social media url of social media is required"`
	UserID         uint
	User           *User
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(sm)
	if err != nil {
		return err
	}

	return nil
}

// func (sm *SocialMedia) BeforeUpdate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(sm)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
	PhotoID   uint   `json:"photo_id"`
	Message   string `json:"message" valid:"required~Your message of comment is required" `
	User      *User
	Photo     *Photo
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}

// func (c *Comment) BeforeUpdate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(c)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

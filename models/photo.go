package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"type:varchar(191)" json:"title" valid:"required~Your title of photo is required"`
	Caption   string `json:"caption"`
	PhotoUrl  string `gorm:"" json:"photo_url" valid:"required~Your photo url is required"`
	UserID    uint   `json:"user_id"`
	User      *User
	Comments  []Comment
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}

// func (p *Photo) BeforeUpdate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

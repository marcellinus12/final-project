package models

import (
	"go-gorm/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint          `gorm:"primaryKey" json:"id"`
	Username     string        `gorm:"unique;type:varchar(191)" json:"username" valid:"required~Username is required"`
	Email        string        `gorm:"unique;type:varchar(191)" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password     string        `json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum of 6 characters"`
	Age          uint          `json:"age" valid:"required~Your age is required,range(9|200)~Age has to be greater than 9"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}

// func (u *User) BeforeUpdate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(u)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

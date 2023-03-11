package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"primaryKey;unique"`
	FirstName   string `json:"first_name" gorm:"not null" validate:"max=30"`
	LastName    string `json:"last_name" gorm:"not null" validate:"max=30"`
	Email       string `json:"email" gorm:"not null;email;unique"`
	PhoneNumber string `json:"phone" gorm:"type:varchar(20);unique"`
	Password    string `json:"password" gorm:"not null" validate:"password"`
	IsModerator bool   `json:"is_moderator" gorm:"default:false"`
}

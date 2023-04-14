package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `json:"-" gorm:"primaryKey"`
	Username    string `json:"username" gorm:"unique"`
	FirstName   string `json:"first_name" gorm:"not null" validate:"max=30"`
	LastName    string `json:"last_name" gorm:"not null" validate:"max=30"`
	Email       string `json:"email" gorm:"not null;email;unique"`
	PhoneNumber string `json:"phone" gorm:"type:varchar(20);unique"`
	Password    string `json:"password" gorm:"not null" validate:"password"`
	IsModerator bool   `json:"is_moderator" gorm:"default:false"`
}

type Post struct {
	gorm.Model
	Title       string
	Description string
}

var ErrUserNotFound = errors.New("user not found")

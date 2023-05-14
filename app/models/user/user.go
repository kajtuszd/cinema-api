package user

import (
	"errors"
)

type User struct {
	ID          uint   `json:"-" gorm:"primaryKey"`
	Username    string `json:"username" gorm:"not null" validate:"omitempty,unique_username"`
	FirstName   string `json:"first_name" gorm:"not null" validate:"max=30"`
	LastName    string `json:"last_name" gorm:"not null" validate:"max=30"`
	Email       string `json:"email" gorm:"not null;email" validate:"omitempty,unique_email"`
	PhoneNumber string `json:"phone" gorm:"type:varchar(20);unique" validate:"omitempty,unique_phone"`
	Password    string `json:"password" gorm:"not null" validate:"required,password"`
	IsModerator bool   `json:"is_moderator" gorm:"default:false"`
}

var ErrUserNotFound = errors.New("user not found")

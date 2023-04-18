package user

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(username string) (*User, error)
	GetAll() ([]User, error)
	entity.Repository
}

type userRepository struct {
	entity.Repository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Repository: entity.NewRepository(db),
		db:         db,
	}
}

func (r *userRepository) Save(entity interface{}) error {
	user := entity.(User)
	hash, _ := utils.HashPassword(user.Password)
	if !utils.CheckPasswordHash(user.Password, hash) {
		return utils.PasswordHashError
	}
	user.Password = hash
	return r.db.Create(&user).Error
}

func (r *userRepository) GetByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

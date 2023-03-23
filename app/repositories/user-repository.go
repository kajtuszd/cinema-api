package repositories

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models"
	"github.com/kajtuszd/cinema-api/app/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) error
	Update(user models.User) error
	Delete(user models.User) error
	GetByUsername(username string) (*models.User, error)
	GetAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Save(user models.User) error {
	hash, _ := utils.HashPassword(user.Password)
	if !utils.CheckPasswordHash(user.Password, hash) {
		return utils.PasswordHashError
	}
	user.Password = hash
	return r.db.Create(&user).Error
}

func (r *userRepository) Update(user models.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(user models.User) error {
	return r.db.Delete(&user).Error
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
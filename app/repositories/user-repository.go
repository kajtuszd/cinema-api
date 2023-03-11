package repositories

import (
	"github.com/kajtuszd/cinema-api/app/models"
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

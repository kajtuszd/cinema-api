package services

import (
	"github.com/kajtuszd/cinema-api/app/models"
	"github.com/kajtuszd/cinema-api/app/repositories"
	//"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(user models.User) error
	GetByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func New(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service *userService) CreateUser(user models.User) error {
	return service.userRepo.Save(user)
}

func (service *userService) UpdateUser(user models.User) error {
	return service.userRepo.Update(user)
}

func (service *userService) DeleteUser(user models.User) error {
	return service.userRepo.Delete(user)
}

func (service *userService) GetByUsername(username string) (*models.User, error) {
	return service.userRepo.GetByUsername(username)
}

func (service *userService) GetAllUsers() ([]models.User, error) {
	return service.userRepo.GetAll()
}

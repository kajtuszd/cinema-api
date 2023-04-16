package user

import (
	"github.com/kajtuszd/cinema-api/app/utils"
)

type UserService interface {
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(user User) error
	GetByUsername(username string) (*User, error)
	GetAllUsers() ([]User, error)
	CheckLogin(username, password string) (string, error)
}

type userService struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service *userService) CreateUser(user User) error {
	return service.userRepo.Save(user)
}

func (service *userService) UpdateUser(user User) error {
	return service.userRepo.Update(user)
}

func (service *userService) DeleteUser(user User) error {
	return service.userRepo.Delete(user)
}

func (service *userService) GetByUsername(username string) (*User, error) {
	return service.userRepo.GetByUsername(username)
}

func (service *userService) GetAllUsers() ([]User, error) {
	return service.userRepo.GetAll()
}

func (service *userService) CheckLogin(username, password string) (string, error) {
	user, err := service.GetByUsername(username)
	var token string
	if err != nil {
		return token, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return token, utils.PasswordMismatchError
	}
	token, err = utils.GenerateToken(user)
	return token, err
}

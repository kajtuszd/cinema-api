package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/utils"
)

type UserService interface {
	GetByUsername(username string) (*User, error)
	GetAllUsers() ([]User, error)
	CheckLogin(username, password string) (string, error)
	UniquePhoneValidator(fl validator.FieldLevel) bool
	UniqueUsernameValidator(fl validator.FieldLevel) bool
	UniqueEmailValidator(fl validator.FieldLevel) bool
	entity.Service
}

type userService struct {
	userRepo UserRepository
	entity.Service
}

func NewService(userRepo UserRepository) UserService {
	return &userService{
		Service:  entity.NewService(userRepo),
		userRepo: userRepo,
	}
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
	token, err = GenerateToken(user)
	return token, err
}

func (service *userService) UniquePhoneValidator(fl validator.FieldLevel) bool {
	return service.userRepo.UniquePhoneValidator(fl)
}

func (service *userService) UniqueUsernameValidator(fl validator.FieldLevel) bool {
	return service.userRepo.UniqueUsernameValidator(fl)
}

func (service *userService) UniqueEmailValidator(fl validator.FieldLevel) bool {
	return service.userRepo.UniqueEmailValidator(fl)
}

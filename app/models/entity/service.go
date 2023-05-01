package entity

import (
	"gorm.io/gorm"
)

type Service interface {
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
}

type service struct {
	db         *gorm.DB
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(entity interface{}) error {
	return s.repository.Save(entity)
}

func (s *service) Update(entity interface{}) error {
	return s.repository.Update(entity)
}

func (s *service) Delete(entity interface{}) error {
	return s.repository.Delete(entity)
}

package entity

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(entity interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(entity interface{}) error {
	return r.db.Create(entity).Error
}

func (r *repository) Update(entity interface{}) error {
	return r.db.Save(entity).Error
}

func (r *repository) Delete(entity interface{}) error {
	return r.db.Delete(entity).Error
}

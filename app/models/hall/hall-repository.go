package hall

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"gorm.io/gorm"
)

type HallRepository interface {
	GetByID(id string) (*Hall, error)
	GetAll() ([]Hall, error)
	entity.Repository
}

type hallRepository struct {
	db *gorm.DB
	entity.Repository
}

func NewRepository(db *gorm.DB) HallRepository {
	return &hallRepository{
		Repository: entity.NewRepository(db),
		db:         db,
	}
}

func (r *hallRepository) GetByID(id string) (*Hall, error) {
	var hall Hall
	err := r.db.Where("id = ?", id).First(&hall).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrHallNotFound
		}
		return nil, err
	}
	return &hall, nil
}

func (r *hallRepository) GetAll() ([]Hall, error) {
	var halls []Hall
	if err := r.db.Find(&halls).Error; err != nil {
		return nil, err
	}
	return halls, nil
}

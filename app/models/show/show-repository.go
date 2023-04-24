package show

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"gorm.io/gorm"
)

type ShowRepository interface {
	GetByID(id string) (*Show, error)
	GetAll() ([]Show, error)
	entity.Repository
}

type showRepository struct {
	db *gorm.DB
	entity.Repository
}

func NewRepository(db *gorm.DB) ShowRepository {
	return &showRepository{
		Repository: entity.NewRepository(db),
		db:         db,
	}
}

func (s *showRepository) GetByID(id string) (*Show, error) {
	var show Show
	err := s.db.Where("id = ?", id).Preload("Movie").Preload("Hall").First(&show).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrShowNotFound
		}
		return nil, err
	}
	return &show, nil
}

func (s *showRepository) GetAll() ([]Show, error) {
	var shows []Show
	if err := s.db.Preload("Movie").Preload("Hall").Find(&shows).Error; err != nil {
		return nil, err
	}
	return shows, nil
}

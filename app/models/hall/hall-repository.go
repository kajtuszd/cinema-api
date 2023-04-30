package hall

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"gorm.io/gorm"
)

type HallRepository interface {
	GetByID(id string) (*Hall, error)
	GetByNumber(hallNumber string) (*Hall, error)
	ExistsByNumber(number string) bool
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

func (r *hallRepository) GetByNumber(hallNumber string) (*Hall, error) {
	var hall Hall
	err := r.db.Where("hall_number = ?", hallNumber).First(&hall).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrHallNotFound
		}
		return nil, err
	}
	return &hall, nil
}

func (r *hallRepository) ExistsByNumber(number string) bool {
	var count int64
	if err := r.db.Model(&Hall{}).Where("hall_number = ?", number).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (r *hallRepository) GetAll() ([]Hall, error) {
	var halls []Hall
	if err := r.db.Find(&halls).Error; err != nil {
		return nil, err
	}
	return halls, nil
}

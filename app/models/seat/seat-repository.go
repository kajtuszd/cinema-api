package seat

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"gorm.io/gorm"
)

type SeatRepository interface {
	GetByID(id string) (*Seat, error)
	GetAll() ([]Seat, error)
	entity.Repository
}

type seatRepository struct {
	db *gorm.DB
	entity.Repository
}

func NewRepository(db *gorm.DB) SeatRepository {
	return &seatRepository{
		Repository: entity.NewRepository(db),
		db:         db,
	}
}

func (s *seatRepository) GetByID(id string) (*Seat, error) {
	var seat Seat
	err := s.db.Where("id = ?", id).
		Preload("Show").
		Preload("Show.Hall").
		Preload("Show.Movie").
		First(&seat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSeatNotFound
		}
		return nil, err
	}
	return &seat, nil
}

func (s *seatRepository) GetAll() ([]Seat, error) {
	var seats []Seat
	if err := s.db.
		Preload("Show").
		Preload("Show.Hall").
		Preload("Show.Movie").
		Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

package movie

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	GetByID(id string) (*Movie, error)
	GetAll() ([]Movie, error)
	entity.Repository
}

type movieRepository struct {
	entity.Repository
	db *gorm.DB
}

func NewRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{
		Repository: entity.NewRepository(db),
		db:         db,
	}
}

func (r *movieRepository) GetByID(id string) (*Movie, error) {
	var movie Movie
	err := r.db.Where("id = ?", id).First(&movie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMovieNotFound
		}
		return nil, err
	}
	return &movie, nil
}

func (r *movieRepository) GetAll() ([]Movie, error) {
	var movies []Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

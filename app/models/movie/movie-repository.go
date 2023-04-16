package movie

import (
	"errors"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(movie Movie) error
	Update(movie Movie) error
	Delete(movie Movie) error
	GetByID(id string) (*Movie, error)
	GetAll() ([]Movie, error)
}

type movieRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{
		db: db,
	}
}

func (r *movieRepository) Save(movie Movie) error {
	return r.db.Create(&movie).Error
}

func (r *movieRepository) Update(movie Movie) error {
	return r.db.Save(&movie).Error
}

func (r *movieRepository) Delete(movie Movie) error {
	return r.db.Delete(&movie).Error
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

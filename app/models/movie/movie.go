package movie

import "errors"

type Movie struct {
	ID             uint   `json:"-" gorm:"primaryKey"`
	Title          string `json:"title" gorm:"not null" validate:"max=50"`
	ProductionYear int    `json:"production_year" validate:"min=1950,max=2023"`
	TimeInMinutes  int    `json:"time_in_minutes" validate:"min=0,max=300"`
	Description    string `json:"description" validate:"max=200"`
}

var ErrMovieNotFound = errors.New("movie not found")

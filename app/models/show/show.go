package show

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/hall"
	"github.com/kajtuszd/cinema-api/app/models/movie"
	"time"
)

type Show struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	MovieID   uint
	Movie     movie.Movie `json:"movie" gorm:"foreignkey:MovieID"`
	HallID    uint
	Hall      hall.Hall `json:"hall" gorm:"foreignkey:HallID"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
}

var ErrShowNotFound = errors.New("show not found")

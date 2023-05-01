package hall

import "errors"

type Hall struct {
	ID            uint `json:"-" gorm:"primaryKey"`
	NumberOfSeats int  `json:"number_of_seats" gorm:"not null" validate:"max=300,min=1"`
	HallNumber    int  `json:"hall_number" gorm:"not null;unique" validate:"max=100,min=1"`
}

var ErrHallNotFound = errors.New("hall not found")
var ErrBadHallNumber = errors.New("hall with given number already exists")

package seat

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/show"
)

type Seat struct {
	ID     uint `json:"-" gorm:"primaryKey"`
	ShowID uint
	Show   show.Show `json:"show" gorm:"foreignkey:ShowID"`
	State  string    `json:"state" gorm:"not null"`
}

type SeatState string

const (
	SeatStateReserved  SeatState = "reserved"
	SeatStateAvailable SeatState = "available"
)

var ErrSeatNotFound = errors.New("seat not found")

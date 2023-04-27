package ticket

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/seat"
	"github.com/kajtuszd/cinema-api/app/models/user"
)

type Ticket struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint
	Owner  user.User `json:"owner" gorm:"foreignkey:UserID"`
	SeatID uint
	Seat   seat.Seat `json:"seat" gorm:"foreignkey:SeatID"`
	Price  int       `json:"price" gorm:"not null" validate:"min=0"`
}

var ErrTicketNotFound = errors.New("ticket not found")

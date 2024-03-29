package seat

import (
	"github.com/google/uuid"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/models/show"
	"strconv"
)

type SeatService interface {
	GetByID(id string) (*Seat, error)
	GetSeatsForShow(showID string) ([]Seat, error)
	CreateSeatsForShow(show show.Show) error
	DeleteSeatsForShow(showID string) error
	SetSeatReserved(seat *Seat) error
	SetSeatAvailable(seat *Seat) error
	IsSeatReserved(seat *Seat) bool
	entity.Service
}

type seatService struct {
	seatRepo SeatRepository
	entity.Service
}

func NewService(seatRepo SeatRepository) SeatService {
	return &seatService{
		Service:  entity.NewService(seatRepo),
		seatRepo: seatRepo,
	}
}

func (service *seatService) GetByID(id string) (*Seat, error) {
	return service.seatRepo.GetByID(id)
}

func (service *seatService) GetSeatsForShow(showID string) ([]Seat, error) {
	showSeats, err := service.seatRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var filteredSeats []Seat
	for _, s := range showSeats {
		if strconv.Itoa(int(s.ShowID)) == showID {
			filteredSeats = append(filteredSeats, s)
		}
	}
	return showSeats, err
}

func (service *seatService) CreateSeatsForShow(show show.Show) error {
	for i := 0; i < show.Hall.NumberOfSeats; i++ {
		seat := &Seat{
			ID:     uint(uuid.New().ID()),
			ShowID: show.ID,
			State:  string(SeatStateAvailable),
		}
		if err := service.Create(seat); err != nil {
			return err
		}
	}
	return nil
}

func (service *seatService) DeleteSeatsForShow(showID string) error {
	seats, err := service.GetSeatsForShow(showID)
	if err != nil {
		return err
	}
	for _, seat := range seats {
		if err = service.Delete(seat); err != nil {
			return err
		}
	}
	return nil
}

func (service *seatService) SetSeatReserved(seat *Seat) error {
	seat.State = string(SeatStateReserved)
	return service.Update(seat)
}

func (service *seatService) SetSeatAvailable(seat *Seat) error {
	seat.State = string(SeatStateAvailable)
	return service.Update(seat)
}

func (service *seatService) IsSeatReserved(seat *Seat) bool {
	return seat.State == string(SeatStateReserved)
}

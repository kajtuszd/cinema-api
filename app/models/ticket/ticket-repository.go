package ticket

import (
	"errors"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/models/user"
	"gorm.io/gorm"
)

type TicketRepository interface {
	GetByID(id string) (*Ticket, error)
	GetAll() ([]Ticket, error)
	GetTicketsByUser(user user.User) ([]Ticket, error)
	entity.Repository
}

type ticketRepository struct {
	db *gorm.DB
	entity.Repository
}

func NewRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		db:         db,
		Repository: entity.NewRepository(db),
	}
}

func (t *ticketRepository) GetByID(id string) (*Ticket, error) {
	var ticket Ticket
	err := t.db.Where("id = ?", id).
		Preload("Seat").
		Preload("Seat.Show").
		Preload("Seat.Show.Hall").
		Preload("Seat.Show.Movie").
		Preload("Owner").
		First(&ticket).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTicketNotFound
		}
		return nil, err
	}
	return &ticket, nil
}

func (t *ticketRepository) GetAll() ([]Ticket, error) {
	var tickets []Ticket
	if err := t.db.Preload("Seat").
		Preload("Seat.Show").
		Preload("Seat.Show.Hall").
		Preload("Seat.Show.Movie").
		Preload("Owner").
		Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (t *ticketRepository) GetTicketsByUser(user user.User) ([]Ticket, error) {
	var tickets []Ticket
	err := t.db.Where("user_id = ?", user.ID).
		Preload("Seat").
		Preload("Seat.Show").
		Preload("Seat.Show.Hall").
		Preload("Seat.Show.Movie").
		Preload("Owner").
		Find(&tickets).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTicketNotFound
		}
		return nil, err
	}
	return tickets, nil
}

package ticket

import (
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/models/user"
)

type TicketService interface {
	GetByID(id string) (*Ticket, error)
	GetAllTickets() ([]Ticket, error)
	GetTicketsByUser(user user.User) ([]Ticket, error)
	entity.Service
}

type ticketService struct {
	ticketRepo TicketRepository
	entity.Service
}

func NewService(ticketRepo TicketRepository) TicketService {
	return &ticketService{
		Service:    entity.NewService(ticketRepo),
		ticketRepo: ticketRepo,
	}
}

func (service *ticketService) GetByID(id string) (*Ticket, error) {
	return service.ticketRepo.GetByID(id)
}

func (service *ticketService) GetAllTickets() ([]Ticket, error) {
	return service.ticketRepo.GetAll()
}

func (service *ticketService) GetTicketsByUser(user user.User) ([]Ticket, error) {
	return service.ticketRepo.GetTicketsByUser(user)
}

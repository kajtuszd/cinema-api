package show

import "github.com/kajtuszd/cinema-api/app/models/entity"

type ShowService interface {
	GetByID(id string) (*Show, error)
	GetAllShows() ([]Show, error)
	entity.Service
}

type showService struct {
	showRepo ShowRepository
	entity.Service
}

func NewService(showRepo ShowRepository) ShowService {
	return &showService{
		Service:  entity.NewService(showRepo),
		showRepo: showRepo,
	}
}

func (s *showService) GetByID(id string) (*Show, error) {
	return s.showRepo.GetByID(id)
}

func (s *showService) GetAllShows() ([]Show, error) {
	return s.showRepo.GetAll()
}

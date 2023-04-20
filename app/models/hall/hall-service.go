package hall

import (
	"github.com/kajtuszd/cinema-api/app/models/entity"
)

type HallService interface {
	GetByID(id string) (*Hall, error)
	GetAllHalls() ([]Hall, error)
	entity.Service
}

type hallService struct {
	hallRepo HallRepository
	entity.Service
}

func NewService(hallRepo HallRepository) HallService {
	return &hallService{
		Service:  entity.NewService(hallRepo),
		hallRepo: hallRepo,
	}
}

func (service *hallService) GetByID(id string) (*Hall, error) {
	return service.hallRepo.GetByID(id)
}

func (service *hallService) GetAllHalls() ([]Hall, error) {
	return service.hallRepo.GetAll()
}

package movie

import "github.com/kajtuszd/cinema-api/app/models/entity"

type MovieService interface {
	GetByID(id string) (*Movie, error)
	GetAllMovies() ([]Movie, error)
	entity.Service
}

type movieService struct {
	movieRepo MovieRepository
	entity.Service
}

func NewService(movieRepo MovieRepository) MovieService {
	return &movieService{
		Service:   entity.NewService(movieRepo),
		movieRepo: movieRepo,
	}
}

func (service *movieService) GetByID(id string) (*Movie, error) {
	return service.movieRepo.GetByID(id)
}

func (service *movieService) GetAllMovies() ([]Movie, error) {
	return service.movieRepo.GetAll()
}

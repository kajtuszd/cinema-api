package movie

type MovieService interface {
	CreateMovie(movie *Movie) error
	UpdateMovie(movie *Movie) error
	DeleteMovie(movie *Movie) error
	GetByID(id string) (*Movie, error)
	GetAllMovies() ([]Movie, error)
}

type movieService struct {
	movieRepo MovieRepository
}

func NewService(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (service *movieService) CreateMovie(movie *Movie) error {
	return service.movieRepo.Save(movie)
}

func (service *movieService) UpdateMovie(movie *Movie) error {
	return service.movieRepo.Update(movie)
}

func (service *movieService) DeleteMovie(movie *Movie) error {
	return service.movieRepo.Delete(movie)
}

func (service *movieService) GetByID(id string) (*Movie, error) {
	return service.movieRepo.GetByID(id)
}

func (service *movieService) GetAllMovies() ([]Movie, error) {
	return service.movieRepo.GetAll()
}

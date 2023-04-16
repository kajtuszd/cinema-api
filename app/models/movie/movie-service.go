package movie

type MovieService interface {
	CreateMovie(movie Movie) error
	UpdateMovie(movie Movie) error
	DeleteMovie(movie Movie) error
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

func (service *movieService) CreateMovie(user Movie) error {
	return service.movieRepo.Save(user)
}

func (service *movieService) UpdateMovie(user Movie) error {
	return service.movieRepo.Update(user)
}

func (service *movieService) DeleteMovie(user Movie) error {
	return service.movieRepo.Delete(user)
}

func (service *movieService) GetByID(id string) (*Movie, error) {
	return service.movieRepo.GetByID(id)
}

func (service *movieService) GetAllMovies() ([]Movie, error) {
	return service.movieRepo.GetAll()
}

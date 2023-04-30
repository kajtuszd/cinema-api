package movie

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"net/http"
)

type MovieController interface {
	GetMovie(ctx *gin.Context)
	GetAllMovies(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	entity.Controller
}

type movieController struct {
	movieService MovieService
	validator    *validator.Validate
	entity.Controller
}

func NewController(service MovieService) MovieController {
	return &movieController{
		movieService: service,
		Controller:   entity.NewController(),
	}
}

func (c *movieController) GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := c.movieService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrMovieNotFound); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": movie})
}

func (c *movieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.movieService.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": movies})
}

func (c *movieController) CreateMovie(ctx *gin.Context) {
	var movie Movie
	if !c.ValidateRequest(ctx, &movie, nil) {
		return
	}
	if err := c.movieService.Create(&movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"movie": movie})
}

func (c *movieController) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := c.movieService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrMovieNotFound); err != nil {
		return
	}
	if err = c.movieService.Delete(movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}

func (c *movieController) UpdateMovie(ctx *gin.Context) {
	var updatedMovie Movie
	id := ctx.Param("id")
	movie, err := c.movieService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrMovieNotFound); err != nil {
		return
	}
	if !c.ValidateRequest(ctx, &updatedMovie, nil) {
		return
	}
	movie.Title = updatedMovie.Title
	movie.Description = updatedMovie.Description
	movie.TimeInMinutes = updatedMovie.TimeInMinutes
	movie.ProductionYear = updatedMovie.ProductionYear
	if err = c.movieService.Update(movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie updated successfully"})
}

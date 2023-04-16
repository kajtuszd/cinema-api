package movie

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieController interface {
	GetMovie(ctx *gin.Context)
	GetAllMovies(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	handleError(ctx *gin.Context, err error) error
}

type movieController struct {
	movieService MovieService
}

func NewController(service MovieService) MovieController {
	return &movieController{
		movieService: service,
	}
}

func (c *movieController) handleError(ctx *gin.Context, err error) error {
	if err != nil {
		if errors.Is(err, ErrMovieNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": ErrMovieNotFound.Error()})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (c *movieController) GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.movieService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *movieController) GetAllMovies(ctx *gin.Context) {
	users, err := c.movieService.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *movieController) CreateMovie(ctx *gin.Context) {
	var movie Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.movieService.CreateMovie(movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (c *movieController) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.movieService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	if err = c.movieService.DeleteMovie(*user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}

func (c *movieController) UpdateMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.movieService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.movieService.UpdateMovie(*user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie updated successfully"})
}

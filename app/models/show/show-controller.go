package show

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/models/hall"
	"github.com/kajtuszd/cinema-api/app/models/movie"
	"net/http"
	"time"
)

type ShowController interface {
	GetShow(ctx *gin.Context)
	GetAllShows(ctx *gin.Context)
	CreateShow(ctx *gin.Context)
	DeleteShow(ctx *gin.Context)
	UpdateShow(ctx *gin.Context)
	entity.Controller
}

type showController struct {
	showService  ShowService
	movieService movie.MovieService
	hallService  hall.HallService
	validator    *validator.Validate
	entity.Controller
}

func NewController(s ShowService, m movie.MovieService, h hall.HallService) ShowController {
	v := validator.New()
	return &showController{
		showService:  s,
		movieService: m,
		hallService:  h,
		validator:    v,
		Controller:   entity.NewController(),
	}
}

type ShowInput struct {
	MovieID   uint      `json:"movie_id"`
	HallID    uint      `json:"hall_id"`
	StartTime time.Time `json:"start_time"`
}

func (c *showController) GetShow(ctx *gin.Context) {
	id := ctx.Param("id")
	show, err := c.showService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrShowNotFound); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": show})
}

func (c *showController) GetAllShows(ctx *gin.Context) {
	shows, err := c.showService.GetAllShows()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shows})
}

func (c *showController) CreateShow(ctx *gin.Context) {
	var input ShowInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movieObj, err := c.movieService.GetByID(fmt.Sprintf("%d", input.MovieID))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	hallObj, err := c.hallService.GetByID(fmt.Sprintf("%d", input.HallID))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	show := &Show{
		ID:        uint(uuid.New().ID()),
		MovieID:   input.MovieID,
		HallID:    input.HallID,
		Movie:     *movieObj,
		Hall:      *hallObj,
		StartTime: input.StartTime,
	}
	if err := c.validator.Struct(show); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.showService.Create(show); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"show": show})
}

func (c *showController) DeleteShow(ctx *gin.Context) {
	id := ctx.Param("id")
	show, err := c.showService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrShowNotFound); err != nil {
		return
	}
	if err = c.showService.Delete(show); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "show deleted successfully"})
}

func (c *showController) UpdateShow(ctx *gin.Context) {
	var input ShowInput
	id := ctx.Param("id")
	show, err := c.showService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrShowNotFound); err != nil {
		return
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.MovieID > 0 {
		movieObj, err := c.movieService.GetByID(fmt.Sprintf("%d", input.MovieID))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		show.MovieID = input.MovieID
		show.Movie = *movieObj
	}
	if input.HallID > 0 {
		hallObj, err := c.hallService.GetByID(fmt.Sprintf("%d", input.HallID))
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		show.HallID = input.HallID
		show.Hall = *hallObj
	}
	if !input.StartTime.IsZero() {
		show.StartTime = input.StartTime
	}
	if err := c.validator.Struct(show); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.showService.Update(show); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "show updated successfully"})
}

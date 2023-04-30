package seat

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/models/show"
	"net/http"
)

type SeatController interface {
	GetSeat(ctx *gin.Context)
	GetSeatsForShow(ctx *gin.Context)
	CreateSeatsForShow(ctx *gin.Context)
	DeleteSeatsForShow(ctx *gin.Context)
	entity.Controller
}

type seatController struct {
	seatService SeatService
	showService show.ShowService
	validator   *validator.Validate
	entity.Controller
}

func NewController(seatServ SeatService, showServ show.ShowService) SeatController {
	v := validator.New()
	return &seatController{
		seatService: seatServ,
		showService: showServ,
		validator:   v,
		Controller:  entity.NewController(),
	}
}

type SeatInput struct {
	ShowID uint   `json:"movie_id"`
	State  string `json:"state"`
}

func (c *seatController) GetSeat(ctx *gin.Context) {
	id := ctx.Param("id")
	seat, err := c.seatService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrSeatNotFound); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": seat})
}

func (c *seatController) GetSeatsForShow(ctx *gin.Context) {
	showID := ctx.Param("id")
	seats, err := c.seatService.GetSeatsForShow(showID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": seats})
}

func (c *seatController) CreateSeatsForShow(ctx *gin.Context) {
	showID := ctx.Param("id")
	showObj, err := c.showService.GetByID(showID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seats, err := c.seatService.GetSeatsForShow(showID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(seats) > 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "seats are already created"})
		return
	}
	err = c.seatService.CreateSeatsForShow(*showObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully created seats"})
}

func (c *seatController) DeleteSeatsForShow(ctx *gin.Context) {
	showID := ctx.Param("id")
	if err := c.seatService.DeleteSeatsForShow(showID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted seats"})
}

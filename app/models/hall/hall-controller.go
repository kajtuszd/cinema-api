package hall

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"net/http"
	"strconv"
)

type HallController interface {
	GetHall(ctx *gin.Context)
	GetAllHalls(ctx *gin.Context)
	CreateHall(ctx *gin.Context)
	DeleteHall(ctx *gin.Context)
	UpdateHall(ctx *gin.Context)
	entity.Controller
}

type hallController struct {
	hallService HallService
	entity.Controller
}

func NewController(service HallService) HallController {
	return &hallController{
		hallService: service,
		Controller:  entity.NewController(),
	}
}

func (c *hallController) GetHall(ctx *gin.Context) {
	id := ctx.Param("id")
	hall, err := c.hallService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrHallNotFound); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": hall})
}

func (c *hallController) GetAllHalls(ctx *gin.Context) {
	halls, err := c.hallService.GetAllHalls()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": halls})
}

func (c *hallController) CreateHall(ctx *gin.Context) {
	var hall Hall
	if !c.ValidateRequest(ctx, &hall, func(req interface{}) error {
		if c.hallService.ExistsByNumber(strconv.Itoa(hall.HallNumber)) {
			return ErrBadHallNumber
		}
		return nil
	}) {
		return
	}
	if err := c.hallService.Create(&hall); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"hall": hall})
}

func (c *hallController) DeleteHall(ctx *gin.Context) {
	id := ctx.Param("id")
	hall, err := c.hallService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrHallNotFound); err != nil {
		return
	}
	if err = c.hallService.Delete(hall); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "hall deleted successfully"})
}

type hallInput struct {
	NumberOfSeats int `json:"number_of_seats" gorm:"not null" validate:"max=300,min=1"`
	HallNumber    int `json:"hall_number" gorm:"not null;unique" validate:"max=100,min=1"`
}

func (c *hallController) UpdateHall(ctx *gin.Context) {
	id := ctx.Param("id")
	hall, err := c.hallService.GetByID(id)
	if err = c.HandleError(ctx, err, ErrHallNotFound); err != nil {
		return
	}
	var hallInput hallInput
	if !c.ValidateRequest(ctx, &hallInput, func(req interface{}) error {
		if c.hallService.ExistsByNumber(strconv.Itoa(hallInput.HallNumber)) {
			oldHall, err := c.hallService.GetByNumber(strconv.Itoa(hallInput.HallNumber))
			if err != nil {
				return err
			}
			if oldHall.ID != hall.ID {
				return ErrBadHallNumber
			}
		}
		return nil
	}) {
		return
	}
	hall.HallNumber = hallInput.HallNumber
	hall.NumberOfSeats = hallInput.NumberOfSeats
	if err = c.hallService.Update(hall); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "hall updated successfully"})
}

package hall

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HallController interface {
	GetHall(ctx *gin.Context)
	GetAllHalls(ctx *gin.Context)
	CreateHall(ctx *gin.Context)
	DeleteHall(ctx *gin.Context)
	UpdateHall(ctx *gin.Context)
	handleError(ctx *gin.Context, err error) error
}

type hallController struct {
	hallService HallService
	validator   *validator.Validate
}

func NewController(service HallService) HallController {
	v := validator.New()
	return &hallController{
		hallService: service,
		validator:   v,
	}
}

func (c *hallController) handleError(ctx *gin.Context, err error) error {
	if err != nil {
		if errors.Is(err, ErrHallNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": ErrHallNotFound.Error()})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (c *hallController) GetHall(ctx *gin.Context) {
	id := ctx.Param("id")
	hall, err := c.hallService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
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
	if err := ctx.ShouldBindJSON(&hall); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.validator.Struct(hall); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	if err = c.hallService.Delete(hall); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "hall deleted successfully"})
}

func (c *hallController) UpdateHall(ctx *gin.Context) {
	id := ctx.Param("id")
	hall, err := c.hallService.GetByID(id)
	if err = c.handleError(ctx, err); err != nil {
		return
	}
	if err := ctx.ShouldBindJSON(&hall); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.validator.Struct(hall); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.hallService.Update(hall); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "hall updated successfully"})
}

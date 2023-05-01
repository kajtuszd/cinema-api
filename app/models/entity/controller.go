package entity

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

type Controller interface {
	HandleError(ctx *gin.Context, err error, notFoundError error) error
	ValidateRequest(ctx *gin.Context, request interface{}, extraFunc func(interface{}) error) bool
}

type controller struct {
	validator *validator.Validate
}

func NewController() Controller {
	return &controller{
		validator: validator.New(),
	}
}

func (c *controller) HandleError(ctx *gin.Context, err error, notFoundError error) error {
	if err != nil {
		if errors.Is(err, notFoundError) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": notFoundError.Error()})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (c *controller) ValidateRequest(ctx *gin.Context, request interface{}, extraFunc func(interface{}) error) bool {
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	if extraFunc != nil && reflect.TypeOf(extraFunc).Kind() == reflect.Func {
		if err := extraFunc(request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return false
		}
	}
	if err := c.validator.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

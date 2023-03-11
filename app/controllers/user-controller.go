package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/services"
	"net/http"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func New(service services.UserService) UserController {
	return &userController{
		userService: service,
	}
}

func (c *userController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

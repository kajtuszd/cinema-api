package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/models"
	"github.com/kajtuszd/cinema-api/app/services"
	"net/http"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	handleUserError(ctx *gin.Context, err error) error
}

type userController struct {
	userService services.UserService
}

func New(service services.UserService) UserController {
	return &userController{
		userService: service,
	}
}

func (c *userController) handleUserError(ctx *gin.Context, err error) error {
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": models.ErrUserNotFound.Error()})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func (c *userController) GetUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(username)
	if err = c.handleUserError(ctx, err); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *userController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.userService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(username)
	if err = c.handleUserError(ctx, err); err != nil {
		return
	}
	if err = c.userService.DeleteUser(*user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(username)
	fmt.Println(user)
	if err = c.handleUserError(ctx, err); err != nil {
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)
	if err = c.userService.UpdateUser(*user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

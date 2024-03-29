package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kajtuszd/cinema-api/app/models/entity"
	"github.com/kajtuszd/cinema-api/app/validators"
	"net/http"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	Validate(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
	entity.Controller
}

var validate *validator.Validate

type userController struct {
	userService UserService
	entity.Controller
}

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" validate:"password"`
}

func NewController(service UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("password", validators.PasswordValidator)
	validate.RegisterValidation("unique_username", service.UniqueUsernameValidator)
	validate.RegisterValidation("unique_phone", service.UniquePhoneValidator)
	validate.RegisterValidation("unique_email", service.UniqueEmailValidator)
	return &userController{
		userService: service,
		Controller:  entity.NewController(),
	}
}

func (c *userController) GetUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(username)
	if err = c.HandleError(ctx, err, ErrUserNotFound); err != nil {
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
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.userService.Create(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (c *userController) LoginUser(ctx *gin.Context) {
	var input LoginForm
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.userService.CheckLogin(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user logged successfully",
		"token":   token})
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(username)
	if err = c.HandleError(ctx, err, ErrUserNotFound); err != nil {
		return
	}
	if err = c.userService.Delete(*user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(username)
	email := user.Email
	phone := user.PhoneNumber
	if err = c.HandleError(ctx, err, ErrUserNotFound); err != nil {
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" {
		user.Email = email
	}
	if user.PhoneNumber == "" {
		user.PhoneNumber = phone
	}
	if user.Username == "" {
		user.Username = username
	}
	if err = c.userService.Update(*user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (c *userController) Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (c *userController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user logged out successfully",
	})
}

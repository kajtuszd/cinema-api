package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/controllers"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/repositories"
	"github.com/kajtuszd/cinema-api/app/services"
)

func InitializeRoutes(r *gin.Engine, db *database.GormDatabase) {
	userRepo := repositories.New(db.DB())
	userService := services.New(userRepo)
	userController := controllers.New(userService)
	userRoutes := r.Group("/users/")
	{
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.GET(":username", userController.GetUser)
		userRoutes.POST("", userController.CreateUser)
		userRoutes.DELETE(":username", userController.DeleteUser)
		userRoutes.PUT(":username", userController.UpdateUser)
		userRoutes.PATCH(":username", userController.UpdateUser)
	}
}
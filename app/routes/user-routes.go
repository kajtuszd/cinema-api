package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/middleware"
	"github.com/kajtuszd/cinema-api/app/models/user"
)

func InitializeUserRoutes(r *gin.Engine, db *database.GormDatabase) {
	userService := user.NewService(user.NewRepository(db.DB()))
	userController := user.NewController(userService)
	authRoutes := r.Group("/auth/")
	{
		authRoutes.POST("login/", userController.LoginUser)
		authRoutes.GET("logout/", middleware.JWTAuth(db), userController.LogoutUser)
		authRoutes.GET("validate/", middleware.JWTAuth(db), userController.Validate)
	}
	userRoutes := r.Group("/users/")
	{
		userRoutes.GET("", middleware.JWTAuth(db), middleware.Moderator(), userController.GetAllUsers)
		userRoutes.GET(":username", middleware.JWTAuth(db), middleware.AccountOwnerOrModerator(), userController.GetUser)
		userRoutes.POST("", userController.CreateUser)
		userRoutes.DELETE(":username", middleware.JWTAuth(db), middleware.AccountOwnerOrModerator(), userController.DeleteUser)
		userRoutes.PUT(":username", middleware.JWTAuth(db), middleware.AccountOwnerOrModerator(), userController.UpdateUser)
	}
}

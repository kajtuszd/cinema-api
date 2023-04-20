package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/models/hall"
)

func InitializeHallRoutes(r *gin.Engine, db *database.GormDatabase) {
	hallRepo := hall.NewRepository(db.DB())
	hallService := hall.NewService(hallRepo)
	hallController := hall.NewController(hallService)
	hallRoutes := r.Group("/halls/")
	{
		hallRoutes.GET("", hallController.GetAllHalls)
		hallRoutes.GET(":id", hallController.GetHall)
		hallRoutes.POST("", hallController.CreateHall)
		hallRoutes.DELETE(":id", hallController.DeleteHall)
		hallRoutes.PUT(":id", hallController.UpdateHall)
		hallRoutes.PATCH(":id", hallController.UpdateHall)
	}
}

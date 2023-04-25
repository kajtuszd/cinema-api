package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/models/hall"
	"github.com/kajtuszd/cinema-api/app/models/movie"
	"github.com/kajtuszd/cinema-api/app/models/seat"
	"github.com/kajtuszd/cinema-api/app/models/show"
)

func InitializeShowRoutes(r *gin.Engine, db *database.GormDatabase) {
	showRepo := show.NewRepository(db.DB())
	showService := show.NewService(showRepo)
	movieRepo := movie.NewRepository(db.DB())
	movieService := movie.NewService(movieRepo)
	hallRepo := hall.NewRepository(db.DB())
	hallService := hall.NewService(hallRepo)
	seatRepo := seat.NewRepository(db.DB())
	seatService := seat.NewService(seatRepo)

	showController := show.NewController(showService, movieService, hallService)
	seatController := seat.NewController(seatService, showService)
	showRoutes := r.Group("/shows/")
	{
		showRoutes.GET("", showController.GetAllShows)
		showRoutes.GET(":id", showController.GetShow)
		showRoutes.POST("", showController.CreateShow)
		showRoutes.DELETE(":id", showController.DeleteShow)
		showRoutes.PUT(":id", showController.UpdateShow)
		showRoutes.PATCH(":id", showController.UpdateShow)
		showRoutes.GET(":id/get_seats", seatController.GetSeatsForShow)
		showRoutes.GET(":id/create_seats", seatController.CreateSeatsForShow)
	}
}

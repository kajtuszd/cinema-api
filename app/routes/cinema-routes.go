package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/middleware"
	"github.com/kajtuszd/cinema-api/app/models/hall"
	"github.com/kajtuszd/cinema-api/app/models/movie"
	"github.com/kajtuszd/cinema-api/app/models/seat"
	"github.com/kajtuszd/cinema-api/app/models/show"
	"github.com/kajtuszd/cinema-api/app/models/ticket"
	"github.com/kajtuszd/cinema-api/app/models/user"
)

func InitializeCinemaRoutes(r *gin.Engine, db *database.GormDatabase) {
	showService := show.NewService(show.NewRepository(db.DB()))
	movieService := movie.NewService(movie.NewRepository(db.DB()))
	hallService := hall.NewService(hall.NewRepository(db.DB()))
	seatService := seat.NewService(seat.NewRepository(db.DB()))
	userService := user.NewService(user.NewRepository(db.DB()))
	ticketService := ticket.NewService(ticket.NewRepository(db.DB()))

	hallController := hall.NewController(hallService)
	movieController := movie.NewController(movieService)
	ticketController := ticket.NewController(ticketService, userService, seatService)
	showController := show.NewController(showService, movieService, hallService)
	seatController := seat.NewController(seatService, showService)

	hallRoutes := r.Group("/halls/")
	{
		hallRoutes.GET("", hallController.GetAllHalls)
		hallRoutes.GET(":id", hallController.GetHall)
		hallRoutes.POST("", hallController.CreateHall)
		hallRoutes.DELETE(":id", hallController.DeleteHall)
		hallRoutes.PUT(":id", hallController.UpdateHall)
	}
	movieRoutes := r.Group("/movies/")
	{
		movieRoutes.GET("", movieController.GetAllMovies)
		movieRoutes.GET(":id", movieController.GetMovie)
		movieRoutes.POST("", movieController.CreateMovie)
		movieRoutes.DELETE(":id", movieController.DeleteMovie)
		movieRoutes.PUT(":id", movieController.UpdateMovie)
	}
	showRoutes := r.Group("/shows/")
	{
		showRoutes.GET("", showController.GetAllShows)
		showRoutes.GET(":id", showController.GetShow)
		showRoutes.POST("", showController.CreateShow)
		showRoutes.DELETE(":id", showController.DeleteShow)
		showRoutes.PUT(":id", showController.UpdateShow)
		showRoutes.GET(":id/get_seats", seatController.GetSeatsForShow)
		showRoutes.GET(":id/create_seats", seatController.CreateSeatsForShow)
		showRoutes.GET(":id/delete_seats", seatController.DeleteSeatsForShow)
	}
	ticketRoutes := r.Group("/tickets/")
	{
		ticketRoutes.GET(":id", middleware.JWTAuth(db), ticketController.GetTicket)
		ticketRoutes.GET("", middleware.JWTAuth(db), ticketController.GetTickets)
		ticketRoutes.POST("", middleware.JWTAuth(db), ticketController.CreateTicket)
		ticketRoutes.PUT(":id", middleware.JWTAuth(db), ticketController.UpdateTicket)
		ticketRoutes.DELETE(":id", middleware.JWTAuth(db), ticketController.DeleteTicket)
	}
}

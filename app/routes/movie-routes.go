package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/models/movie"
)

func InitializeMovieRoutes(r *gin.Engine, db *database.GormDatabase) {
	movieRepo := movie.NewRepository(db.DB())
	movieService := movie.NewService(movieRepo)
	movieController := movie.NewController(movieService, validator.New())
	movieRoutes := r.Group("/movies/")
	{
		movieRoutes.GET("", movieController.GetAllMovies)
		movieRoutes.GET(":id", movieController.GetMovie)
		movieRoutes.POST("", movieController.CreateMovie)
		movieRoutes.DELETE(":id", movieController.DeleteMovie)
		movieRoutes.PUT(":id", movieController.UpdateMovie)
		movieRoutes.PATCH(":id", movieController.UpdateMovie)
	}
}

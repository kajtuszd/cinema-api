package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/controllers"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/middleware"
	"github.com/kajtuszd/cinema-api/app/repositories"
	"github.com/kajtuszd/cinema-api/app/services"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())
	db := &database.GormDatabase{}
	err := db.Connect()
	if err != nil {
		panic("Can not connect to database")
	}
	defer db.Close()
	if err != nil {
		panic(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRepo := repositories.New(db.DB())
	userService := services.New(userRepo)
	userController := controllers.New(userService)

	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:username", userController.GetUser)
	r.POST("/users/", userController.CreateUser)
	r.DELETE("/users/:username", userController.DeleteUser)
	r.PUT("/users/:username", userController.UpdateUser)
	r.PATCH("/users/:username", userController.UpdateUser)

	err = db.Migrate()
	if err != nil {
		panic(err)
	}
	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/database"
	"github.com/kajtuszd/cinema-api/app/middleware"
	"github.com/kajtuszd/cinema-api/app/routes"
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

	routes.InitializeRoutes(r, db)
	err = db.Migrate()
	if err != nil {
		panic(err)
	}
	r.Run()
}

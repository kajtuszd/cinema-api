package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // Import the sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                      // Import the theme

	"github.com/gin-gonic/gin"
	"github.com/kajtuszd/cinema-api/app/initializers"
	"github.com/kajtuszd/cinema-api/app/middleware"
	"github.com/kajtuszd/cinema-api/app/routes"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())
	db := &initializers.GormDatabase{}
	err := db.Connect()
	if err != nil {
		panic("Can not connect to database")
	}
	defer db.Close()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}
	initializers.ConfigureAdmin(r)
	routes.InitializeRoutes(r, db)

	r.Run()
}

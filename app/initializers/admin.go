package initializers

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"
	"os"
)

func ConfigureAdmin(router *gin.Engine) {
	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	eng.AddPlugins(adminPlugin)
	cfg := loadConfig()
	_ = eng.AddConfig(&cfg).Use(router)
}

func loadConfig() config.Config {
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:         os.Getenv("POSTGRES_HOST"),
				Port:         os.Getenv("POSTGRES_PORT"),
				User:         os.Getenv("POSTGRES_USER"),
				Pwd:          os.Getenv("POSTGRES_PASSWORD"),
				Name:         os.Getenv("POSTGRES_DB"),
				MaxIdleConns: 50,
				MaxOpenConns: 150,
				Driver:       os.Getenv("DB_DRIVER"),
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}
	return cfg
}

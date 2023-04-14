package tables

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/joho/godotenv"
	"os"
)

func loadEnv() (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
		return "", err
	}
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbSSLMode := os.Getenv("SSLMODE")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)
	return dsn, nil
}

func GetUserTable(ctx *context.Context) table.Table {
	dsn, err := loadEnv()
	if err != nil {
		return nil
	}
	userTable := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", dsn))
	formList := userTable.GetForm()
	formList.AddField("ID", "ID", db.Integer, form.Default)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("First Name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last Name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Phone Number", "phone", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Is Moderator", "is_moderator", db.Boolean, form.Switch)
	return userTable
}

func GetPostTable(ctx *context.Context) table.Table {
	//dsn, err := loadEnv()
	//if err != nil {
	//	return nil
	//}
	//postTable := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("postgresql", dsn))
	postTable := table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})
	info := postTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)

	formList := postTable.GetForm()

	formList.SetTable("posts")
	formList.AddField("ID", "id", db.Integer, form.Default)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")
	return postTable
}

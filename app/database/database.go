package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kajtuszd/cinema-api/app/models/hall"
	"github.com/kajtuszd/cinema-api/app/models/movie"
	"github.com/kajtuszd/cinema-api/app/models/show"
	"github.com/kajtuszd/cinema-api/app/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database interface {
	Connect() error
	Close() error
	Migrate() error
}

type GormDatabase struct {
	database *gorm.DB
}

func (db *GormDatabase) DB() *gorm.DB {
	return db.database
}

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

func (db *GormDatabase) Connect() error {
	dsn, err := loadEnv()
	if err != nil {
		return err
	}
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.database = conn
	return nil
}

func (db *GormDatabase) Close() error {
	conn, err := db.database.DB()
	if err != nil {
		return err
	}
	err = conn.Close()
	if err != nil {
		return err
	}
	fmt.Println("Database connection closed")
	return nil
}

func (db *GormDatabase) Migrate() error {
	err := db.database.AutoMigrate(&user.User{}, &movie.Movie{}, &hall.Hall{}, &show.Show{})
	if err != nil {
		panic(err)
		return err
	}
	fmt.Println("Migrations done")
	return nil
}

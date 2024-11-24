package databse

import (
	"fmt"
	"log"
	"os"

	"github.com/danieldenton/meditation-app-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var DB DBinstance

func ConnectDB() {
	dsn := fmt.Sprint(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Los_Angeles",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating the schema")
	db.AutoMigrate(&models.User{})

	DB = DBinstance{
		DB: db,
	}
}

package databse

import (
	"fmt"
	"os"

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
	gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

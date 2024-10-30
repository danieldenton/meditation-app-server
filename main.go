package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	postgres_secret := os.Getenv("POSTGRES_SECRET")
	connStr := fmt.Sprintf("postgres://postgres:%s@localhost:5432/meditaiondb?sslmode=disable", postgres_secret)
}

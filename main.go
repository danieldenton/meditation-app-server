package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	Uid       string
	FirstName string
	LastName  string
	Email     string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	postgres_secret := os.Getenv("POSTGRES_PASSWORD")

	connStr := fmt.Sprintf("postgres://postgres:%s@localhost:5432/meditationdb?sslmode=disable", postgres_secret)

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductionTable(db)

	user := User{"HFbwu4jiaLMxEjC8nMarOjJ55Ou2", "Daniel", "Denton", "danieldentondev@gmail.com"}
	pk := insertUser(db, user)

	fmt.Println("User created with id:", pk)
}

func createProductionTable(db *sql.DB) {
	//ID
	//first name
	//last name
	//email
	//created_at
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		uid VARCHAR(100) NOT NULL,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
		)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func insertUser(db *sql.DB, user User) int {
	query := `INSERT INTO users (uid, first_name, last_name, email) VALUES ($1, $2, $3, $4) RETURNING id`

	var pk int
	err := db.QueryRow(query, user.Uid, user.FirstName, user.LastName, user.Email).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

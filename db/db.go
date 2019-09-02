package db

import (
	"database/sql"
	"fmt"
	"os"

	// Auto-load env vars from .env file
	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB

// ConnectToDB attempts to connect to the Postgres database running in a container
// Spin up with: docker-compose up -d when in root dir
//
// https://godoc.org/github.com/lib/pq
func ConnectToDB() {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
	)

	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		// TODO: real error handling
		panic(err)
	}

	// Check that our connection to the db seems alright
	err = connection.Ping()
	if err != nil {
		// TODO: real error handling
		panic(err)
	}

	// Set global db var equal to the connection pointer
	db = connection

	fmt.Println("Successfully connected to Postgres")
}

// GetDB returns a pointer to the database that ConnectToDB() connects to
func GetDB() *sql.DB {
	return db
}

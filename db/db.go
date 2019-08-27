package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// ConnectToDB attempts to connect to the Postgres database running in a container
// Spin up with: docker-compose up -d when in root dir
//
// https://godoc.org/github.com/lib/pq
func ConnectToDB() {
	connection, err := sql.Open("postgres", "user=postgres password=pw dbname=postgres sslmode=disable")
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

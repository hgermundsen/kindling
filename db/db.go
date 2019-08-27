package db

import (
	"database/sql"
	"fmt"
)

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

	err = connection.Ping()
	if err != nil {
		// TODO: real error handling
		panic(err)
	}

	fmt.Println("Successfully connected to Postgres")
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/nchaloult/kindling/api/message"
)

func main() {
	router := httprouter.New()

	router.GET("/api/message", message.GetAllMessages)

	connectToDB()

	fmt.Println("Listening @ localhost:8080....")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Attempt to connect to the Postgres database running in a container
// Spin up with: docker-compose up -d when in root dir
//
// https://godoc.org/github.com/lib/pq
func connectToDB() {
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

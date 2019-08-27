package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/nchaloult/kindling/api/message"
	"github.com/nchaloult/kindling/db"
)

func main() {
	router := httprouter.New()

	router.GET("/api/message", message.GetAllMessages)

	db.ConnectToDB()

	fmt.Println("Listening @ localhost:8080....")
	log.Fatal(http.ListenAndServe(":8080", router))
}

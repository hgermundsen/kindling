package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"github.com/nchaloult/kindling/api/message"
)

func main() {
	router := httprouter.New()

	router.GET("/api/message", message.GetAllMessages)

	fmt.Println("Listening @ localhost:8080....")
	log.Fatal(http.ListenAndServe(":8080", router))
}

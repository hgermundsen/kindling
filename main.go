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

	// Dependency injection for message entity
	messageRepo := message.NewRepo(
		message.FetchAllMessages,
		message.FetchMessageByID,
		message.InsertMessage,
	)
	messageController := message.NewController(messageRepo)

	// Defining message routes
	router.GET("/api/message", messageController.GetAllMessages)
	router.GET("/api/message/:id", messageController.GetMessageByID)
	router.POST("/api/message", messageController.CreateMessage)

	db.ConnectToDB()

	fmt.Println("Listening @ localhost:8080....")
	log.Fatal(http.ListenAndServe(":8080", router))
}

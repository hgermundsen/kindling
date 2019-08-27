package message

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/nchaloult/kindling/common"
)

// GetAllMessages responds with all messages in the db
//
// GET /api/message
func GetAllMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	messages, err := FetchAllMessages()
	common.ConstructResponse(w, messages, err)
}

// GetMessageByID responds with the message that has the provided id
//
// GET /api/message/:id
func GetMessageByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	message, err := FetchMessageByID(ps.ByName("id"))
	common.ConstructResponse(w, message, err)
}

// CreateMessage inserts a new message into the database from the provided request body
//
// POST /api/message
func CreateMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// May decide to move lots of the logic in here to other files, like common

	// Create struct that request body will be parsed "into"
	newMessage := Message{}

	// Parse the request body & try to "fit it" to newMessage
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		//TODO: real error handling
		common.ConstructResponse(w, newMessage, err)
	}

	// For now, just print the struct that we "fit into" to see if that went alright
	fmt.Println(newMessage)
	w.WriteHeader(http.StatusNoContent)

	// Insert newMessage into the DB
}

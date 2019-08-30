package message

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/nchaloult/kindling/common"
)

func dummyFetchAll() ([]Message, error) {
	dummyMessage := Message{
		ID:        1,
		Title:     "dummy message",
		Content:   "dummy content. i sure hope this works",
		Upvotes:   200,
		Downvotes: 0,
		Flags:     0,
	}
	output := make([]Message, 1)
	output = append(output, dummyMessage)
	return output, nil
}

var repo *Repo = NewRepo(dummyFetchAll, FetchMessageByID, InsertMessage)

// GetAllMessages responds with all messages in the db
//
// GET /api/message
func GetAllMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	messages, err := repo.fetchAllMessages()
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

	// Insert newMessage into the DB
	err = InsertMessage(newMessage)
	if err != nil {
		//TODO: real error handling
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Everything seemed to go alright. Return 204
	w.WriteHeader(http.StatusNoContent)
}

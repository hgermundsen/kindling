package message

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/nchaloult/kindling/common"
)

// Controller is a struct that exposes access to the functions defined in this
// file.
//
// Mainly exists for dependency injection, which makes testing very simple.
type Controller struct {
	repo *Repo
}

// NewController is the default constructor for the Controller struct
func NewController(repo *Repo) *Controller {
	return &Controller{
		repo: repo,
	}
}

// GetAllMessages responds with all messages in the db
//
// GET /api/message
func (c *Controller) GetAllMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	messages, err := c.repo.fetchAllMessages()
	common.ConstructResponse(w, messages, err)
}

// GetMessageByID responds with the message that has the provided id
//
// GET /api/message/:id
func (c *Controller) GetMessageByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	message, err := c.repo.fetchMessageByID(ps.ByName("id"))
	common.ConstructResponse(w, message, err)
}

// CreateMessage inserts a new message into the database from the provided request body
//
// POST /api/message
func (c *Controller) CreateMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// May decide to move lots of the logic in here to other files, like common

	// Create struct that request body will be parsed "into"
	newMessage := Message{}

	// Parse the request body & try to "fit it" to newMessage
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		//TODO: real error handling
		common.ConstructResponse(w, newMessage, err)
	}

	// If request body didn't contain at least a title and content, then return
	// status code 400
	if newMessage.Title == "" || newMessage.Content == "" {
		// TODO: real error handling
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	// Insert newMessage into the DB
	err = c.repo.insertMessage(newMessage)
	if err != nil {
		//TODO: real error handling
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Everything seemed to go alright. Return 204
	w.WriteHeader(http.StatusNoContent)
}

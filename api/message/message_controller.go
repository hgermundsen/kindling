package message

import (
	"encoding/json"
	"errors"
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
		common.ConstructResponse(w, nil, errors.New(common.EInvalidJSON))
		return
	}

	// If request body didn't contain at least a title and content, then return
	// status code 400
	if newMessage.Title == "" || newMessage.Content == "" {
		common.ConstructResponse(w, nil, errors.New(common.EMessageMissingRequiredFields))
		return
	}

	// Insert newMessage into the DB
	err = c.repo.insertMessage(newMessage)
	if err != nil {
		common.ConstructResponse(w, nil, err)
		return
	}

	// Everything seemed to go alright. Return 204
	w.WriteHeader(http.StatusNoContent)
}

// DeleteMessageByID deletes the message that has the provided id from the db
//
// DELETE /api/message/:id
func (c *Controller) DeleteMessageByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Attempt to delete the appropriate message
	err := c.repo.deleteMessageByID(ps.ByName("id"))
	if err != nil {
		common.ConstructResponse(w, nil, err)
		return
	}

	// Everything seemed to go alright. Return 204
	w.WriteHeader(http.StatusNoContent)
}

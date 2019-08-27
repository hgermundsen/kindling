package message

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// GetAllMessages responds with all messages in the db
//
// GET /api/message
func GetAllMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set important header info (tidying up the response sent)
	w.Header().Set("Content-Type", "application/json")

	// Get all messages fron the repository
	messages, err := FetchAllMessages()
	if err != nil {
		// TODO: real error handling
		// Return 500 if something went wrong
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Turn our dummyMessages into json & write that json to the ResponseWriter
	json.NewEncoder(w).Encode(messages)
}

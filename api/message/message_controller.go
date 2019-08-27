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
	// Get all messages fron the repository
	dummyMessages := FetchAllMessages()

	// Turn our dummyMessages into json & write that json to the ResponseWriter
	json.NewEncoder(w).Encode(dummyMessages)
}

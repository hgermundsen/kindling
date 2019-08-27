package message

import (
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

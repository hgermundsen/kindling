package message

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// GetAllMessages responds with all messages in the db
//
// GET /api/message
func GetAllMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "You've hit the /api/message endpoint")
}

package common

import (
	"fmt"
	"net/http"
)

// Common error messages that are handled by common.ConstructResponse
var (
	// 400s
	EInvalidJSON                  = fmt.Sprintf("%s: Invalid JSON", http.StatusText(http.StatusBadRequest))
	EMessageMissingRequiredFields = fmt.Sprintf("%s: Message must contain a title and content", http.StatusText(http.StatusBadRequest))
	ENotFound                     = http.StatusText(http.StatusNotFound)
	// 500s
	EDBInsert = fmt.Sprintf("%s: Failed to insert new record into database", http.StatusText(http.StatusInternalServerError))
	EDBDelete = fmt.Sprintf("%s: Failed to delete a record from the database", http.StatusText(http.StatusInternalServerError))
	ESQL      = fmt.Sprintf("%s: Something went wrong executing an SQL statement", http.StatusText(http.StatusInternalServerError))
)

// ECodes maps custom error messages with the status codes that go along
// with them
var ECodes = map[string]int{
	// 400s
	EInvalidJSON:                  http.StatusBadRequest,
	EMessageMissingRequiredFields: http.StatusBadRequest,
	ENotFound:                     http.StatusNotFound,
	// 500s
	EDBInsert: http.StatusInternalServerError,
	EDBDelete: http.StatusInternalServerError,
	ESQL:      http.StatusInternalServerError,
}

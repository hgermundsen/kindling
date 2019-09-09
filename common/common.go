package common

import (
	"encoding/json"
	"net/http"
	"reflect"
)

// ConstructResponse is a helpful utility func that gets called in all
// *_controllers. It adds important headers to a response, and marshals the
// body of the response into json. ConstructResponse also handles taking care
// of responding with the right status code, depending on the provided error.
func ConstructResponse(w http.ResponseWriter, body interface{}, err error) {
	// Set important header info (tidying up the response sent)
	w.Header().Set("Content-Type", "application/json")

	// If something went wrong, handle based on the custom error message
	if err != nil {
		// switch err.Error() {
		// // 400s
		// case EInvalidJSON:
		// 	http.Error(w, EInvalidJSON, ECodes[EInvalidJSON])
		// case EMessageMissingRequiredFields:
		// 	http.Error(w, EMessageMissingRequiredFields, ECodes[EMessageMissingRequiredFields])
		// // 500s
		// case EDBInsert:
		// 	http.Error(w, EDBInsert, ECodes[EDBInsert])
		// }

		http.Error(w, err.Error(), ECodes[err.Error()])
		return
	}

	// If nothing went wrong, but there's no content to respond with, assume
	// (for now) that we should respond with a 404
	if body == nil || isZeroOrEmpty(body) {
		http.Error(w, ENotFound, ECodes[ENotFound])
		return
	}

	// Respond with a 200
	// Default status code is 200, so we don't have to specify that manually
	json.NewEncoder(w).Encode(body)
}

// Returns whether the provided input contains the "zero value" of its type.
//
// Ex: isZeroOrEmpty(0) -> true
// Ex: isZeroOrEmpty(Employee{}) -> true
// Ex: isZeroOrEmpty("foo") -> false
//
// https://stackoverflow.com/questions/33115946/how-to-know-if-a-variable-of-arbitrary-type-is-zero-in-golang
func isZeroOrEmpty(input interface{}) bool {
	return reflect.DeepEqual(input, reflect.Zero(reflect.TypeOf(input)).Interface())
}

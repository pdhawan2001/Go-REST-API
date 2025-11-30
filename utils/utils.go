package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// the error returned in these functions will be used or returned, wherever they are called

var Validate = validator.New()

// this is taking the request as we are converting the data we are passing into the api to the payload
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}

	// the body is being decoded into the payload in the ParseJSON function
	return json.NewDecoder(r.Body).Decode(payload)
}

// this is taking the response, as we will look if there is any error in that response, and then we will handle it
// so, we take 1) response, 2) status that was returned, 3) anything which is a valid output
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// this will also write a json response, but only if there is an error
func WriteError(w http.ResponseWriter, status int, err error) {
	// map below is a key-value store, which will put the key and values in the error
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

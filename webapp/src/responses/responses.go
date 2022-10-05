package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error is a struct to map the error coming in the request response to JSON
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// Defines in the header that the response type is JSON
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	// If have data
	if data != nil {
		// Convert data to JSON
		if error := json.NewEncoder(w).Encode(data); error != nil {
			log.Fatal(error)
		}
	}

}

// StatusCodeError handle requests with StatusCode 400 or higher
func StatusCodeError(w http.ResponseWriter, r *http.Response) {
	var error ErrorAPI
	// Send the response body to parameter
	// Convert response body from JSON to struct
	json.NewDecoder(r.Body).Decode(&error)

	JSON(w, r.StatusCode, error)
}

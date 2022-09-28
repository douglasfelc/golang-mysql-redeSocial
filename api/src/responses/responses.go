package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

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

// Error returns an error in JSON format
func Error(w http.ResponseWriter, statusCode int, error error) {
	// Calls JSON(), passing the error response structure
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: error.Error(),
	})
}

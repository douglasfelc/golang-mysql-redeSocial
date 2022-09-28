package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	// Convert data to JSON
	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

// Error returns an error in JSON format
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	// Calls JSON(), passing the error response structure
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}

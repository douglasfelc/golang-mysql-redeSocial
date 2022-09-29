// It will be a layer that will be between the request and the response
// It is used when there is a function that must be applied to all routes
// Instead of entering route by route and placing a function X, you create a middleware that makes the application of this function

package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		// Execute what is coming by parameter
		next(w, r)
	}
}

// Authenticate checks if the user making the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Validates the token passed in the request
		if error := authentication.ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}

		// Execute what is coming by parameter
		next(w, r)
	}
}

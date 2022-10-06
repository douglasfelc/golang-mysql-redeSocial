package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger write request information to terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		// Execute what is coming by parameter
		next(w, r)
	}
}

// Authenticate checks if the authentication cookie exists
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// If CANNOT read the cookie
		if _, error := cookies.Read(r); error != nil {
			http.Redirect(w, r, "/signin", 302)
			return
		}

		// Execute what is coming by parameter
		next(w, r)
	}
}

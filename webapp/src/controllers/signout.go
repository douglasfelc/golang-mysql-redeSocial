package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// SignOut clears the authentication data saved in the user's browser
func SignOut(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/signin", 302)
}

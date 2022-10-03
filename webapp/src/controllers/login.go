package controllers

import (
	"net/http"
)

// LoginScreen renders the login screen
func LoginScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Screen"))
}

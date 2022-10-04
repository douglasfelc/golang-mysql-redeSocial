// All functions that render pages

package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoginScreen renders the login screen
func LoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// SignUpScreen renders the user registration page
func SignUpScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

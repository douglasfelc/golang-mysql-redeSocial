package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoginScreen renders the login screen
func LoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

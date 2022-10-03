package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login is responsible for authenticating a user to the API
func Login(w http.ResponseWriter, r *http.Request) {
	// Reads the Request.Body
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	// Convert JSON (requestBody) to struct (user)
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// Connect to the database
	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	// Create the repository, passing the database as a parameter
	repository := repositories.NewUsersRepository(db)

	// Get the repository for the user with this email and return ID and password with hash
	databaseUser, error := repository.GetByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// Checks if the password passed is the same as the one in the database for this user
	if error = security.CheckPassword(databaseUser.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Create a token signed with the user's permissions
	token, error := authentication.CreateToken(databaseUser.ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	w.Write([]byte(token))
}

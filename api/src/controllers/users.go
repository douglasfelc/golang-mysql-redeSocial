package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser insert a user...
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Reads the Request.Body
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	// Convert JSON (requestBody) to struct (user)
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// User validation
	if erro = user.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Connect to the database
	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Create the repository, passing the database as a parameter
	repository := repositories.NewUsersRepository(db)

	// Create the user in the database, using the Create method
	// The Create method returns the user id
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// If successful, reply with StatusCreated, and user
	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários!"))
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}

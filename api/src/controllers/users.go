package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser insert a user...
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Reads the Request.Body
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	// Convert JSON (requestBody) to struct (user)
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	// Connect to the database
	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	// Create the repository, passing the database as a parameter
	repository := repositories.NewUsersRepository(db)

	// Create the user in the database, using the Create method
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))
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

package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser insert a user...
func CreateUser(w http.ResponseWriter, r *http.Request) {

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

	// User validation
	if error = user.Prepare("newRegister"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// Connect to the database
	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// Create the repository, passing the database as a parameter
	repository := repositories.NewUsersRepository(db)

	// Create the user in the database, using the Create method
	// The Create method returns the user id
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusCreated, and user
	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers receives the name or nick in the get user field, and responds with a JSON with the filtered users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// ToLower: Convert request to lowercase
	// Get: Get the value passed with this name
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	// Connect to the database
	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	// Create the repository, passing the database as a parameter
	repository := repositories.NewUsersRepository(db)

	// Requests the repository to search for the name or nick passed with user in get
	users, error := repository.Search(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	// If successful, reply with StatusOK and filtered users
	responses.JSON(w, http.StatusOK, users)
}

// GetUser request the repository for the user by ID and returns the data in JSON
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
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

	// Requests that the repository search and return the user with id sent by parameter
	user, error := repository.SearchByID(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	// If successful, reply with StatusOK and filtered user
	responses.JSON(w, http.StatusOK, user)
}

// Update sends received user data to repository change in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

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

	// User validation
	if error = user.Prepare("update"); error != nil {
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

	// Create the user in the database, using the Create method
	// The Create method returns the user id
	if error = repository.Update(userID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser sends a request to the repository to delete a user by the id passed
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
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
	if error = repository.Delete(userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

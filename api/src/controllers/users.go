package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
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
	users, error := repository.Get(nameOrNick)
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
	user, error := repository.GetByID(userID)
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

	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Check permission
	// Checks if the user being updated is different from the user requesting the change
	if userID != userIDinToken {
		responses.Error(w, http.StatusForbidden,
			errors.New("You are not allowed to change this user"),
		)
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

	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Check permission
	// Checks if the user being deleted is different from the user requesting the change
	if userID != userIDinToken {
		responses.Error(w, http.StatusForbidden,
			errors.New("You do not have permission to delete this user"),
		)
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

// FollowUser allows one user to follow another
func FollowUser(w http.ResponseWriter, r *http.Request) {

	// Extract the userId from the token, to check its permissions
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// If the follower ID is the same as the user ID to be followed
	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Can't follow yourself"))
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
	if error = repository.Follow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

// UnFollowUser allows one user to unfollow another
func UnFollowUser(w http.ResponseWriter, r *http.Request) {

	// Extract the userId from the token, to check its permissions
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// If the follower ID is the same as the user ID to be unfollowed
	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Can't unfollow yourself"))
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
	if error = repository.UnFollow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

// GetFollowers get all followers of a user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
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
	followers, error := repository.GetFollowers(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusOK and followers in JSON
	responses.JSON(w, http.StatusOK, followers)
}

// GetFollowing get all users that a user is following
func GetFollowing(w http.ResponseWriter, r *http.Request) {
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
	users, error := repository.GetFollowing(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusOK and followers in JSON
	responses.JSON(w, http.StatusOK, users)
}

// UpdatePassword update a user's password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// Check permission
	// Checks if the user being updated is different from the user requesting the change
	if userIDinToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You are not allowed to change this user's password"))
		return
	}

	// Reads the Request.Body
	requestBody, error := ioutil.ReadAll(r.Body)

	var password models.Password
	if error = json.Unmarshal(requestBody, &password); error != nil {
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

	// Get password from user database
	PasswordInDataBase, error := repository.GetPassword(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// Checks if the current password sent is the same as the one in the database
	if error = security.CheckPassword(PasswordInDataBase, password.Current); error != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Current password is incorrect"))
		return
	}

	// Hash the password
	passwordWithHash, error := security.Hash(password.New)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// Update password in database
	if error = repository.UpdatePassword(userID, string(passwordWithHash)); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreateUser calls the API to register a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// To access form fields with FormValue
	r.ParseForm()

	// Convert the data submitted in the form to JSON
	user, error := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nick":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:3000/users
	url := fmt.Sprintf("%s/users", config.APIURL)
	// Send the request to the API with the data
	responseHttp, error := http.Post(url, "application/json", bytes.NewBuffer(user))
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer responseHttp.Body.Close()

	// If in range of the Error StatusCode
	if responseHttp.StatusCode >= 400 {
		responses.StatusCodeError(w, responseHttp)
		return
	}

	responses.JSON(w, responseHttp.StatusCode, nil)
}

// UnFollowUser calls the API to stop following a user
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:3000/users/{userId}/unfollow
	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)

	// Send the request with authentication to the API with the data
	response, error := requests.RequestWithAuthentication(r, http.MethodPost, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	// If in range of the Error StatusCode
	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// FollowUser calls the API to follow a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:3000/users/{userId}/follow
	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)

	// Send the request with authentication to the API with the data
	response, error := requests.RequestWithAuthentication(r, http.MethodPost, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	// If in range of the Error StatusCode
	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdateProfile calls the API to edit a user
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, error := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	// Convert the id in the cookie to uint64
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// Mount the url, eg http://localhost:3000/users/{userId}
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	// Send the request with authentication to the API with the data
	response, error := requests.RequestWithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(user))
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	// If in range of the Error StatusCode
	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePassword calls the API to update the password of the logged in user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	senhas, error := json.Marshal(map[string]string{
		"current": r.FormValue("current"),
		"new":     r.FormValue("new"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	// Convert the id in the cookie to uint64
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// Mount the url, eg http://localhost:3000/users/{userId}/update-password
	url := fmt.Sprintf("%s/users/%d/update-password", config.APIURL, userID)

	// Send the request with authentication to the API with the data
	response, error := requests.RequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	// If in range of the Error StatusCode
	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// DeleteUser calls API to remove user account
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	// Convert the id in the cookie to uint64
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// Mount the url, eg http://localhost:3000/users/{userId}
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	// Send the request with authentication to the API with the data
	response, error := requests.RequestWithAuthentication(r, http.MethodDelete, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	// If in range of the Error StatusCode
	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

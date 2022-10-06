package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// SignIn is responsible for authenticating a user
func SignIn(w http.ResponseWriter, r *http.Request) {
	// To access form fields with FormValue
	r.ParseForm()

	// Convert the data submitted in the form to JSON
	user, error := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:3000/signin
	url := fmt.Sprintf("%s/signin", config.APIURL)
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

	var returnAuthentication models.Authentication
	// Convert response body from JSON to struct
	if error = json.NewDecoder(responseHttp.Body).Decode(&returnAuthentication); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Save authentication return data in cookies
	if error = cookies.Save(w, returnAuthentication.ID, returnAuthentication.Token); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

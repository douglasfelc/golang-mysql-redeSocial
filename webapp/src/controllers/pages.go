// All functions that render pages

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

// SignInScreen renders the signin screen
func SignInScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signin.html", nil)
}

// SignUpScreen renders the user registration page
func SignUpScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

// FeedScreen renders the feed screen
func FeedScreen(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, error := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	var posts []models.Post
	// Convert response body from JSON to struct
	if error = json.NewDecoder(response.Body).Decode(&posts); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	fmt.Println(posts)

	utils.ExecuteTemplate(w, "feed.html", posts)
}

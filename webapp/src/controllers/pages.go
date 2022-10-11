// All functions that render pages

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
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

	// Read the Cookie, ignoring the error as the middleware has already verified this
	cookie, _ := cookies.Read(r)

	// Convert the id in the cookie to uint64
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// Send request posts and cookie userID to the template
	utils.ExecuteTemplate(w, "feed.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

// UpdatePostScreen renders the post edit page
func UpdatePostScreen(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
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

	var post models.Post
	// Convert response body from JSON to struct
	if error = json.NewDecoder(response.Body).Decode(&post); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Send request posts to the template
	utils.ExecuteTemplate(w, "post-update.html", post)
}

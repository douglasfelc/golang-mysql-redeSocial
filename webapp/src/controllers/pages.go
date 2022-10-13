// All functions that render pages

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	cookie, _ := cookies.Read(r)

	// Check if token exists
	// In addition, below it is necessary to check if the token is still valid
	if cookie["token"] != "" {

		// Convert the id in the cookie to uint64
		userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

		url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
		response, _ := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
		defer response.Body.Close()

		// Checks if the token is valid
		if response.StatusCode <= 201 {
			// Send to feed
			http.Redirect(w, r, "/feed", 302)
			return
		}
	}

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

	// If the StatusCode is an error
	if response.StatusCode >= 400 {
		// Send to signin
		http.Redirect(w, r, "/signin", 302)
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

// UsersScreen renders the page with the users filtered out
func UsersScreen(w http.ResponseWriter, r *http.Request) {

	// Get the "search" user coming by get
	nameOrNick := strings.ToLower(r.URL.Query().Get("search"))

	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)
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

	var users []models.User
	// Convert response body from JSON to struct
	if error = json.NewDecoder(response.Body).Decode(&users); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Send request users to the template
	utils.ExecuteTemplate(w, "users.html", users)
}

// UserScreen renders the user profile page
func UserScreen(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{userId}
	params := mux.Vars(r)

	// Convert ID to uint64
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	// Convert the id in the cookie to uint64
	LoggedInUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// Get the full user (with following, followers and posts)
	user, error := models.GetFullUser(userID, r)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User           models.User
		LoggedInUserID uint64
	}{
		User:           user,
		LoggedInUserID: LoggedInUserID,
	})
}

// Profile redirects to the logged in user page
func Profile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	LoggedInUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	url := fmt.Sprintf("/users/%d", LoggedInUserID)
	http.Redirect(w, r, url, 302)
}

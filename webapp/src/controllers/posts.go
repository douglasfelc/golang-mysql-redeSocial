package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreatePost calls the API to create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// To access form fields with FormValue
	r.ParseForm()

	// Convert the data submitted in the form to JSON
	post, error := json.Marshal(map[string]string{
		"content": r.FormValue("content"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:5000/posts
	url := fmt.Sprintf("%s/posts", config.APIURL)
	// Send the request with authentication to the API with the data
	responseHttp, error := requests.RequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(post))
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

// LikePost calls the API to like a post
func LikePost(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:3000/posts/{postId}/like
	url := fmt.Sprintf("%s/posts/%d/like", config.APIURL, postID)
	// Send the request with authentication to the API with the data
	responseHttp, error := requests.RequestWithAuthentication(r, http.MethodPost, url, nil)
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

// DisLikePost calls the API to like a post
func DisLikePost(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	// Mount the url, eg http://localhost:5000/posts/{postId}/dislike
	url := fmt.Sprintf("%s/posts/%d/dislike", config.APIURL, postID)

	fmt.Println(url)

	// Send the request with authentication to the API with the data
	responseHttp, error := requests.RequestWithAuthentication(r, http.MethodPost, url, nil)
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

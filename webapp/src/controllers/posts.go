package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
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

	// Mount the url, eg http://localhost:3000/posts
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

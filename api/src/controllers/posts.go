package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Reads the Request.Body
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	// Convert JSON (requestBody) to struct (post)
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	post.AuthorID = userIDinToken

	// Post validation
	if error = post.Prepare(); error != nil {
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
	repository := repositories.NewPostsRepository(db)

	// Create the post in the database, using the Create method
	// The Create method returns the post id
	post.ID, error = repository.Create(post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusCreated, and post
	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts get the posts to display in the feed
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Connect to the database
	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	// Create the repository, passing the database as a parameter
	repository := repositories.NewPostsRepository(db)

	// Get from database posts from followers and from the user who made a request
	posts, error := repository.Get(userIDinToken)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	// If successful, reply with StatusOK and filtered posts
	responses.JSON(w, http.StatusOK, posts)
}

// GetPost get a post
func GetPost(w http.ResponseWriter, r *http.Request) {

	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
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
	repository := repositories.NewPostsRepository(db)

	// Requests that the repository search and return the post with id sent by parameter
	post, error := repository.GetByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	// If successful, reply with StatusOK and filtered post
	responses.JSON(w, http.StatusOK, post)
}

// UpdatePost change the data of a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
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
	repository := repositories.NewPostsRepository(db)

	// Get the post data from the database by ID, to know who is the author of the post, and check permissions below
	postInDataBase, error := repository.GetByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// Check permission
	// Check if the user in the requesting token is the same as the post author
	if postInDataBase.AuthorID != userIDinToken {
		responses.Error(w, http.StatusForbidden,
			errors.New("You cannot change a post that is not yours"),
		)
		return
	}

	// Reads the Request.Body
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	// Convert JSON (requestBody) to struct (post)
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// Post validation
	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	// Update the post in the database
	if error = repository.Update(postID, post); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Extract the userId from the token, to check its permissions
	userIDinToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
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
	repository := repositories.NewPostsRepository(db)

	// Get the post data from the database by ID, to know who is the author of the post, and check permissions below
	postInDataBase, error := repository.GetByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// Check permission
	// Check if the user in the requesting token is the same as the post author
	if postInDataBase.AuthorID != userIDinToken {
		responses.Error(w, http.StatusForbidden,
			errors.New("You cannot delete a post that is not yours"),
		)
		return
	}

	// Delete post in the database
	if error = repository.Delete(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

// GetPostsByUser get all posts from a user
func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
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
	repository := repositories.NewPostsRepository(db)

	// Update the post in the database
	posts, error := repository.GetByUser(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusOK and posts
	responses.JSON(w, http.StatusOK, posts)
}

// LikePost like a post
func LikePost(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
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
	repository := repositories.NewPostsRepository(db)

	// Add a like to the post in database
	if error := repository.Like(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

// DislikePost dislike a post
func DislikePost(w http.ResponseWriter, r *http.Request) {
	// Get the parameters sent in the route, ex: /{postId}
	params := mux.Vars(r)

	// Convert ID to uint64
	postID, error := strconv.ParseUint(params["postId"], 10, 64)
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
	repository := repositories.NewPostsRepository(db)

	// Decrease a like to the post in database
	if error := repository.Dislike(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	// If successful, reply with StatusNoContent
	responses.JSON(w, http.StatusNoContent, nil)
}

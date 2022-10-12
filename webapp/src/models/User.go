package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User struct
type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// GetFullUser makes 4 API requests to mount the user
func GetFullUser(userID uint64, r *http.Request) (User, error) {

	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	// Using concurrency
	go GetUser(userChannel, userID, r)
	go GetFollowers(followersChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	go GetPosts(postsChannel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	// Iterates 4x = channel numbers
	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("Error getting user")
			}

			user = userLoaded

		case followersLoaded := <-followersChannel:
			if followersLoaded == nil {
				return User{}, errors.New("Error getting followers")
			}

			followers = followersLoaded

		case followingLoaded := <-followingChannel:
			if followingLoaded == nil {
				return User{}, errors.New("Error getting who the user is following")
			}

			following = followingLoaded

		case postsLoaded := <-postsChannel:
			if postsLoaded == nil {
				return User{}, errors.New("Error getting posts")
			}

			posts = postsLoaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// GetUser calls the API to get the user's base data
func GetUser(canal chan<- User, userID uint64, r *http.Request) {
	// Mount the url, eg http://localhost:5000/users/{userId}
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	// API request with authentication
	response, erro := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	// Convert response body from JSON to struct
	if erro = json.NewDecoder(response.Body).Decode(&user); erro != nil {
		canal <- User{}
		return
	}

	canal <- user
}

// GetFollowers calls the API to get the user's followers
func GetFollowers(canal chan<- []User, userID uint64, r *http.Request) {
	// Mount the url, eg http://localhost:5000/users/{userId}/followers
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)

	// API request with authentication
	response, erro := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	// Convert response body from JSON to struct
	if erro = json.NewDecoder(response.Body).Decode(&followers); erro != nil {
		canal <- nil
		return
	}

	// If the user has no followers
	if followers == nil {
		// It sends an empty slice, because above when the information is being sent to the channel, the value nil means that it was not able to capture the information, which is different from not having the information
		canal <- make([]User, 0)
		return
	}

	canal <- followers
}

// GetFollowing calls the API to get the users followed by a user
func GetFollowing(canal chan<- []User, userID uint64, r *http.Request) {
	// Mount the url, eg http://localhost:5000/users/{userId}/following
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)

	// API request with authentication
	response, erro := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	// Convert response body from JSON to struct
	if erro = json.NewDecoder(response.Body).Decode(&following); erro != nil {
		canal <- nil
		return
	}

	// If the user has no following
	if following == nil {
		// It sends an empty slice, because above when the information is being sent to the channel, the value nil means that it was not able to capture the information, which is different from not having the information
		canal <- make([]User, 0)
		return
	}

	canal <- following
}

// GetPosts calls the API to get a user's posts
func GetPosts(canal chan<- []Post, userID uint64, r *http.Request) {
	// Mount the url, eg http://localhost:5000/users/{userId}/posts
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)

	// API request with authentication
	response, erro := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	// Convert response body from JSON to struct
	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		canal <- nil
		return
	}

	// If the user has no posts
	if posts == nil {
		// It sends an empty slice, because above when the information is being sent to the channel, the value nil means that it was not able to capture the information, which is different from not having the information
		canal <- make([]Post, 0)
		return
	}

	canal <- posts
}

package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var userRouters = []Router{
	// Route to load the user registration page
	{
		URI:                    "/signup",
		Method:                 http.MethodGet,
		Function:               controllers.SignUpScreen,
		authenticationRequired: false,
	},
	// Route to add a new user
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		authenticationRequired: false,
	},
	// Route to load filtered users page
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.UsersScreen,
		authenticationRequired: true,
	},
	// Route to load user page
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.UserScreen,
		authenticationRequired: true,
	},
	// Route to follow user by id
	// userId of the user being followed
	{
		URI:                    "/users/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		authenticationRequired: true,
	},
	// Route to unfollow user by id
	// userId of the user being unfollowed
	{
		URI:                    "/users/{userId}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnFollowUser,
		authenticationRequired: true,
	},
}

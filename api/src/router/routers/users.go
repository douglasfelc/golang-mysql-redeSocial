package routers

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Router{
	// Route to register a user
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		authenticationRequired: false,
	},
	// Route to return all users
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		authenticationRequired: true,
	},
	// Route to return a user by id
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		authenticationRequired: true,
	},
	// Route to change user by id
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		authenticationRequired: true,
	},
	// Route to delete user by id
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
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
	// Route to see a user's followers
	{
		URI:                    "/users/{userId}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowers,
		authenticationRequired: true,
	},
	// Route to see who the user is following
	{
		URI:                    "/users/{userId}/following",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowing,
		authenticationRequired: true,
	},
	// Route to see who the user can follow
	{
		URI:                    "/users/{userId}/whotofollow",
		Method:                 http.MethodGet,
		Function:               controllers.WhoToFollow,
		authenticationRequired: true,
	},
	// Route to update user password
	// Normally, POST is used to update the password.
	// Because if you make two PUT requests passing exactly the same data, the second one will have no effect.
	// But every time I make a request, a new hash will always be generated, that is, even if I pass the same data, the return is different, because the hash always changes
	{
		URI:                    "/users/{userId}/update-password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		authenticationRequired: true,
	},
}

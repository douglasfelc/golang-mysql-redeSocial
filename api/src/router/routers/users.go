package routers

import (
	"api/src/router/controllers"
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
		authenticationRequired: false,
	},
	// Route to return a user by id
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		authenticationRequired: false,
	},
	// Route to change user by id
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		authenticationRequired: false,
	},
	// Route to delete user by id
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		authenticationRequired: false,
	},
}

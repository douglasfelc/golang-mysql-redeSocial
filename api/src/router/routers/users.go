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
}

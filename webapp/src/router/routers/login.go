package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRouters = []Router{
	// Route to load login screen
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.LoginScreen,
		authenticationRequired: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.LoginScreen,
		authenticationRequired: false,
	},
}

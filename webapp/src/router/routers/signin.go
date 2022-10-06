package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var signinRouters = []Router{
	// Route to load signin screen
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.SignInScreen,
		authenticationRequired: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.SignInScreen,
		authenticationRequired: false,
	},
	{
		URI:                    "/signin",
		Method:                 http.MethodGet,
		Function:               controllers.SignInScreen,
		authenticationRequired: false,
	},
	{
		URI:                    "/signin",
		Method:                 http.MethodPost,
		Function:               controllers.SignIn,
		authenticationRequired: false,
	},
}

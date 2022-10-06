package routers

import (
	"api/src/controllers"
	"net/http"
)

// Route to register a user
var signinRouter = Router{
	URI:                    "/signin",
	Method:                 http.MethodPost,
	Function:               controllers.Signin,
	authenticationRequired: false,
}

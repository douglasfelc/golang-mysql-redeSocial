package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var signoutRouter = Router{
	URI:                    "/signout",
	Method:                 http.MethodGet,
	Function:               controllers.SignOut,
	authenticationRequired: true,
}

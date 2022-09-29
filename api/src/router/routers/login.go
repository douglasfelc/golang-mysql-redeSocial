package routers

import (
	"api/src/controllers"
	"net/http"
)

// Route to register a user
var loginRouter = Router{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	authenticationRequired: false,
}

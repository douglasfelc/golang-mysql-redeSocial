package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var feedRouter = Router{
	URI:                    "/feed",
	Method:                 http.MethodGet,
	Function:               controllers.FeedScreen,
	authenticationRequired: true,
}

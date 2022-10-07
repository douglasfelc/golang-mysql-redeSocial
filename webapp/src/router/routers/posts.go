package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var postRouters = []Router{
	// Route to add a new post
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		authenticationRequired: true,
	},
}

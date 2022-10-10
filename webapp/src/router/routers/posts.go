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
	// Route to like a post
	{
		URI:                    "/posts/{postId}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		authenticationRequired: true,
	},
	// Route to dislike a post
	{
		URI:                    "/posts/{postId}/dislike",
		Method:                 http.MethodPost,
		Function:               controllers.DisLikePost,
		authenticationRequired: true,
	},
}

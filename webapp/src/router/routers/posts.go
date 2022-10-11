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
	// Route to load a page to update a post
	{
		URI:                    "/posts/{postId}/update",
		Method:                 http.MethodGet,
		Function:               controllers.UpdatePostScreen,
		authenticationRequired: true,
	},
	// Route to update a post
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		authenticationRequired: true,
	},
	// Route to delete a post
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		authenticationRequired: true,
	},
}

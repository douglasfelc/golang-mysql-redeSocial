package routers

import (
	"api/src/controllers"
	"net/http"
)

var postRouters = []Router{
	// Route to create a post
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		authenticationRequired: true,
	},
	// Route to get posts from users that a user follows
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPosts,
		authenticationRequired: true,
	},
	// Route to get a post
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPost,
		authenticationRequired: true,
	},
	// Route to update a post
	{
		URI:                    "/posts",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		authenticationRequired: true,
	},
	// Route to delete a post
	{
		URI:                    "/posts",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		authenticationRequired: true,
	},
}

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
	// Route of a user's posts
	{
		URI:                    "/users/{userId}/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPostsByUser,
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
		Function:               controllers.DislikePost,
		authenticationRequired: true,
	},
}

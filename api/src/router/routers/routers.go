package routers

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Router is the structure of API routes
type Router struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	authenticationRequired bool
}

// Configure put all routes inside the router
func Configure(r *mux.Router) *mux.Router {
	routers := userRouters
	routers = append(routers, signinRouter)
	// To append this slice of routes (that is, every item in this slice has to be a struct of route), it is necessary to use "...", so that the items inside the slice are included, and not the slice itself
	routers = append(routers, postRouters...)

	// For each route of the routes
	for _, route := range routers {

		// If the route requires authentication
		if route.authenticationRequired {
			// Register a new route with a matcher for the URL path
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Function),
				),
			).Methods(route.Method)
		} else {
			//Register a new route with a matcher for the URL path
			//Ex: r.HandleFunc("/users", function).Methods("POST")
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}

	}

	return r
}

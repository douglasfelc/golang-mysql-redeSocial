package routers

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Router is the structure of WebApp routes
type Router struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	authenticationRequired bool
}

// Configure put all routes inside the router
func Configure(r *mux.Router) *mux.Router {
	routers := signinRouters
	// Append with "..." to get each slice item individually, and add
	routers = append(routers, userRouters...)
	routers = append(routers, feedRouter)
	routers = append(routers, postRouters...)
	routers = append(routers, signoutRouter)

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

	// Defines the assets folder
	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}

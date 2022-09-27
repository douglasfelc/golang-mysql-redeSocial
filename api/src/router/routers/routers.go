package routers

import (
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

	// For each route of the routes
	for _, route := range routers {
		//Register a new route with a matcher for the URL path
		//Ex: r.HandleFunc("/users", function).Methods("POST")
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
package router

import (
	"webapp/src/router/routers"

	"github.com/gorilla/mux"
)

// Generate will return a router with the configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.Configure(r)
}

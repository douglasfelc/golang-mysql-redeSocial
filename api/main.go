package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Server listening on port: %d", config.Port)
	// Server listening port: (Ex:":5000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

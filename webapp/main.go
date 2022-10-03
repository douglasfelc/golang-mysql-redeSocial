package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	r := router.Generate()

	fmt.Printf("Server listening on port: %d", 3000)
	// Server listening port: (Ex:":3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

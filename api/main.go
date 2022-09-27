package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")

	r := router.Generate()

	// Server listening port:
	log.Fatal(http.ListenAndServe(":5000", r))
}

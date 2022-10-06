package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Load()
	cookies.Config()
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Printf("Server listening on port: %d", config.Port)
	// Server listening port: (Ex:":3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

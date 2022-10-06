package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is the url of communication with the API
	APIURL = ""
	// Port where the web application is running
	Port = 0
	// HashKey is used to authenticate the cookie
	HashKey []byte
	//BlockKey is used to encrypt cookie data
	BlockKey []byte
)

// Load will initialize environment variables
func Load() {
	var error error

	// If unable to load environment variables
	if error = godotenv.Load(); error != nil {
		// Stops application and displays error
		log.Fatal(error)
	}

	// Convert APP_PORT(.env) string to integer
	Port, error = strconv.Atoi(os.Getenv("APP_PORT"))
	// if you can't read
	if error != nil {
		// Set default port 3000
		Port = 3000
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}

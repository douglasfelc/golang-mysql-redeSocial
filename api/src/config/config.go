package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseConnection is the connection string with MySQL
	DatabaseConnection = ""
	// Port where the API will run
	Port = 0
)

// Load will initialize environment variables
func Load() {
	var error error

	// If unable to load environment variables
	if error = godotenv.Load(); error != nil {
		// Stops application and displays error
		log.Fatal(error)
	}

	// Convert API_PORT(.env) string to integer
	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	// if you can't read
	if error != nil {
		// Set default port 9000
		Port = 9000
	}

	// Connection string with environment variables
	DatabaseConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}

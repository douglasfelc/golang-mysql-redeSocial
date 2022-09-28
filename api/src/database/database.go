package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Connect database
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.DatabaseConnection)
	if error != nil {
		return nil, error
	}

	// If the connection is not responding
	if error = db.Ping(); error != nil {
		// Close connection
		db.Close()

		return nil, error
	}

	return db, nil
}

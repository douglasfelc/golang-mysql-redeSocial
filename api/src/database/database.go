package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Connect database
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.DatabaseConnection)
	if erro != nil {
		return nil, erro
	}

	// If the connection is not responding
	if erro = db.Ping(); erro != nil {
		// Close connection
		db.Close()

		return nil, erro
	}

	return db, nil
}

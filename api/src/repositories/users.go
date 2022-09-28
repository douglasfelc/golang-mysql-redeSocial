// The database will be opened by the controller, which will call the repositories
// The functions will grab the database, and throw it into a struct
// An instance of this struct will be created with the database
// It is important because inside the struct, there are methods that will communicate directly with the database tables.
// In this way, things are isolated
// The controller will only worry about opening the connection with the bank
// Here the communication with the database tables will be made

package repositories

import (
	"api/src/models"
	"database/sql"
)

// users represents a user repository
type users struct {
	db *sql.DB
}

// NewUsersRepository create a user repository; will receive the database (*sql.DB), and will return a `users` pointer
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insert a user into the database
func (repository users) Create(user models.User) (uint64, error) {
	return 0, nil

	// Prepare statement
	statement, erro := repository.db.Prepare("INSERT INTO users(name, nick, email, password) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	// Execute the insert
	insert, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	// Get the last ID inserted in the database
	LastInsertID, erro := insert.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	// Returns the last insert id converted to uint64
	return uint64(LastInsertID), nil
}

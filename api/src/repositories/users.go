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
	"fmt"
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

	// Prepare statement
	statement, error := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	// Execute the insert
	insert, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	// Get the last ID inserted in the database
	LastInsertID, error := insert.LastInsertId()
	if error != nil {
		return 0, error
	}

	// Returns the last insert id converted to uint64
	return uint64(LastInsertID), nil
}

// Search user by name or nick
func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	// Make the request in the database
	rows, error := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick,
	)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []models.User

	// Iterates through the rows returned from the database
	for rows.Next() {
		var user models.User

		// Read the line
		if error = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}

		// Include the row user in the users slice
		users = append(users, user)
	}

	// If successful, returns users with the filter applied
	return users, nil
}

// SearchByID fetches the user by ID and returns it
func (repository users) SearchByID(ID uint64) (models.User, error) {
	// Make the request in the database
	row, error := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", ID,
	)
	if error != nil {
		// Returns empty user to match type, and error
		return models.User{}, error
	}
	defer row.Close()

	var user models.User

	// if you have line
	if row.Next() {
		// Read the line
		if error = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			// Returns empty user to match type, and error
			return models.User{}, error
		}
	}

	// If successful, returns user
	return user, nil
}

// Update user information in the database
func (repository users) Update(ID uint64, user models.User) error {
	// Prepare statement
	statement, error := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}

	return nil
}

// Delete user from database by ID
func (repository users) Delete(ID uint64) error {
	// Prepare statement
	statement, error := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(ID); error != nil {
		return error
	}

	return nil
}

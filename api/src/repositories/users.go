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

// Get user by name or nick
func (repository users) Get(nameOrNick string) ([]models.User, error) {
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

// GetByID get user by ID and returns it
func (repository users) GetByID(ID uint64) (models.User, error) {
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

// GetByEmail search for a user by email and return their ID and hashed password
func (repository users) GetByEmail(email string) (models.User, error) {
	// Make the request in the database
	row, error := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?", email,
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
		if error = row.Scan(&user.ID, &user.Password); error != nil {
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

// Follow saves in the database the information of one user following another
func (repository users) Follow(userID, followerID uint64) error {
	// Prepare statement
	// Insert with IGNORE: to ignore if this data already exists
	statement, error := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the insert
	if _, error := statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// UnFollow deletes from the database the information of one user following another
func (repository users) UnFollow(userID, followerID uint64) error {
	// Prepare statement
	statement, error := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the delete
	if _, error := statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// GetFollowers get in the database all the followers of a user
func (repository users) GetFollowers(userID uint64) ([]models.User, error) {

	// Make the request in the database
	rows, error := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt 
		FROM users u 
		INNER JOIN followers f ON u.id = f.follower_id 
		WHERE f.user_id = ?
		`, userID,
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

// GetFollowing get in the database all users who are following a user
func (repository users) GetFollowing(userID uint64) ([]models.User, error) {

	// Make the request in the database
	rows, error := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt 
		FROM users u 
		INNER JOIN followers f ON u.id = f.user_id 
		WHERE f.follower_id = ?
		`, userID,
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

// GetPassword get a user's password
func (repository users) GetPassword(userID uint64) (string, error) {
	// Make the request in the database
	row, error := repository.db.Query(
		"SELECT password FROM users WHERE id = ?", userID,
	)
	if error != nil {
		return "", error
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if error = row.Scan(&user.Password); error != nil {
			return "", error
		}
	}

	return user.Password, nil
}

// UpdatePassword update a user's password in the database
func (repository users) UpdatePassword(userID uint64, newPassword string) error {

	// Prepare statement
	statement, error := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(newPassword, userID); error != nil {
		return error
	}

	return nil
}

// WhoToFollow returns a list of users for the user to follow
func (repository users) WhoToFollow(userID uint64) ([]models.User, error) {

	// Make the request in the database
	rows, error := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt 
		FROM users u 
		WHERE u.id <> ? 
		AND NOT FIND_IN_SET(u.id, 
			(
				SELECT 
					CASE 
						WHEN LENGTH(GROUP_CONCAT(f.user_id)) > 0 
						THEN GROUP_CONCAT(f.user_id) 
						ELSE 0
					END
				FROM followers f 
				WHERE f.follower_id = ?
			)
		)
		LIMIT 5
		`, userID, userID,
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

package repositories

import (
	"api/src/models"
	"database/sql"
)

// posts represents a post repository
type posts struct {
	db *sql.DB
}

// NewPostsRepository create a post repository; will receive the database (*sql.DB), and will return a `posts` pointer
func NewPostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

// Create insert a post into the database
func (repository posts) Create(post models.Post) (uint64, error) {

	// Prepare statement
	statement, error := repository.db.Prepare("INSERT INTO posts (content, author_id) VALUES (?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	// Execute the insert
	insert, error := statement.Exec(post.Content, post.AuthorID)
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

// GetByID get post by ID and returns it
func (repository posts) GetByID(ID uint64) (models.Post, error) {
	// Make the request in the database
	row, error := repository.db.Query(`
		SELECT 
			p.id, p.content, p.author_id, p.likes, p.createdAt, 
			u.nick 
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		WHERE p.id = ?
	`, ID,
	)
	if error != nil {
		// Returns empty post to match type, and error
		return models.Post{}, error
	}
	defer row.Close()

	var post models.Post

	// if you have line
	if row.Next() {
		// Read the line
		if error = row.Scan(
			&post.ID,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			// Returns empty post to match type, and error
			return models.Post{}, error
		}
	}

	// If successful, returns post
	return post, nil
}

// Get posts from followers and from the user who made a request
func (repository posts) Get(userID uint64) ([]models.Post, error) {

	// Make the request in the database
	rows, error := repository.db.Query(`
		SELECT 
			DISTINCT 
			p.id, p.content, p.author_id, p.likes, p.createdAt, 
			u.nick 
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		INNER JOIN followers f ON p.author_id = f.user_id
		WHERE p.author_id = ? OR f.follower_id = ? 
		ORDER BY 1 DESC
	`, userID, userID,
	)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var posts []models.Post

	// Iterates through the rows returned from the database
	for rows.Next() {
		var post models.Post

		// Read the line
		if error = rows.Scan(
			&post.ID,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}

		// Include the row post in the posts slice
		posts = append(posts, post)
	}

	// If successful, returns posts with the filter applied
	return posts, nil
}

// Update post information in the database
func (repository posts) Update(postID uint64, post models.Post) error {
	// Prepare statement
	statement, error := repository.db.Prepare("UPDATE posts SET content = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(post.Content, postID); error != nil {
		return error
	}

	return nil
}

// Delete post from database by ID
func (repository posts) Delete(postID uint64) error {
	// Prepare statement
	statement, error := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(postID); error != nil {
		return error
	}

	return nil
}

// GetByUser get all posts of a user in the database
func (repository posts) GetByUser(userID uint64) ([]models.Post, error) {

	// Make the request in the database
	rows, error := repository.db.Query(`
		SELECT 
			DISTINCT 
			p.id, p.content, p.author_id, p.likes, p.createdAt, 
			u.nick 
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id 
		WHERE p.author_id = ? 
	`, userID,
	)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var posts []models.Post

	// Iterates through the rows returned from the database
	for rows.Next() {
		var post models.Post

		// Read the line
		if error = rows.Scan(
			&post.ID,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}

		// Include the row post in the posts slice
		posts = append(posts, post)
	}

	// If successful, returns posts with the filter applied
	return posts, nil
}

// Like add a like to the post
func (repository posts) Like(postID uint64) error {
	// Prepare statement
	statement, error := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(postID); error != nil {
		return error
	}

	return nil
}

// Dislike decrease a like on the post
func (repository posts) Dislike(postID uint64) error {
	// Prepare statement
	statement, error := repository.db.Prepare(`
		UPDATE posts 
		SET likes = 
			CASE 
				WHEN likes > 0 
					THEN likes - 1 
				ELSE 0 
			END
		WHERE id = ?
	`)
	if error != nil {
		return error
	}
	defer statement.Close()

	// Execute the update
	if _, error := statement.Exec(postID); error != nil {
		return error
	}

	return nil
}

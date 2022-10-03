package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post made by a user
type Post struct {
	ID         uint64    `json:id,omitempty`
	Title      string    `json:title,omitempty`
	Content    string    `json:content,omitempty`
	AuthorID   uint64    `json:authorId,omitempty`
	AuthorNick string    `json:authorNick,omitempty`
	Likes      uint64    `json:likes`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare post (format and validate)
func (post *Post) Prepare() error {
	// format post
	post.format()

	// validate post
	if error := post.validate(); error != nil {
		return error
	}

	return nil
}

// validade validates if the fields are filled
func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("Title is required")
	}

	if post.Content == "" {
		return errors.New("Content is required")
	}

	return nil
}

// format post data
func (post *Post) format() error {
	// TrimSpace remove leading and trailing whitespace
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	return nil
}

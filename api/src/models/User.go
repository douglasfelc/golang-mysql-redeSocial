package models

import (
	"errors"
	"strings"
	"time"
)

// User struct
type User struct {
	//omitempty: only pass user to JSON if the field is not blank
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare user (format and validate)
func (user *User) Prepare(stage string) error {
	// format user
	user.format()

	// validate user
	if error := user.validate(stage); error != nil {
		return error
	}

	return nil
}

// validade validates if the fields are filled
func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}

	if user.Nick == "" {
		return errors.New("Nickname is required")
	}

	if user.Email == "" {
		return errors.New("Email is required")
	}

	if stage == "newRegister" && user.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

// format user data
func (user *User) format() {
	// TrimSpace remove leading and trailing whitespace
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
	if error := user.format(stage); error != nil {
		return error
	}

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

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("Invalid email address") // Custom error return
	}

	if stage == "newRegister" && user.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

// format user data
func (user *User) format(stage string) error {
	// TrimSpace remove leading and trailing whitespace
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "newRegister" {
		passwordWithHash, error := security.Hash(user.Password)
		if error != nil {
			return error
		}

		// Convert to string
		user.Password = string(passwordWithHash)
	}

	return nil
}

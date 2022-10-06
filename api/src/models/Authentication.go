package models

// Authentication is an authentication return structure
type Authentication struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

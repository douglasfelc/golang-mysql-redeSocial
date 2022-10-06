package models

// Authentication is the structure that contains the authenticated user's id and token
type Authentication struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

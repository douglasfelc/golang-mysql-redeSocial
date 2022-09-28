package security

import "golang.org/x/crypto/bcrypt"

// Hash takes a string and returns the hashed password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword compares a password and a hash and returns if they match
func CheckPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}

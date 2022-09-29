package authentication

import (
	// Import with alias (jwt)
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken returns a token signed with the user's permissions
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() //expire in 6 hours
	permissions["userId"] = userID

	// Create a new token with the Claims defined above
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	// Sign the token with the security key
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken checks if the token passed in the request is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	// Convert to jwt read
	// Sends the tokenString as a parameter and a function that returns the SecretKey after checks if the signing method is correct (the token is automatically passed by parameter to checkSigningMethod)
	// As per jwo documentation: it's a good idea that before returning the verification key, check if the signature method is what you're expecting, as you can't sign a token using one method, and parse it using another
	token, error := jwt.Parse(tokenString, checkSigningMethod)
	if error != nil {
		return error
	}

	// If you can get the Claims, and the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

// ExtractUserID returns the userId that is saved in the token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)

	// Convert to jwt read
	// Sends the tokenString as a parameter and a function that returns the SecretKey after checks if the signing method is correct (the token is automatically passed by parameter to checkSigningMethod)
	// As per jwo documentation: it's a good idea that before returning the verification key, check if the signature method is what you're expecting, as you can't sign a token using one method, and parse it using another
	token, error := jwt.Parse(tokenString, checkSigningMethod)
	if error != nil {
		return 0, error
	}

	// If you can get the Claims, and the token is valid
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Convert "userId" to string, and after convert to uint64
		// permissions["userId"] é do tipo interface and ParseUint: wait a string
		userID, error := strconv.ParseUint(
			fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if error != nil {
			return 0, error
		}

		// If successful, got the Claims and the token is valid
		return userID, nil
	}

	return 0, errors.New("Invalid token")
}

// extractToken extract the request token
func extractToken(r *http.Request) string {
	// Get the Authorization content passed in the request header
	token := r.Header.Get("Authorization")

	// Check if two words are coming
	// Bearer + {token}
	if len(strings.Split(token, " ")) == 2 {
		// Returns 2nd word [1] = token
		return strings.Split(token, " ")[1]
	}

	return ""
}

// checkSigningMethod checks if the signing method is correct, and if so, returns the SecretKey
func checkSigningMethod(token *jwt.Token) (interface{}, error) {
	// Checks if using the signature method is correct
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
	}

	// If successful, return the SecretKey
	return config.SecretKey, nil
}

package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config uses environment variables to create SecureCookie
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save authentication information
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, error := s.Encode("redeSocialCookie", data)
	if error != nil {
		return error
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "redeSocialCookie",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read the cookie and return the values
func Read(r *http.Request) (map[string]string, error) {

	cookie, error := r.Cookie("redeSocialCookie")
	if error != nil {
		return nil, error
	}

	// Decode the cookie and map the values
	values := make(map[string]string)
	if error := s.Decode("redeSocialCookie", cookie.Value, &values); error != nil {
		return nil, error
	}

	return values, nil
}

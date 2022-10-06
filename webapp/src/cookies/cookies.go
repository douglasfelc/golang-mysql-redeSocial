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

package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// RequestWithAuthentication is used to place the token in the request
func RequestWithAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	//r = request made to WebApp

	// NewRequest to API
	request, error := http.NewRequest(method, url, data)
	if error != nil {
		return nil, error
	}

	// Read the Cookie, ignoring the error as the middleware has already verified this
	cookie, _ := cookies.Read(r)

	// Requisition header
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	responseHttp, error := client.Do(request)
	if error != nil {
		return nil, error
	}

	return responseHttp, nil
}

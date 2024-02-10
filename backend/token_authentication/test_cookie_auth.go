package tokenauthentication

import "net/http"

var _ TokenAuthentication = TestCookieAuth{}

type TestCookieAuth struct {
}

func (t TestCookieAuth) AuthenticateToken(r *http.Request) (interface{}, error) {
	// WARNING !!! THIS IS SUBJECT TO CSRF ATTACKS !!! DO NOT USE IN PRODUCTION !!!
	// This is a test implementation of token authentication that uses a cookie to store the token.

	// Get the token from the cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		return nil, err
	}

	// Decode the token
	username := cookie.Value

	return AuthenticatedUser{Subject: username}, nil
}

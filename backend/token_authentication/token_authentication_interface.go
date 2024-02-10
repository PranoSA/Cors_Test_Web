package tokenauthentication

import "net/http"

type AuthenticatedUser struct {
	Subject  string
	Issuer   string
	Audience string
	Roles    []string
}

type TokenAuthentication interface {
	AuthenticateToken(*http.Request) (interface{}, error)
}

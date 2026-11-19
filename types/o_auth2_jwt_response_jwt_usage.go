package types

import (
	"fmt"
)

// How the JWT is used at the token endpoint. 'authorization_grant' sends it as a jwt-bearer grant. 'client_authentication' sends it as a client_assertion on a client_credentials request, with the subject as client_id and scopes in the request body.
type OAuth2JwtResponseJwtUsage string

const (
	OAuth2JwtResponseJwtUsageAuthorizationGrant   OAuth2JwtResponseJwtUsage = "authorization_grant"
	OAuth2JwtResponseJwtUsageClientAuthentication OAuth2JwtResponseJwtUsage = "client_authentication"
)

func NewOAuth2JwtResponseJwtUsageFromString(s string) (OAuth2JwtResponseJwtUsage, error) {
	switch s {
	case "authorization_grant":
		return OAuth2JwtResponseJwtUsageAuthorizationGrant, nil
	case "client_authentication":
		return OAuth2JwtResponseJwtUsageClientAuthentication, nil
	}
	var t OAuth2JwtResponseJwtUsage
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OAuth2JwtResponseJwtUsage) Ptr() *OAuth2JwtResponseJwtUsage {
	return &e
}

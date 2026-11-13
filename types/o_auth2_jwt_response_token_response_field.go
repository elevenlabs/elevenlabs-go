package types

import (
	"fmt"
)

// Token field to extract from the token endpoint response.
type OAuth2JwtResponseTokenResponseField string

const (
	OAuth2JwtResponseTokenResponseFieldAccessToken OAuth2JwtResponseTokenResponseField = "access_token"
	OAuth2JwtResponseTokenResponseFieldIdToken     OAuth2JwtResponseTokenResponseField = "id_token"
)

func NewOAuth2JwtResponseTokenResponseFieldFromString(s string) (OAuth2JwtResponseTokenResponseField, error) {
	switch s {
	case "access_token":
		return OAuth2JwtResponseTokenResponseFieldAccessToken, nil
	case "id_token":
		return OAuth2JwtResponseTokenResponseFieldIdToken, nil
	}
	var t OAuth2JwtResponseTokenResponseField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OAuth2JwtResponseTokenResponseField) Ptr() *OAuth2JwtResponseTokenResponseField {
	return &e
}

package types

import (
	"fmt"
)

type UpdateOAuth2JwtRequestJwtUsage string

const (
	UpdateOAuth2JwtRequestJwtUsageAuthorizationGrant   UpdateOAuth2JwtRequestJwtUsage = "authorization_grant"
	UpdateOAuth2JwtRequestJwtUsageClientAuthentication UpdateOAuth2JwtRequestJwtUsage = "client_authentication"
)

func NewUpdateOAuth2JwtRequestJwtUsageFromString(s string) (UpdateOAuth2JwtRequestJwtUsage, error) {
	switch s {
	case "authorization_grant":
		return UpdateOAuth2JwtRequestJwtUsageAuthorizationGrant, nil
	case "client_authentication":
		return UpdateOAuth2JwtRequestJwtUsageClientAuthentication, nil
	}
	var t UpdateOAuth2JwtRequestJwtUsage
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UpdateOAuth2JwtRequestJwtUsage) Ptr() *UpdateOAuth2JwtRequestJwtUsage {
	return &e
}

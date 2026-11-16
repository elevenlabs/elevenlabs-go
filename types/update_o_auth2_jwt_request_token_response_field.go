package types

import (
	"fmt"
)

type UpdateOAuth2JwtRequestTokenResponseField string

const (
	UpdateOAuth2JwtRequestTokenResponseFieldAccessToken UpdateOAuth2JwtRequestTokenResponseField = "access_token"
	UpdateOAuth2JwtRequestTokenResponseFieldIdToken     UpdateOAuth2JwtRequestTokenResponseField = "id_token"
)

func NewUpdateOAuth2JwtRequestTokenResponseFieldFromString(s string) (UpdateOAuth2JwtRequestTokenResponseField, error) {
	switch s {
	case "access_token":
		return UpdateOAuth2JwtRequestTokenResponseFieldAccessToken, nil
	case "id_token":
		return UpdateOAuth2JwtRequestTokenResponseFieldIdToken, nil
	}
	var t UpdateOAuth2JwtRequestTokenResponseField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UpdateOAuth2JwtRequestTokenResponseField) Ptr() *UpdateOAuth2JwtRequestTokenResponseField {
	return &e
}

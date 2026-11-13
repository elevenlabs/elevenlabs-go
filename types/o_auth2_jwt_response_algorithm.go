package types

import (
	"fmt"
)

// JWT signing algorithm
type OAuth2JwtResponseAlgorithm string

const (
	OAuth2JwtResponseAlgorithmHs256 OAuth2JwtResponseAlgorithm = "HS256"
	OAuth2JwtResponseAlgorithmHs384 OAuth2JwtResponseAlgorithm = "HS384"
	OAuth2JwtResponseAlgorithmHs512 OAuth2JwtResponseAlgorithm = "HS512"
	OAuth2JwtResponseAlgorithmRs256 OAuth2JwtResponseAlgorithm = "RS256"
	OAuth2JwtResponseAlgorithmRs384 OAuth2JwtResponseAlgorithm = "RS384"
	OAuth2JwtResponseAlgorithmRs512 OAuth2JwtResponseAlgorithm = "RS512"
	OAuth2JwtResponseAlgorithmPs256 OAuth2JwtResponseAlgorithm = "PS256"
	OAuth2JwtResponseAlgorithmPs384 OAuth2JwtResponseAlgorithm = "PS384"
	OAuth2JwtResponseAlgorithmPs512 OAuth2JwtResponseAlgorithm = "PS512"
)

func NewOAuth2JwtResponseAlgorithmFromString(s string) (OAuth2JwtResponseAlgorithm, error) {
	switch s {
	case "HS256":
		return OAuth2JwtResponseAlgorithmHs256, nil
	case "HS384":
		return OAuth2JwtResponseAlgorithmHs384, nil
	case "HS512":
		return OAuth2JwtResponseAlgorithmHs512, nil
	case "RS256":
		return OAuth2JwtResponseAlgorithmRs256, nil
	case "RS384":
		return OAuth2JwtResponseAlgorithmRs384, nil
	case "RS512":
		return OAuth2JwtResponseAlgorithmRs512, nil
	case "PS256":
		return OAuth2JwtResponseAlgorithmPs256, nil
	case "PS384":
		return OAuth2JwtResponseAlgorithmPs384, nil
	case "PS512":
		return OAuth2JwtResponseAlgorithmPs512, nil
	}
	var t OAuth2JwtResponseAlgorithm
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OAuth2JwtResponseAlgorithm) Ptr() *OAuth2JwtResponseAlgorithm {
	return &e
}

package types

import (
	"fmt"
)

type UpdateOAuth2JwtRequestAlgorithm string

const (
	UpdateOAuth2JwtRequestAlgorithmHs256 UpdateOAuth2JwtRequestAlgorithm = "HS256"
	UpdateOAuth2JwtRequestAlgorithmHs384 UpdateOAuth2JwtRequestAlgorithm = "HS384"
	UpdateOAuth2JwtRequestAlgorithmHs512 UpdateOAuth2JwtRequestAlgorithm = "HS512"
	UpdateOAuth2JwtRequestAlgorithmRs256 UpdateOAuth2JwtRequestAlgorithm = "RS256"
	UpdateOAuth2JwtRequestAlgorithmRs384 UpdateOAuth2JwtRequestAlgorithm = "RS384"
	UpdateOAuth2JwtRequestAlgorithmRs512 UpdateOAuth2JwtRequestAlgorithm = "RS512"
	UpdateOAuth2JwtRequestAlgorithmPs256 UpdateOAuth2JwtRequestAlgorithm = "PS256"
	UpdateOAuth2JwtRequestAlgorithmPs384 UpdateOAuth2JwtRequestAlgorithm = "PS384"
	UpdateOAuth2JwtRequestAlgorithmPs512 UpdateOAuth2JwtRequestAlgorithm = "PS512"
)

func NewUpdateOAuth2JwtRequestAlgorithmFromString(s string) (UpdateOAuth2JwtRequestAlgorithm, error) {
	switch s {
	case "HS256":
		return UpdateOAuth2JwtRequestAlgorithmHs256, nil
	case "HS384":
		return UpdateOAuth2JwtRequestAlgorithmHs384, nil
	case "HS512":
		return UpdateOAuth2JwtRequestAlgorithmHs512, nil
	case "RS256":
		return UpdateOAuth2JwtRequestAlgorithmRs256, nil
	case "RS384":
		return UpdateOAuth2JwtRequestAlgorithmRs384, nil
	case "RS512":
		return UpdateOAuth2JwtRequestAlgorithmRs512, nil
	case "PS256":
		return UpdateOAuth2JwtRequestAlgorithmPs256, nil
	case "PS384":
		return UpdateOAuth2JwtRequestAlgorithmPs384, nil
	case "PS512":
		return UpdateOAuth2JwtRequestAlgorithmPs512, nil
	}
	var t UpdateOAuth2JwtRequestAlgorithm
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UpdateOAuth2JwtRequestAlgorithm) Ptr() *UpdateOAuth2JwtRequestAlgorithm {
	return &e
}

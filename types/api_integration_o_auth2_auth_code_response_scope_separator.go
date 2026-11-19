package types

import (
	"fmt"
)

// Separator for scopes
type ApiIntegrationOAuth2AuthCodeResponseScopeSeparator string

const (
	ApiIntegrationOAuth2AuthCodeResponseScopeSeparatorSpace ApiIntegrationOAuth2AuthCodeResponseScopeSeparator = " "
	ApiIntegrationOAuth2AuthCodeResponseScopeSeparatorComma ApiIntegrationOAuth2AuthCodeResponseScopeSeparator = ","
)

func NewApiIntegrationOAuth2AuthCodeResponseScopeSeparatorFromString(s string) (ApiIntegrationOAuth2AuthCodeResponseScopeSeparator, error) {
	switch s {
	case " ":
		return ApiIntegrationOAuth2AuthCodeResponseScopeSeparatorSpace, nil
	case ",":
		return ApiIntegrationOAuth2AuthCodeResponseScopeSeparatorComma, nil
	}
	var t ApiIntegrationOAuth2AuthCodeResponseScopeSeparator
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ApiIntegrationOAuth2AuthCodeResponseScopeSeparator) Ptr() *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator {
	return &e
}

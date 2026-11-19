package types

import (
	"fmt"
)

type SecretDependencyResourceType string

const (
	SecretDependencyResourceTypeTools        SecretDependencyResourceType = "tools"
	SecretDependencyResourceTypeAgents       SecretDependencyResourceType = "agents"
	SecretDependencyResourceTypePhoneNumbers SecretDependencyResourceType = "phone_numbers"
)

func NewSecretDependencyResourceTypeFromString(s string) (SecretDependencyResourceType, error) {
	switch s {
	case "tools":
		return SecretDependencyResourceTypeTools, nil
	case "agents":
		return SecretDependencyResourceTypeAgents, nil
	case "phone_numbers":
		return SecretDependencyResourceTypePhoneNumbers, nil
	}
	var t SecretDependencyResourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SecretDependencyResourceType) Ptr() *SecretDependencyResourceType {
	return &e
}

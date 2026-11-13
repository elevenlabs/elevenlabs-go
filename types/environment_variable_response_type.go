package types

import (
	"fmt"
)

type EnvironmentVariableResponseType string

const (
	EnvironmentVariableResponseTypeString         EnvironmentVariableResponseType = "string"
	EnvironmentVariableResponseTypeSecret         EnvironmentVariableResponseType = "secret"
	EnvironmentVariableResponseTypeAuthConnection EnvironmentVariableResponseType = "auth_connection"
)

func NewEnvironmentVariableResponseTypeFromString(s string) (EnvironmentVariableResponseType, error) {
	switch s {
	case "string":
		return EnvironmentVariableResponseTypeString, nil
	case "secret":
		return EnvironmentVariableResponseTypeSecret, nil
	case "auth_connection":
		return EnvironmentVariableResponseTypeAuthConnection, nil
	}
	var t EnvironmentVariableResponseType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EnvironmentVariableResponseType) Ptr() *EnvironmentVariableResponseType {
	return &e
}

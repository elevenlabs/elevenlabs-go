package types

import (
	"fmt"
)

type ListEnvironmentVariablesRequestType string

const (
	ListEnvironmentVariablesRequestTypeString         ListEnvironmentVariablesRequestType = "string"
	ListEnvironmentVariablesRequestTypeSecret         ListEnvironmentVariablesRequestType = "secret"
	ListEnvironmentVariablesRequestTypeAuthConnection ListEnvironmentVariablesRequestType = "auth_connection"
)

func NewListEnvironmentVariablesRequestTypeFromString(s string) (ListEnvironmentVariablesRequestType, error) {
	switch s {
	case "string":
		return ListEnvironmentVariablesRequestTypeString, nil
	case "secret":
		return ListEnvironmentVariablesRequestTypeSecret, nil
	case "auth_connection":
		return ListEnvironmentVariablesRequestTypeAuthConnection, nil
	}
	var t ListEnvironmentVariablesRequestType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListEnvironmentVariablesRequestType) Ptr() *ListEnvironmentVariablesRequestType {
	return &e
}

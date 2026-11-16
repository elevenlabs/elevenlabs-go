package types

import (
	"fmt"
)

// Which kind of regulated number a request or an issued number is for.
type ApplicationNumberType string

const (
	ApplicationNumberTypeLocal  ApplicationNumberType = "local"
	ApplicationNumberTypeMobile ApplicationNumberType = "mobile"
)

func NewApplicationNumberTypeFromString(s string) (ApplicationNumberType, error) {
	switch s {
	case "local":
		return ApplicationNumberTypeLocal, nil
	case "mobile":
		return ApplicationNumberTypeMobile, nil
	}
	var t ApplicationNumberType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ApplicationNumberType) Ptr() *ApplicationNumberType {
	return &e
}

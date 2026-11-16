package types

import (
	"fmt"
)

// The type of the principal: user, group, or service account (under 'key').
type ShareOptionResponseModelType string

const (
	ShareOptionResponseModelTypeUser  ShareOptionResponseModelType = "user"
	ShareOptionResponseModelTypeGroup ShareOptionResponseModelType = "group"
	ShareOptionResponseModelTypeKey   ShareOptionResponseModelType = "key"
)

func NewShareOptionResponseModelTypeFromString(s string) (ShareOptionResponseModelType, error) {
	switch s {
	case "user":
		return ShareOptionResponseModelTypeUser, nil
	case "group":
		return ShareOptionResponseModelTypeGroup, nil
	case "key":
		return ShareOptionResponseModelTypeKey, nil
	}
	var t ShareOptionResponseModelType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ShareOptionResponseModelType) Ptr() *ShareOptionResponseModelType {
	return &e
}

package types

import (
	"fmt"
)

type LivekitStackType string

const (
	LivekitStackTypeStandard LivekitStackType = "standard"
	LivekitStackTypeStatic   LivekitStackType = "static"
)

func NewLivekitStackTypeFromString(s string) (LivekitStackType, error) {
	switch s {
	case "standard":
		return LivekitStackTypeStandard, nil
	case "static":
		return LivekitStackTypeStatic, nil
	}
	var t LivekitStackType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LivekitStackType) Ptr() *LivekitStackType {
	return &e
}

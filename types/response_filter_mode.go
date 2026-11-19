package types

import (
	"fmt"
)

// Controls how tool responses are filtered before being visible to the agent.
type ResponseFilterMode string

const (
	ResponseFilterModeAll     ResponseFilterMode = "all"
	ResponseFilterModeAllow   ResponseFilterMode = "allow"
	ResponseFilterModeHideAll ResponseFilterMode = "hide_all"
)

func NewResponseFilterModeFromString(s string) (ResponseFilterMode, error) {
	switch s {
	case "all":
		return ResponseFilterModeAll, nil
	case "allow":
		return ResponseFilterModeAllow, nil
	case "hide_all":
		return ResponseFilterModeHideAll, nil
	}
	var t ResponseFilterMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResponseFilterMode) Ptr() *ResponseFilterMode {
	return &e
}

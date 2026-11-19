package types

import (
	"fmt"
)

type SipLogMessageDirection string

const (
	SipLogMessageDirectionIn  SipLogMessageDirection = "in"
	SipLogMessageDirectionOut SipLogMessageDirection = "out"
)

func NewSipLogMessageDirectionFromString(s string) (SipLogMessageDirection, error) {
	switch s {
	case "in":
		return SipLogMessageDirectionIn, nil
	case "out":
		return SipLogMessageDirectionOut, nil
	}
	var t SipLogMessageDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SipLogMessageDirection) Ptr() *SipLogMessageDirection {
	return &e
}

package types

import (
	"fmt"
)

type CallStatus string

const (
	CallStatusQueued    CallStatus = "queued"
	CallStatusSubmitted CallStatus = "submitted"
	CallStatusCompleted CallStatus = "completed"
	CallStatusCancelled CallStatus = "cancelled"
	CallStatusAttention CallStatus = "attention"
)

func NewCallStatusFromString(s string) (CallStatus, error) {
	switch s {
	case "queued":
		return CallStatusQueued, nil
	case "submitted":
		return CallStatusSubmitted, nil
	case "completed":
		return CallStatusCompleted, nil
	case "cancelled":
		return CallStatusCancelled, nil
	case "attention":
		return CallStatusAttention, nil
	}
	var t CallStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CallStatus) Ptr() *CallStatus {
	return &e
}

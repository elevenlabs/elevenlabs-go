package types

import (
	"fmt"
)

type AttentionStatus string

const (
	AttentionStatusOpen     AttentionStatus = "open"
	AttentionStatusResolved AttentionStatus = "resolved"
	AttentionStatusRetrying AttentionStatus = "retrying"
)

func NewAttentionStatusFromString(s string) (AttentionStatus, error) {
	switch s {
	case "open":
		return AttentionStatusOpen, nil
	case "resolved":
		return AttentionStatusResolved, nil
	case "retrying":
		return AttentionStatusRetrying, nil
	}
	var t AttentionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AttentionStatus) Ptr() *AttentionStatus {
	return &e
}

package types

import (
	"fmt"
)

type LoadMemoryEntryToolErrorStatus string

const (
	LoadMemoryEntryToolErrorStatusNotFound  LoadMemoryEntryToolErrorStatus = "not_found"
	LoadMemoryEntryToolErrorStatusInvalidId LoadMemoryEntryToolErrorStatus = "invalid_id"
)

func NewLoadMemoryEntryToolErrorStatusFromString(s string) (LoadMemoryEntryToolErrorStatus, error) {
	switch s {
	case "not_found":
		return LoadMemoryEntryToolErrorStatusNotFound, nil
	case "invalid_id":
		return LoadMemoryEntryToolErrorStatusInvalidId, nil
	}
	var t LoadMemoryEntryToolErrorStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LoadMemoryEntryToolErrorStatus) Ptr() *LoadMemoryEntryToolErrorStatus {
	return &e
}

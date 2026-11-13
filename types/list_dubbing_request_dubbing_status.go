package types

import (
	"fmt"
)

// What state the dub is currently in.
type ListDubbingRequestDubbingStatus string

const (
	ListDubbingRequestDubbingStatusDubbing ListDubbingRequestDubbingStatus = "dubbing"
	ListDubbingRequestDubbingStatusDubbed  ListDubbingRequestDubbingStatus = "dubbed"
	ListDubbingRequestDubbingStatusFailed  ListDubbingRequestDubbingStatus = "failed"
)

func NewListDubbingRequestDubbingStatusFromString(s string) (ListDubbingRequestDubbingStatus, error) {
	switch s {
	case "dubbing":
		return ListDubbingRequestDubbingStatusDubbing, nil
	case "dubbed":
		return ListDubbingRequestDubbingStatusDubbed, nil
	case "failed":
		return ListDubbingRequestDubbingStatusFailed, nil
	}
	var t ListDubbingRequestDubbingStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestDubbingStatus) Ptr() *ListDubbingRequestDubbingStatus {
	return &e
}

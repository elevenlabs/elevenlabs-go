package types

import (
	"fmt"
)

type BatchCallStatus string

const (
	BatchCallStatusPending    BatchCallStatus = "pending"
	BatchCallStatusInProgress BatchCallStatus = "in_progress"
	BatchCallStatusCompleted  BatchCallStatus = "completed"
	BatchCallStatusFailed     BatchCallStatus = "failed"
	BatchCallStatusCancelled  BatchCallStatus = "cancelled"
)

func NewBatchCallStatusFromString(s string) (BatchCallStatus, error) {
	switch s {
	case "pending":
		return BatchCallStatusPending, nil
	case "in_progress":
		return BatchCallStatusInProgress, nil
	case "completed":
		return BatchCallStatusCompleted, nil
	case "failed":
		return BatchCallStatusFailed, nil
	case "cancelled":
		return BatchCallStatusCancelled, nil
	}
	var t BatchCallStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BatchCallStatus) Ptr() *BatchCallStatus {
	return &e
}

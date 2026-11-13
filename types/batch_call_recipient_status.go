package types

import (
	"fmt"
)

type BatchCallRecipientStatus string

const (
	BatchCallRecipientStatusPending    BatchCallRecipientStatus = "pending"
	BatchCallRecipientStatusDispatched BatchCallRecipientStatus = "dispatched"
	BatchCallRecipientStatusInitiated  BatchCallRecipientStatus = "initiated"
	BatchCallRecipientStatusInProgress BatchCallRecipientStatus = "in_progress"
	BatchCallRecipientStatusCompleted  BatchCallRecipientStatus = "completed"
	BatchCallRecipientStatusFailed     BatchCallRecipientStatus = "failed"
	BatchCallRecipientStatusCancelled  BatchCallRecipientStatus = "cancelled"
	BatchCallRecipientStatusVoicemail  BatchCallRecipientStatus = "voicemail"
)

func NewBatchCallRecipientStatusFromString(s string) (BatchCallRecipientStatus, error) {
	switch s {
	case "pending":
		return BatchCallRecipientStatusPending, nil
	case "dispatched":
		return BatchCallRecipientStatusDispatched, nil
	case "initiated":
		return BatchCallRecipientStatusInitiated, nil
	case "in_progress":
		return BatchCallRecipientStatusInProgress, nil
	case "completed":
		return BatchCallRecipientStatusCompleted, nil
	case "failed":
		return BatchCallRecipientStatusFailed, nil
	case "cancelled":
		return BatchCallRecipientStatusCancelled, nil
	case "voicemail":
		return BatchCallRecipientStatusVoicemail, nil
	}
	var t BatchCallRecipientStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BatchCallRecipientStatus) Ptr() *BatchCallRecipientStatus {
	return &e
}

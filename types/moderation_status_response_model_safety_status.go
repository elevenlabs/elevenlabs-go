package types

import (
	"fmt"
)

type ModerationStatusResponseModelSafetyStatus string

const (
	ModerationStatusResponseModelSafetyStatusAppealApproved ModerationStatusResponseModelSafetyStatus = "appeal_approved"
	ModerationStatusResponseModelSafetyStatusAppealDenied   ModerationStatusResponseModelSafetyStatus = "appeal_denied"
	ModerationStatusResponseModelSafetyStatusFalsePositive  ModerationStatusResponseModelSafetyStatus = "false_positive"
)

func NewModerationStatusResponseModelSafetyStatusFromString(s string) (ModerationStatusResponseModelSafetyStatus, error) {
	switch s {
	case "appeal_approved":
		return ModerationStatusResponseModelSafetyStatusAppealApproved, nil
	case "appeal_denied":
		return ModerationStatusResponseModelSafetyStatusAppealDenied, nil
	case "false_positive":
		return ModerationStatusResponseModelSafetyStatusFalsePositive, nil
	}
	var t ModerationStatusResponseModelSafetyStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ModerationStatusResponseModelSafetyStatus) Ptr() *ModerationStatusResponseModelSafetyStatus {
	return &e
}

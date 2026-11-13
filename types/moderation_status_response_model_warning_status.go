package types

import (
	"fmt"
)

type ModerationStatusResponseModelWarningStatus string

const (
	ModerationStatusResponseModelWarningStatusWarning        ModerationStatusResponseModelWarningStatus = "warning"
	ModerationStatusResponseModelWarningStatusWarningCleared ModerationStatusResponseModelWarningStatus = "warning_cleared"
)

func NewModerationStatusResponseModelWarningStatusFromString(s string) (ModerationStatusResponseModelWarningStatus, error) {
	switch s {
	case "warning":
		return ModerationStatusResponseModelWarningStatusWarning, nil
	case "warning_cleared":
		return ModerationStatusResponseModelWarningStatusWarningCleared, nil
	}
	var t ModerationStatusResponseModelWarningStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ModerationStatusResponseModelWarningStatus) Ptr() *ModerationStatusResponseModelWarningStatus {
	return &e
}

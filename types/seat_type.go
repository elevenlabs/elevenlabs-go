package types

import (
	"fmt"
)

// Seat types for workspace members.
type SeatType string

const (
	SeatTypeWorkspaceAdmin      SeatType = "workspace_admin"
	SeatTypeWorkspaceMember     SeatType = "workspace_member"
	SeatTypeWorkspaceLiteMember SeatType = "workspace_lite_member"
)

func NewSeatTypeFromString(s string) (SeatType, error) {
	switch s {
	case "workspace_admin":
		return SeatTypeWorkspaceAdmin, nil
	case "workspace_member":
		return SeatTypeWorkspaceMember, nil
	case "workspace_lite_member":
		return SeatTypeWorkspaceLiteMember, nil
	}
	var t SeatType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SeatType) Ptr() *SeatType {
	return &e
}

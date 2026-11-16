package types

import (
	"fmt"
)

// What a time-off entry blocks: a staff member, an asset, or a location.
type TimeOffType string

const (
	TimeOffTypeStaff    TimeOffType = "staff"
	TimeOffTypeAsset    TimeOffType = "asset"
	TimeOffTypeLocation TimeOffType = "location"
)

func NewTimeOffTypeFromString(s string) (TimeOffType, error) {
	switch s {
	case "staff":
		return TimeOffTypeStaff, nil
	case "asset":
		return TimeOffTypeAsset, nil
	case "location":
		return TimeOffTypeLocation, nil
	}
	var t TimeOffType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TimeOffType) Ptr() *TimeOffType {
	return &e
}

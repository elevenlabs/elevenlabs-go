package types

import (
	"fmt"
)

// Whether the entry closes the target for the whole range or replaces its working hours.
type HolidayType string

const (
	HolidayTypeClosed   HolidayType = "closed"
	HolidayTypeOverride HolidayType = "override"
)

func NewHolidayTypeFromString(s string) (HolidayType, error) {
	switch s {
	case "closed":
		return HolidayTypeClosed, nil
	case "override":
		return HolidayTypeOverride, nil
	}
	var t HolidayType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e HolidayType) Ptr() *HolidayType {
	return &e
}

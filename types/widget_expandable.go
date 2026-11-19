package types

import (
	"fmt"
)

type WidgetExpandable string

const (
	WidgetExpandableNever   WidgetExpandable = "never"
	WidgetExpandableMobile  WidgetExpandable = "mobile"
	WidgetExpandableDesktop WidgetExpandable = "desktop"
	WidgetExpandableAlways  WidgetExpandable = "always"
)

func NewWidgetExpandableFromString(s string) (WidgetExpandable, error) {
	switch s {
	case "never":
		return WidgetExpandableNever, nil
	case "mobile":
		return WidgetExpandableMobile, nil
	case "desktop":
		return WidgetExpandableDesktop, nil
	case "always":
		return WidgetExpandableAlways, nil
	}
	var t WidgetExpandable
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WidgetExpandable) Ptr() *WidgetExpandable {
	return &e
}

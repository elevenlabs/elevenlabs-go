package types

import (
	"fmt"
)

type WidgetPlacement string

const (
	WidgetPlacementTopLeft     WidgetPlacement = "top-left"
	WidgetPlacementTop         WidgetPlacement = "top"
	WidgetPlacementTopRight    WidgetPlacement = "top-right"
	WidgetPlacementBottomLeft  WidgetPlacement = "bottom-left"
	WidgetPlacementBottom      WidgetPlacement = "bottom"
	WidgetPlacementBottomRight WidgetPlacement = "bottom-right"
)

func NewWidgetPlacementFromString(s string) (WidgetPlacement, error) {
	switch s {
	case "top-left":
		return WidgetPlacementTopLeft, nil
	case "top":
		return WidgetPlacementTop, nil
	case "top-right":
		return WidgetPlacementTopRight, nil
	case "bottom-left":
		return WidgetPlacementBottomLeft, nil
	case "bottom":
		return WidgetPlacementBottom, nil
	case "bottom-right":
		return WidgetPlacementBottomRight, nil
	}
	var t WidgetPlacement
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WidgetPlacement) Ptr() *WidgetPlacement {
	return &e
}

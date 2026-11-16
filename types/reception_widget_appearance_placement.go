package types

import (
	"fmt"
)

// Where on the page the widget sits.
type ReceptionWidgetAppearancePlacement string

const (
	ReceptionWidgetAppearancePlacementTopLeft     ReceptionWidgetAppearancePlacement = "top-left"
	ReceptionWidgetAppearancePlacementTop         ReceptionWidgetAppearancePlacement = "top"
	ReceptionWidgetAppearancePlacementTopRight    ReceptionWidgetAppearancePlacement = "top-right"
	ReceptionWidgetAppearancePlacementBottomLeft  ReceptionWidgetAppearancePlacement = "bottom-left"
	ReceptionWidgetAppearancePlacementBottom      ReceptionWidgetAppearancePlacement = "bottom"
	ReceptionWidgetAppearancePlacementBottomRight ReceptionWidgetAppearancePlacement = "bottom-right"
)

func NewReceptionWidgetAppearancePlacementFromString(s string) (ReceptionWidgetAppearancePlacement, error) {
	switch s {
	case "top-left":
		return ReceptionWidgetAppearancePlacementTopLeft, nil
	case "top":
		return ReceptionWidgetAppearancePlacementTop, nil
	case "top-right":
		return ReceptionWidgetAppearancePlacementTopRight, nil
	case "bottom-left":
		return ReceptionWidgetAppearancePlacementBottomLeft, nil
	case "bottom":
		return ReceptionWidgetAppearancePlacementBottom, nil
	case "bottom-right":
		return ReceptionWidgetAppearancePlacementBottomRight, nil
	}
	var t ReceptionWidgetAppearancePlacement
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionWidgetAppearancePlacement) Ptr() *ReceptionWidgetAppearancePlacement {
	return &e
}

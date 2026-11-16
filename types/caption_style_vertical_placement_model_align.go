package types

import (
	"fmt"
)

type CaptionStyleVerticalPlacementModelAlign string

const (
	CaptionStyleVerticalPlacementModelAlignTop    CaptionStyleVerticalPlacementModelAlign = "top"
	CaptionStyleVerticalPlacementModelAlignCenter CaptionStyleVerticalPlacementModelAlign = "center"
	CaptionStyleVerticalPlacementModelAlignBottom CaptionStyleVerticalPlacementModelAlign = "bottom"
)

func NewCaptionStyleVerticalPlacementModelAlignFromString(s string) (CaptionStyleVerticalPlacementModelAlign, error) {
	switch s {
	case "top":
		return CaptionStyleVerticalPlacementModelAlignTop, nil
	case "center":
		return CaptionStyleVerticalPlacementModelAlignCenter, nil
	case "bottom":
		return CaptionStyleVerticalPlacementModelAlignBottom, nil
	}
	var t CaptionStyleVerticalPlacementModelAlign
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleVerticalPlacementModelAlign) Ptr() *CaptionStyleVerticalPlacementModelAlign {
	return &e
}

package types

import (
	"fmt"
)

type CaptionStyleHorizontalPlacementModelAlign string

const (
	CaptionStyleHorizontalPlacementModelAlignLeft   CaptionStyleHorizontalPlacementModelAlign = "left"
	CaptionStyleHorizontalPlacementModelAlignCenter CaptionStyleHorizontalPlacementModelAlign = "center"
	CaptionStyleHorizontalPlacementModelAlignRight  CaptionStyleHorizontalPlacementModelAlign = "right"
)

func NewCaptionStyleHorizontalPlacementModelAlignFromString(s string) (CaptionStyleHorizontalPlacementModelAlign, error) {
	switch s {
	case "left":
		return CaptionStyleHorizontalPlacementModelAlignLeft, nil
	case "center":
		return CaptionStyleHorizontalPlacementModelAlignCenter, nil
	case "right":
		return CaptionStyleHorizontalPlacementModelAlignRight, nil
	}
	var t CaptionStyleHorizontalPlacementModelAlign
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleHorizontalPlacementModelAlign) Ptr() *CaptionStyleHorizontalPlacementModelAlign {
	return &e
}

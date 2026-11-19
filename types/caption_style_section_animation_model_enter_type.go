package types

import (
	"fmt"
)

type CaptionStyleSectionAnimationModelEnterType string

const (
	CaptionStyleSectionAnimationModelEnterTypeNone      CaptionStyleSectionAnimationModelEnterType = "none"
	CaptionStyleSectionAnimationModelEnterTypeFade      CaptionStyleSectionAnimationModelEnterType = "fade"
	CaptionStyleSectionAnimationModelEnterTypeScale     CaptionStyleSectionAnimationModelEnterType = "scale"
	CaptionStyleSectionAnimationModelEnterTypePop       CaptionStyleSectionAnimationModelEnterType = "pop"
	CaptionStyleSectionAnimationModelEnterTypeSlideUp   CaptionStyleSectionAnimationModelEnterType = "slide_up"
	CaptionStyleSectionAnimationModelEnterTypeSlideDown CaptionStyleSectionAnimationModelEnterType = "slide_down"
	CaptionStyleSectionAnimationModelEnterTypeSlam      CaptionStyleSectionAnimationModelEnterType = "slam"
	CaptionStyleSectionAnimationModelEnterTypeScaleDown CaptionStyleSectionAnimationModelEnterType = "scale_down"
	CaptionStyleSectionAnimationModelEnterTypeSlideIn   CaptionStyleSectionAnimationModelEnterType = "slide_in"
)

func NewCaptionStyleSectionAnimationModelEnterTypeFromString(s string) (CaptionStyleSectionAnimationModelEnterType, error) {
	switch s {
	case "none":
		return CaptionStyleSectionAnimationModelEnterTypeNone, nil
	case "fade":
		return CaptionStyleSectionAnimationModelEnterTypeFade, nil
	case "scale":
		return CaptionStyleSectionAnimationModelEnterTypeScale, nil
	case "pop":
		return CaptionStyleSectionAnimationModelEnterTypePop, nil
	case "slide_up":
		return CaptionStyleSectionAnimationModelEnterTypeSlideUp, nil
	case "slide_down":
		return CaptionStyleSectionAnimationModelEnterTypeSlideDown, nil
	case "slam":
		return CaptionStyleSectionAnimationModelEnterTypeSlam, nil
	case "scale_down":
		return CaptionStyleSectionAnimationModelEnterTypeScaleDown, nil
	case "slide_in":
		return CaptionStyleSectionAnimationModelEnterTypeSlideIn, nil
	}
	var t CaptionStyleSectionAnimationModelEnterType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleSectionAnimationModelEnterType) Ptr() *CaptionStyleSectionAnimationModelEnterType {
	return &e
}

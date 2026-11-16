package types

import (
	"fmt"
)

type CaptionStyleSectionAnimationModelExitType string

const (
	CaptionStyleSectionAnimationModelExitTypeNone      CaptionStyleSectionAnimationModelExitType = "none"
	CaptionStyleSectionAnimationModelExitTypeFade      CaptionStyleSectionAnimationModelExitType = "fade"
	CaptionStyleSectionAnimationModelExitTypeScale     CaptionStyleSectionAnimationModelExitType = "scale"
	CaptionStyleSectionAnimationModelExitTypePop       CaptionStyleSectionAnimationModelExitType = "pop"
	CaptionStyleSectionAnimationModelExitTypeSlideUp   CaptionStyleSectionAnimationModelExitType = "slide_up"
	CaptionStyleSectionAnimationModelExitTypeSlideDown CaptionStyleSectionAnimationModelExitType = "slide_down"
	CaptionStyleSectionAnimationModelExitTypeSlam      CaptionStyleSectionAnimationModelExitType = "slam"
	CaptionStyleSectionAnimationModelExitTypeScaleDown CaptionStyleSectionAnimationModelExitType = "scale_down"
	CaptionStyleSectionAnimationModelExitTypeSlideIn   CaptionStyleSectionAnimationModelExitType = "slide_in"
)

func NewCaptionStyleSectionAnimationModelExitTypeFromString(s string) (CaptionStyleSectionAnimationModelExitType, error) {
	switch s {
	case "none":
		return CaptionStyleSectionAnimationModelExitTypeNone, nil
	case "fade":
		return CaptionStyleSectionAnimationModelExitTypeFade, nil
	case "scale":
		return CaptionStyleSectionAnimationModelExitTypeScale, nil
	case "pop":
		return CaptionStyleSectionAnimationModelExitTypePop, nil
	case "slide_up":
		return CaptionStyleSectionAnimationModelExitTypeSlideUp, nil
	case "slide_down":
		return CaptionStyleSectionAnimationModelExitTypeSlideDown, nil
	case "slam":
		return CaptionStyleSectionAnimationModelExitTypeSlam, nil
	case "scale_down":
		return CaptionStyleSectionAnimationModelExitTypeScaleDown, nil
	case "slide_in":
		return CaptionStyleSectionAnimationModelExitTypeSlideIn, nil
	}
	var t CaptionStyleSectionAnimationModelExitType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleSectionAnimationModelExitType) Ptr() *CaptionStyleSectionAnimationModelExitType {
	return &e
}

package types

import (
	"fmt"
)

type CaptionStyleWordAnimationModelExitType string

const (
	CaptionStyleWordAnimationModelExitTypeNone      CaptionStyleWordAnimationModelExitType = "none"
	CaptionStyleWordAnimationModelExitTypeFade      CaptionStyleWordAnimationModelExitType = "fade"
	CaptionStyleWordAnimationModelExitTypeScale     CaptionStyleWordAnimationModelExitType = "scale"
	CaptionStyleWordAnimationModelExitTypePop       CaptionStyleWordAnimationModelExitType = "pop"
	CaptionStyleWordAnimationModelExitTypeSlideUp   CaptionStyleWordAnimationModelExitType = "slide_up"
	CaptionStyleWordAnimationModelExitTypeSlideDown CaptionStyleWordAnimationModelExitType = "slide_down"
	CaptionStyleWordAnimationModelExitTypeSlam      CaptionStyleWordAnimationModelExitType = "slam"
	CaptionStyleWordAnimationModelExitTypeScaleDown CaptionStyleWordAnimationModelExitType = "scale_down"
	CaptionStyleWordAnimationModelExitTypeSlideIn   CaptionStyleWordAnimationModelExitType = "slide_in"
)

func NewCaptionStyleWordAnimationModelExitTypeFromString(s string) (CaptionStyleWordAnimationModelExitType, error) {
	switch s {
	case "none":
		return CaptionStyleWordAnimationModelExitTypeNone, nil
	case "fade":
		return CaptionStyleWordAnimationModelExitTypeFade, nil
	case "scale":
		return CaptionStyleWordAnimationModelExitTypeScale, nil
	case "pop":
		return CaptionStyleWordAnimationModelExitTypePop, nil
	case "slide_up":
		return CaptionStyleWordAnimationModelExitTypeSlideUp, nil
	case "slide_down":
		return CaptionStyleWordAnimationModelExitTypeSlideDown, nil
	case "slam":
		return CaptionStyleWordAnimationModelExitTypeSlam, nil
	case "scale_down":
		return CaptionStyleWordAnimationModelExitTypeScaleDown, nil
	case "slide_in":
		return CaptionStyleWordAnimationModelExitTypeSlideIn, nil
	}
	var t CaptionStyleWordAnimationModelExitType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleWordAnimationModelExitType) Ptr() *CaptionStyleWordAnimationModelExitType {
	return &e
}

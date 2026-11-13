package types

import (
	"fmt"
)

type CaptionStyleWordAnimationModelEnterType string

const (
	CaptionStyleWordAnimationModelEnterTypeNone      CaptionStyleWordAnimationModelEnterType = "none"
	CaptionStyleWordAnimationModelEnterTypeFade      CaptionStyleWordAnimationModelEnterType = "fade"
	CaptionStyleWordAnimationModelEnterTypeScale     CaptionStyleWordAnimationModelEnterType = "scale"
	CaptionStyleWordAnimationModelEnterTypePop       CaptionStyleWordAnimationModelEnterType = "pop"
	CaptionStyleWordAnimationModelEnterTypeSlideUp   CaptionStyleWordAnimationModelEnterType = "slide_up"
	CaptionStyleWordAnimationModelEnterTypeSlideDown CaptionStyleWordAnimationModelEnterType = "slide_down"
	CaptionStyleWordAnimationModelEnterTypeSlam      CaptionStyleWordAnimationModelEnterType = "slam"
	CaptionStyleWordAnimationModelEnterTypeScaleDown CaptionStyleWordAnimationModelEnterType = "scale_down"
	CaptionStyleWordAnimationModelEnterTypeSlideIn   CaptionStyleWordAnimationModelEnterType = "slide_in"
)

func NewCaptionStyleWordAnimationModelEnterTypeFromString(s string) (CaptionStyleWordAnimationModelEnterType, error) {
	switch s {
	case "none":
		return CaptionStyleWordAnimationModelEnterTypeNone, nil
	case "fade":
		return CaptionStyleWordAnimationModelEnterTypeFade, nil
	case "scale":
		return CaptionStyleWordAnimationModelEnterTypeScale, nil
	case "pop":
		return CaptionStyleWordAnimationModelEnterTypePop, nil
	case "slide_up":
		return CaptionStyleWordAnimationModelEnterTypeSlideUp, nil
	case "slide_down":
		return CaptionStyleWordAnimationModelEnterTypeSlideDown, nil
	case "slam":
		return CaptionStyleWordAnimationModelEnterTypeSlam, nil
	case "scale_down":
		return CaptionStyleWordAnimationModelEnterTypeScaleDown, nil
	case "slide_in":
		return CaptionStyleWordAnimationModelEnterTypeSlideIn, nil
	}
	var t CaptionStyleWordAnimationModelEnterType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleWordAnimationModelEnterType) Ptr() *CaptionStyleWordAnimationModelEnterType {
	return &e
}

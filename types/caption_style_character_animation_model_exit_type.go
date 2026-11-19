package types

import (
	"fmt"
)

type CaptionStyleCharacterAnimationModelExitType string

const (
	CaptionStyleCharacterAnimationModelExitTypeNone CaptionStyleCharacterAnimationModelExitType = "none"
	CaptionStyleCharacterAnimationModelExitTypeFade CaptionStyleCharacterAnimationModelExitType = "fade"
)

func NewCaptionStyleCharacterAnimationModelExitTypeFromString(s string) (CaptionStyleCharacterAnimationModelExitType, error) {
	switch s {
	case "none":
		return CaptionStyleCharacterAnimationModelExitTypeNone, nil
	case "fade":
		return CaptionStyleCharacterAnimationModelExitTypeFade, nil
	}
	var t CaptionStyleCharacterAnimationModelExitType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleCharacterAnimationModelExitType) Ptr() *CaptionStyleCharacterAnimationModelExitType {
	return &e
}

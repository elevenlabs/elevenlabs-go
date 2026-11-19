package types

import (
	"fmt"
)

type CaptionStyleCharacterAnimationModelEnterType string

const (
	CaptionStyleCharacterAnimationModelEnterTypeNone       CaptionStyleCharacterAnimationModelEnterType = "none"
	CaptionStyleCharacterAnimationModelEnterTypeFade       CaptionStyleCharacterAnimationModelEnterType = "fade"
	CaptionStyleCharacterAnimationModelEnterTypeTypewriter CaptionStyleCharacterAnimationModelEnterType = "typewriter"
)

func NewCaptionStyleCharacterAnimationModelEnterTypeFromString(s string) (CaptionStyleCharacterAnimationModelEnterType, error) {
	switch s {
	case "none":
		return CaptionStyleCharacterAnimationModelEnterTypeNone, nil
	case "fade":
		return CaptionStyleCharacterAnimationModelEnterTypeFade, nil
	case "typewriter":
		return CaptionStyleCharacterAnimationModelEnterTypeTypewriter, nil
	}
	var t CaptionStyleCharacterAnimationModelEnterType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleCharacterAnimationModelEnterType) Ptr() *CaptionStyleCharacterAnimationModelEnterType {
	return &e
}

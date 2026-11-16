package types

import (
	"fmt"
)

type CaptionStyleModelTextTransform string

const (
	CaptionStyleModelTextTransformNone      CaptionStyleModelTextTransform = "none"
	CaptionStyleModelTextTransformUppercase CaptionStyleModelTextTransform = "uppercase"
)

func NewCaptionStyleModelTextTransformFromString(s string) (CaptionStyleModelTextTransform, error) {
	switch s {
	case "none":
		return CaptionStyleModelTextTransformNone, nil
	case "uppercase":
		return CaptionStyleModelTextTransformUppercase, nil
	}
	var t CaptionStyleModelTextTransform
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleModelTextTransform) Ptr() *CaptionStyleModelTextTransform {
	return &e
}

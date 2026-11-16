package types

import (
	"fmt"
)

type CaptionStyleModelTextBlendMode string

const (
	CaptionStyleModelTextBlendModeNormal     CaptionStyleModelTextBlendMode = "normal"
	CaptionStyleModelTextBlendModeDifference CaptionStyleModelTextBlendMode = "difference"
	CaptionStyleModelTextBlendModeMultiply   CaptionStyleModelTextBlendMode = "multiply"
)

func NewCaptionStyleModelTextBlendModeFromString(s string) (CaptionStyleModelTextBlendMode, error) {
	switch s {
	case "normal":
		return CaptionStyleModelTextBlendModeNormal, nil
	case "difference":
		return CaptionStyleModelTextBlendModeDifference, nil
	case "multiply":
		return CaptionStyleModelTextBlendModeMultiply, nil
	}
	var t CaptionStyleModelTextBlendMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleModelTextBlendMode) Ptr() *CaptionStyleModelTextBlendMode {
	return &e
}

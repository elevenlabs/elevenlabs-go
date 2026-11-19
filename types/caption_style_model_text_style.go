package types

import (
	"fmt"
)

type CaptionStyleModelTextStyle string

const (
	CaptionStyleModelTextStyleNormal CaptionStyleModelTextStyle = "normal"
	CaptionStyleModelTextStyleItalic CaptionStyleModelTextStyle = "italic"
)

func NewCaptionStyleModelTextStyleFromString(s string) (CaptionStyleModelTextStyle, error) {
	switch s {
	case "normal":
		return CaptionStyleModelTextStyleNormal, nil
	case "italic":
		return CaptionStyleModelTextStyleItalic, nil
	}
	var t CaptionStyleModelTextStyle
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleModelTextStyle) Ptr() *CaptionStyleModelTextStyle {
	return &e
}

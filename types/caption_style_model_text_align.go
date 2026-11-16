package types

import (
	"fmt"
)

type CaptionStyleModelTextAlign string

const (
	CaptionStyleModelTextAlignStart  CaptionStyleModelTextAlign = "start"
	CaptionStyleModelTextAlignCenter CaptionStyleModelTextAlign = "center"
	CaptionStyleModelTextAlignEnd    CaptionStyleModelTextAlign = "end"
)

func NewCaptionStyleModelTextAlignFromString(s string) (CaptionStyleModelTextAlign, error) {
	switch s {
	case "start":
		return CaptionStyleModelTextAlignStart, nil
	case "center":
		return CaptionStyleModelTextAlignCenter, nil
	case "end":
		return CaptionStyleModelTextAlignEnd, nil
	}
	var t CaptionStyleModelTextAlign
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleModelTextAlign) Ptr() *CaptionStyleModelTextAlign {
	return &e
}

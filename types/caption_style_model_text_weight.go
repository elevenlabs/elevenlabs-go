package types

import (
	"fmt"
)

type CaptionStyleModelTextWeight string

const (
	CaptionStyleModelTextWeightNormal CaptionStyleModelTextWeight = "normal"
	CaptionStyleModelTextWeightBold   CaptionStyleModelTextWeight = "bold"
	CaptionStyleModelTextWeightNine00 CaptionStyleModelTextWeight = "900"
)

func NewCaptionStyleModelTextWeightFromString(s string) (CaptionStyleModelTextWeight, error) {
	switch s {
	case "normal":
		return CaptionStyleModelTextWeightNormal, nil
	case "bold":
		return CaptionStyleModelTextWeightBold, nil
	case "900":
		return CaptionStyleModelTextWeightNine00, nil
	}
	var t CaptionStyleModelTextWeight
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CaptionStyleModelTextWeight) Ptr() *CaptionStyleModelTextWeight {
	return &e
}

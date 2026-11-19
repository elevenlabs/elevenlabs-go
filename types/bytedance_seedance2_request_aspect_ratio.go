package types

import (
	"fmt"
)

// The aspect ratio of the output video. With `auto`, the model picks an aspect ratio based on the inputs.
type BytedanceSeedance2RequestAspectRatio string

const (
	BytedanceSeedance2RequestAspectRatioAuto   BytedanceSeedance2RequestAspectRatio = "auto"
	BytedanceSeedance2RequestAspectRatioTwo19  BytedanceSeedance2RequestAspectRatio = "21:9"
	BytedanceSeedance2RequestAspectRatioOne69  BytedanceSeedance2RequestAspectRatio = "16:9"
	BytedanceSeedance2RequestAspectRatioFour3  BytedanceSeedance2RequestAspectRatio = "4:3"
	BytedanceSeedance2RequestAspectRatioOne1   BytedanceSeedance2RequestAspectRatio = "1:1"
	BytedanceSeedance2RequestAspectRatioThree4 BytedanceSeedance2RequestAspectRatio = "3:4"
	BytedanceSeedance2RequestAspectRatioNine16 BytedanceSeedance2RequestAspectRatio = "9:16"
)

func NewBytedanceSeedance2RequestAspectRatioFromString(s string) (BytedanceSeedance2RequestAspectRatio, error) {
	switch s {
	case "auto":
		return BytedanceSeedance2RequestAspectRatioAuto, nil
	case "21:9":
		return BytedanceSeedance2RequestAspectRatioTwo19, nil
	case "16:9":
		return BytedanceSeedance2RequestAspectRatioOne69, nil
	case "4:3":
		return BytedanceSeedance2RequestAspectRatioFour3, nil
	case "1:1":
		return BytedanceSeedance2RequestAspectRatioOne1, nil
	case "3:4":
		return BytedanceSeedance2RequestAspectRatioThree4, nil
	case "9:16":
		return BytedanceSeedance2RequestAspectRatioNine16, nil
	}
	var t BytedanceSeedance2RequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance2RequestAspectRatio) Ptr() *BytedanceSeedance2RequestAspectRatio {
	return &e
}

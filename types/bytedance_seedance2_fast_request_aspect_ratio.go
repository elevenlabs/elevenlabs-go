package types

import (
	"fmt"
)

// The aspect ratio of the output video. With `auto`, the model picks an aspect ratio based on the inputs.
type BytedanceSeedance2FastRequestAspectRatio string

const (
	BytedanceSeedance2FastRequestAspectRatioAuto   BytedanceSeedance2FastRequestAspectRatio = "auto"
	BytedanceSeedance2FastRequestAspectRatioTwo19  BytedanceSeedance2FastRequestAspectRatio = "21:9"
	BytedanceSeedance2FastRequestAspectRatioOne69  BytedanceSeedance2FastRequestAspectRatio = "16:9"
	BytedanceSeedance2FastRequestAspectRatioFour3  BytedanceSeedance2FastRequestAspectRatio = "4:3"
	BytedanceSeedance2FastRequestAspectRatioOne1   BytedanceSeedance2FastRequestAspectRatio = "1:1"
	BytedanceSeedance2FastRequestAspectRatioThree4 BytedanceSeedance2FastRequestAspectRatio = "3:4"
	BytedanceSeedance2FastRequestAspectRatioNine16 BytedanceSeedance2FastRequestAspectRatio = "9:16"
)

func NewBytedanceSeedance2FastRequestAspectRatioFromString(s string) (BytedanceSeedance2FastRequestAspectRatio, error) {
	switch s {
	case "auto":
		return BytedanceSeedance2FastRequestAspectRatioAuto, nil
	case "21:9":
		return BytedanceSeedance2FastRequestAspectRatioTwo19, nil
	case "16:9":
		return BytedanceSeedance2FastRequestAspectRatioOne69, nil
	case "4:3":
		return BytedanceSeedance2FastRequestAspectRatioFour3, nil
	case "1:1":
		return BytedanceSeedance2FastRequestAspectRatioOne1, nil
	case "3:4":
		return BytedanceSeedance2FastRequestAspectRatioThree4, nil
	case "9:16":
		return BytedanceSeedance2FastRequestAspectRatioNine16, nil
	}
	var t BytedanceSeedance2FastRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance2FastRequestAspectRatio) Ptr() *BytedanceSeedance2FastRequestAspectRatio {
	return &e
}

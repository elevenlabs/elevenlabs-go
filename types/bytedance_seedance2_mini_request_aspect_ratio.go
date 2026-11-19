package types

import (
	"fmt"
)

// The aspect ratio of the output video. With `auto`, the model picks an aspect ratio based on the inputs.
type BytedanceSeedance2MiniRequestAspectRatio string

const (
	BytedanceSeedance2MiniRequestAspectRatioAuto   BytedanceSeedance2MiniRequestAspectRatio = "auto"
	BytedanceSeedance2MiniRequestAspectRatioTwo19  BytedanceSeedance2MiniRequestAspectRatio = "21:9"
	BytedanceSeedance2MiniRequestAspectRatioOne69  BytedanceSeedance2MiniRequestAspectRatio = "16:9"
	BytedanceSeedance2MiniRequestAspectRatioFour3  BytedanceSeedance2MiniRequestAspectRatio = "4:3"
	BytedanceSeedance2MiniRequestAspectRatioOne1   BytedanceSeedance2MiniRequestAspectRatio = "1:1"
	BytedanceSeedance2MiniRequestAspectRatioThree4 BytedanceSeedance2MiniRequestAspectRatio = "3:4"
	BytedanceSeedance2MiniRequestAspectRatioNine16 BytedanceSeedance2MiniRequestAspectRatio = "9:16"
)

func NewBytedanceSeedance2MiniRequestAspectRatioFromString(s string) (BytedanceSeedance2MiniRequestAspectRatio, error) {
	switch s {
	case "auto":
		return BytedanceSeedance2MiniRequestAspectRatioAuto, nil
	case "21:9":
		return BytedanceSeedance2MiniRequestAspectRatioTwo19, nil
	case "16:9":
		return BytedanceSeedance2MiniRequestAspectRatioOne69, nil
	case "4:3":
		return BytedanceSeedance2MiniRequestAspectRatioFour3, nil
	case "1:1":
		return BytedanceSeedance2MiniRequestAspectRatioOne1, nil
	case "3:4":
		return BytedanceSeedance2MiniRequestAspectRatioThree4, nil
	case "9:16":
		return BytedanceSeedance2MiniRequestAspectRatioNine16, nil
	}
	var t BytedanceSeedance2MiniRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance2MiniRequestAspectRatio) Ptr() *BytedanceSeedance2MiniRequestAspectRatio {
	return &e
}

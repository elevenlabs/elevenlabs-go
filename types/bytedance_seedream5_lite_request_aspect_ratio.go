package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type BytedanceSeedream5LiteRequestAspectRatio string

const (
	BytedanceSeedream5LiteRequestAspectRatioAuto   BytedanceSeedream5LiteRequestAspectRatio = "auto"
	BytedanceSeedream5LiteRequestAspectRatioOne1   BytedanceSeedream5LiteRequestAspectRatio = "1:1"
	BytedanceSeedream5LiteRequestAspectRatioThree4 BytedanceSeedream5LiteRequestAspectRatio = "3:4"
	BytedanceSeedream5LiteRequestAspectRatioOne69  BytedanceSeedream5LiteRequestAspectRatio = "16:9"
	BytedanceSeedream5LiteRequestAspectRatioFour3  BytedanceSeedream5LiteRequestAspectRatio = "4:3"
	BytedanceSeedream5LiteRequestAspectRatioNine16 BytedanceSeedream5LiteRequestAspectRatio = "9:16"
)

func NewBytedanceSeedream5LiteRequestAspectRatioFromString(s string) (BytedanceSeedream5LiteRequestAspectRatio, error) {
	switch s {
	case "auto":
		return BytedanceSeedream5LiteRequestAspectRatioAuto, nil
	case "1:1":
		return BytedanceSeedream5LiteRequestAspectRatioOne1, nil
	case "3:4":
		return BytedanceSeedream5LiteRequestAspectRatioThree4, nil
	case "16:9":
		return BytedanceSeedream5LiteRequestAspectRatioOne69, nil
	case "4:3":
		return BytedanceSeedream5LiteRequestAspectRatioFour3, nil
	case "9:16":
		return BytedanceSeedream5LiteRequestAspectRatioNine16, nil
	}
	var t BytedanceSeedream5LiteRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedream5LiteRequestAspectRatio) Ptr() *BytedanceSeedream5LiteRequestAspectRatio {
	return &e
}

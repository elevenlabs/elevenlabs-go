package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type BytedanceSeedream5ProRequestAspectRatio string

const (
	BytedanceSeedream5ProRequestAspectRatioAuto   BytedanceSeedream5ProRequestAspectRatio = "auto"
	BytedanceSeedream5ProRequestAspectRatioOne1   BytedanceSeedream5ProRequestAspectRatio = "1:1"
	BytedanceSeedream5ProRequestAspectRatioThree4 BytedanceSeedream5ProRequestAspectRatio = "3:4"
	BytedanceSeedream5ProRequestAspectRatioOne69  BytedanceSeedream5ProRequestAspectRatio = "16:9"
	BytedanceSeedream5ProRequestAspectRatioFour3  BytedanceSeedream5ProRequestAspectRatio = "4:3"
	BytedanceSeedream5ProRequestAspectRatioNine16 BytedanceSeedream5ProRequestAspectRatio = "9:16"
)

func NewBytedanceSeedream5ProRequestAspectRatioFromString(s string) (BytedanceSeedream5ProRequestAspectRatio, error) {
	switch s {
	case "auto":
		return BytedanceSeedream5ProRequestAspectRatioAuto, nil
	case "1:1":
		return BytedanceSeedream5ProRequestAspectRatioOne1, nil
	case "3:4":
		return BytedanceSeedream5ProRequestAspectRatioThree4, nil
	case "16:9":
		return BytedanceSeedream5ProRequestAspectRatioOne69, nil
	case "4:3":
		return BytedanceSeedream5ProRequestAspectRatioFour3, nil
	case "9:16":
		return BytedanceSeedream5ProRequestAspectRatioNine16, nil
	}
	var t BytedanceSeedream5ProRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedream5ProRequestAspectRatio) Ptr() *BytedanceSeedream5ProRequestAspectRatio {
	return &e
}

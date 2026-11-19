package types

import (
	"fmt"
)

// The aspect ratio of the output video. With `auto`, the model picks an aspect ratio based on the inputs. First-frame / first-and-last-frame tasks always use `auto`.
type BytedanceSeedance25RequestAspectRatio string

const (
	BytedanceSeedance25RequestAspectRatioAuto   BytedanceSeedance25RequestAspectRatio = "auto"
	BytedanceSeedance25RequestAspectRatioTwo19  BytedanceSeedance25RequestAspectRatio = "21:9"
	BytedanceSeedance25RequestAspectRatioOne69  BytedanceSeedance25RequestAspectRatio = "16:9"
	BytedanceSeedance25RequestAspectRatioFour3  BytedanceSeedance25RequestAspectRatio = "4:3"
	BytedanceSeedance25RequestAspectRatioOne1   BytedanceSeedance25RequestAspectRatio = "1:1"
	BytedanceSeedance25RequestAspectRatioThree4 BytedanceSeedance25RequestAspectRatio = "3:4"
	BytedanceSeedance25RequestAspectRatioNine16 BytedanceSeedance25RequestAspectRatio = "9:16"
)

func NewBytedanceSeedance25RequestAspectRatioFromString(s string) (BytedanceSeedance25RequestAspectRatio, error) {
	switch s {
	case "auto":
		return BytedanceSeedance25RequestAspectRatioAuto, nil
	case "21:9":
		return BytedanceSeedance25RequestAspectRatioTwo19, nil
	case "16:9":
		return BytedanceSeedance25RequestAspectRatioOne69, nil
	case "4:3":
		return BytedanceSeedance25RequestAspectRatioFour3, nil
	case "1:1":
		return BytedanceSeedance25RequestAspectRatioOne1, nil
	case "3:4":
		return BytedanceSeedance25RequestAspectRatioThree4, nil
	case "9:16":
		return BytedanceSeedance25RequestAspectRatioNine16, nil
	}
	var t BytedanceSeedance25RequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance25RequestAspectRatio) Ptr() *BytedanceSeedance25RequestAspectRatio {
	return &e
}

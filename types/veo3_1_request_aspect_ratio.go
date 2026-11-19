package types

import (
	"fmt"
)

// The aspect ratio of the output video.
type Veo31RequestAspectRatio string

const (
	Veo31RequestAspectRatioOne69  Veo31RequestAspectRatio = "16:9"
	Veo31RequestAspectRatioNine16 Veo31RequestAspectRatio = "9:16"
)

func NewVeo31RequestAspectRatioFromString(s string) (Veo31RequestAspectRatio, error) {
	switch s {
	case "16:9":
		return Veo31RequestAspectRatioOne69, nil
	case "9:16":
		return Veo31RequestAspectRatioNine16, nil
	}
	var t Veo31RequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Veo31RequestAspectRatio) Ptr() *Veo31RequestAspectRatio {
	return &e
}

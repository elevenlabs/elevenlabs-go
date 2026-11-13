package types

import (
	"fmt"
)

// The aspect ratio of the output video.
type Veo31FastRequestAspectRatio string

const (
	Veo31FastRequestAspectRatioOne69  Veo31FastRequestAspectRatio = "16:9"
	Veo31FastRequestAspectRatioNine16 Veo31FastRequestAspectRatio = "9:16"
)

func NewVeo31FastRequestAspectRatioFromString(s string) (Veo31FastRequestAspectRatio, error) {
	switch s {
	case "16:9":
		return Veo31FastRequestAspectRatioOne69, nil
	case "9:16":
		return Veo31FastRequestAspectRatioNine16, nil
	}
	var t Veo31FastRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Veo31FastRequestAspectRatio) Ptr() *Veo31FastRequestAspectRatio {
	return &e
}

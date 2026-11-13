package types

import (
	"fmt"
)

// The background of the output image. With `auto`, the model picks the background that suits the image.
type GptImage1RequestBackground string

const (
	GptImage1RequestBackgroundTransparent GptImage1RequestBackground = "transparent"
	GptImage1RequestBackgroundOpaque      GptImage1RequestBackground = "opaque"
	GptImage1RequestBackgroundAuto        GptImage1RequestBackground = "auto"
)

func NewGptImage1RequestBackgroundFromString(s string) (GptImage1RequestBackground, error) {
	switch s {
	case "transparent":
		return GptImage1RequestBackgroundTransparent, nil
	case "opaque":
		return GptImage1RequestBackgroundOpaque, nil
	case "auto":
		return GptImage1RequestBackgroundAuto, nil
	}
	var t GptImage1RequestBackground
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage1RequestBackground) Ptr() *GptImage1RequestBackground {
	return &e
}

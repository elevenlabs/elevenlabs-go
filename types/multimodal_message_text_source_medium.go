package types

import (
	"fmt"
)

type MultimodalMessageTextSourceMedium string

const (
	MultimodalMessageTextSourceMediumAudio MultimodalMessageTextSourceMedium = "audio"
	MultimodalMessageTextSourceMediumDtmf  MultimodalMessageTextSourceMedium = "dtmf"
	MultimodalMessageTextSourceMediumText  MultimodalMessageTextSourceMedium = "text"
	MultimodalMessageTextSourceMediumImage MultimodalMessageTextSourceMedium = "image"
	MultimodalMessageTextSourceMediumFile  MultimodalMessageTextSourceMedium = "file"
)

func NewMultimodalMessageTextSourceMediumFromString(s string) (MultimodalMessageTextSourceMedium, error) {
	switch s {
	case "audio":
		return MultimodalMessageTextSourceMediumAudio, nil
	case "dtmf":
		return MultimodalMessageTextSourceMediumDtmf, nil
	case "text":
		return MultimodalMessageTextSourceMediumText, nil
	case "image":
		return MultimodalMessageTextSourceMediumImage, nil
	case "file":
		return MultimodalMessageTextSourceMediumFile, nil
	}
	var t MultimodalMessageTextSourceMedium
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MultimodalMessageTextSourceMedium) Ptr() *MultimodalMessageTextSourceMedium {
	return &e
}

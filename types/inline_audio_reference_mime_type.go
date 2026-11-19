package types

import (
	"fmt"
)

// The MIME type of the encoded audio.
type InlineAudioReferenceMimeType string

const (
	InlineAudioReferenceMimeTypeAudioMpeg InlineAudioReferenceMimeType = "audio/mpeg"
	InlineAudioReferenceMimeTypeAudioWav  InlineAudioReferenceMimeType = "audio/wav"
)

func NewInlineAudioReferenceMimeTypeFromString(s string) (InlineAudioReferenceMimeType, error) {
	switch s {
	case "audio/mpeg":
		return InlineAudioReferenceMimeTypeAudioMpeg, nil
	case "audio/wav":
		return InlineAudioReferenceMimeTypeAudioWav, nil
	}
	var t InlineAudioReferenceMimeType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InlineAudioReferenceMimeType) Ptr() *InlineAudioReferenceMimeType {
	return &e
}

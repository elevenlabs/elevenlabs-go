package types

import (
	"fmt"
)

type DirectPublishingReadResponseModelDisplayMode string

const (
	DirectPublishingReadResponseModelDisplayModeText          DirectPublishingReadResponseModelDisplayMode = "text"
	DirectPublishingReadResponseModelDisplayModeAudioOnly     DirectPublishingReadResponseModelDisplayMode = "audio-only"
	DirectPublishingReadResponseModelDisplayModeTextWithAudio DirectPublishingReadResponseModelDisplayMode = "text-with-audio"
)

func NewDirectPublishingReadResponseModelDisplayModeFromString(s string) (DirectPublishingReadResponseModelDisplayMode, error) {
	switch s {
	case "text":
		return DirectPublishingReadResponseModelDisplayModeText, nil
	case "audio-only":
		return DirectPublishingReadResponseModelDisplayModeAudioOnly, nil
	case "text-with-audio":
		return DirectPublishingReadResponseModelDisplayModeTextWithAudio, nil
	}
	var t DirectPublishingReadResponseModelDisplayMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DirectPublishingReadResponseModelDisplayMode) Ptr() *DirectPublishingReadResponseModelDisplayMode {
	return &e
}

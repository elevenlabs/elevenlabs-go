package types

import (
	"fmt"
)

type StudioClipLocatorClipType string

const (
	StudioClipLocatorClipTypeVideo         StudioClipLocatorClipType = "video"
	StudioClipLocatorClipTypeImage         StudioClipLocatorClipType = "image"
	StudioClipLocatorClipTypeExternalAudio StudioClipLocatorClipType = "external_audio"
	StudioClipLocatorClipTypeTtsNode       StudioClipLocatorClipType = "tts_node"
)

func NewStudioClipLocatorClipTypeFromString(s string) (StudioClipLocatorClipType, error) {
	switch s {
	case "video":
		return StudioClipLocatorClipTypeVideo, nil
	case "image":
		return StudioClipLocatorClipTypeImage, nil
	case "external_audio":
		return StudioClipLocatorClipTypeExternalAudio, nil
	case "tts_node":
		return StudioClipLocatorClipTypeTtsNode, nil
	}
	var t StudioClipLocatorClipType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e StudioClipLocatorClipType) Ptr() *StudioClipLocatorClipType {
	return &e
}

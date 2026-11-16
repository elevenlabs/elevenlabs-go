package types

import (
	"fmt"
)

type OpenAiAudioFormat string

const (
	OpenAiAudioFormatAudioPcm  OpenAiAudioFormat = "audio/pcm"
	OpenAiAudioFormatAudioPcmu OpenAiAudioFormat = "audio/pcmu"
)

func NewOpenAiAudioFormatFromString(s string) (OpenAiAudioFormat, error) {
	switch s {
	case "audio/pcm":
		return OpenAiAudioFormatAudioPcm, nil
	case "audio/pcmu":
		return OpenAiAudioFormatAudioPcmu, nil
	}
	var t OpenAiAudioFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OpenAiAudioFormat) Ptr() *OpenAiAudioFormat {
	return &e
}

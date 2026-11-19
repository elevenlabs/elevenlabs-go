package types

import (
	"fmt"
)

// The type of the word or sound. 'audio_event' is used for non-word sounds like laughter or footsteps.
type SpeechToTextWordResponseModelType string

const (
	SpeechToTextWordResponseModelTypeWord       SpeechToTextWordResponseModelType = "word"
	SpeechToTextWordResponseModelTypeSpacing    SpeechToTextWordResponseModelType = "spacing"
	SpeechToTextWordResponseModelTypeAudioEvent SpeechToTextWordResponseModelType = "audio_event"
)

func NewSpeechToTextWordResponseModelTypeFromString(s string) (SpeechToTextWordResponseModelType, error) {
	switch s {
	case "word":
		return SpeechToTextWordResponseModelTypeWord, nil
	case "spacing":
		return SpeechToTextWordResponseModelTypeSpacing, nil
	case "audio_event":
		return SpeechToTextWordResponseModelTypeAudioEvent, nil
	}
	var t SpeechToTextWordResponseModelType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeechToTextWordResponseModelType) Ptr() *SpeechToTextWordResponseModelType {
	return &e
}

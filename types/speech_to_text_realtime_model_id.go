package types

import (
	"fmt"
)

// The ID of the model to use for speech-to-text transcription.
type SpeechToTextRealtimeModelId string

const (
	SpeechToTextRealtimeModelIdScribeV2Realtime SpeechToTextRealtimeModelId = "scribe_v2_realtime"
	SpeechToTextRealtimeModelIdScribeV3Live     SpeechToTextRealtimeModelId = "scribe_v3_live"
)

func NewSpeechToTextRealtimeModelIdFromString(s string) (SpeechToTextRealtimeModelId, error) {
	switch s {
	case "scribe_v2_realtime":
		return SpeechToTextRealtimeModelIdScribeV2Realtime, nil
	case "scribe_v3_live":
		return SpeechToTextRealtimeModelIdScribeV3Live, nil
	}
	var t SpeechToTextRealtimeModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeechToTextRealtimeModelId) Ptr() *SpeechToTextRealtimeModelId {
	return &e
}

package types

import (
	"fmt"
)

type SpeechHistoryItemResponseSource string

const (
	SpeechHistoryItemResponseSourceTts             SpeechHistoryItemResponseSource = "TTS"
	SpeechHistoryItemResponseSourceSts             SpeechHistoryItemResponseSource = "STS"
	SpeechHistoryItemResponseSourceProjects        SpeechHistoryItemResponseSource = "Projects"
	SpeechHistoryItemResponseSourcePd              SpeechHistoryItemResponseSource = "PD"
	SpeechHistoryItemResponseSourceAn              SpeechHistoryItemResponseSource = "AN"
	SpeechHistoryItemResponseSourceDubbing         SpeechHistoryItemResponseSource = "Dubbing"
	SpeechHistoryItemResponseSourcePlayApi         SpeechHistoryItemResponseSource = "PlayAPI"
	SpeechHistoryItemResponseSourceConvAi          SpeechHistoryItemResponseSource = "ConvAI"
	SpeechHistoryItemResponseSourceVoiceGeneration SpeechHistoryItemResponseSource = "VoiceGeneration"
	SpeechHistoryItemResponseSourceInVpc           SpeechHistoryItemResponseSource = "InVPC"
	SpeechHistoryItemResponseSourceFlows           SpeechHistoryItemResponseSource = "Flows"
	SpeechHistoryItemResponseSourceVoicePreview    SpeechHistoryItemResponseSource = "VoicePreview"
)

func NewSpeechHistoryItemResponseSourceFromString(s string) (SpeechHistoryItemResponseSource, error) {
	switch s {
	case "TTS":
		return SpeechHistoryItemResponseSourceTts, nil
	case "STS":
		return SpeechHistoryItemResponseSourceSts, nil
	case "Projects":
		return SpeechHistoryItemResponseSourceProjects, nil
	case "PD":
		return SpeechHistoryItemResponseSourcePd, nil
	case "AN":
		return SpeechHistoryItemResponseSourceAn, nil
	case "Dubbing":
		return SpeechHistoryItemResponseSourceDubbing, nil
	case "PlayAPI":
		return SpeechHistoryItemResponseSourcePlayApi, nil
	case "ConvAI":
		return SpeechHistoryItemResponseSourceConvAi, nil
	case "VoiceGeneration":
		return SpeechHistoryItemResponseSourceVoiceGeneration, nil
	case "InVPC":
		return SpeechHistoryItemResponseSourceInVpc, nil
	case "Flows":
		return SpeechHistoryItemResponseSourceFlows, nil
	case "VoicePreview":
		return SpeechHistoryItemResponseSourceVoicePreview, nil
	}
	var t SpeechHistoryItemResponseSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeechHistoryItemResponseSource) Ptr() *SpeechHistoryItemResponseSource {
	return &e
}

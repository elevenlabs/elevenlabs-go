package types

import (
	"fmt"
)

// Commit strategy for speech segmentation. 'manual' (the default for all models) requires explicit commits, 'vad' automatically segments speech using silence detection, 'turn_prediction' additionally commits as soon as the model predicts the speaker has finished their turn (only for models that predict turns).
type SpeechToTextRealtimeCommitStrategy string

const (
	SpeechToTextRealtimeCommitStrategyManual         SpeechToTextRealtimeCommitStrategy = "manual"
	SpeechToTextRealtimeCommitStrategyVad            SpeechToTextRealtimeCommitStrategy = "vad"
	SpeechToTextRealtimeCommitStrategyTurnPrediction SpeechToTextRealtimeCommitStrategy = "turn_prediction"
)

func NewSpeechToTextRealtimeCommitStrategyFromString(s string) (SpeechToTextRealtimeCommitStrategy, error) {
	switch s {
	case "manual":
		return SpeechToTextRealtimeCommitStrategyManual, nil
	case "vad":
		return SpeechToTextRealtimeCommitStrategyVad, nil
	case "turn_prediction":
		return SpeechToTextRealtimeCommitStrategyTurnPrediction, nil
	}
	var t SpeechToTextRealtimeCommitStrategy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeechToTextRealtimeCommitStrategy) Ptr() *SpeechToTextRealtimeCommitStrategy {
	return &e
}

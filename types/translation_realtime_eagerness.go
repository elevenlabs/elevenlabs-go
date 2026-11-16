package types

import (
	"fmt"
)

// How eagerly translation starts while the speaker is still talking. 'high' starts translating soonest for the lowest latency. 'low' waits for more of the sentence, giving higher quality translations with more delay.
type TranslationRealtimeEagerness string

const (
	TranslationRealtimeEagernessLow    TranslationRealtimeEagerness = "low"
	TranslationRealtimeEagernessMedium TranslationRealtimeEagerness = "medium"
	TranslationRealtimeEagernessHigh   TranslationRealtimeEagerness = "high"
)

func NewTranslationRealtimeEagernessFromString(s string) (TranslationRealtimeEagerness, error) {
	switch s {
	case "low":
		return TranslationRealtimeEagernessLow, nil
	case "medium":
		return TranslationRealtimeEagernessMedium, nil
	case "high":
		return TranslationRealtimeEagernessHigh, nil
	}
	var t TranslationRealtimeEagerness
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TranslationRealtimeEagerness) Ptr() *TranslationRealtimeEagerness {
	return &e
}

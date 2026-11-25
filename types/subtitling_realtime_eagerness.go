package types

import (
	"fmt"
)

// How eagerly translation starts while the speaker is still talking. 'high' starts translating soonest, so subtitles appear sooner. 'low' waits for more of the sentence, giving higher quality translations with more delay.
type SubtitlingRealtimeEagerness string

const (
	SubtitlingRealtimeEagernessLow    SubtitlingRealtimeEagerness = "low"
	SubtitlingRealtimeEagernessMedium SubtitlingRealtimeEagerness = "medium"
	SubtitlingRealtimeEagernessHigh   SubtitlingRealtimeEagerness = "high"
)

func NewSubtitlingRealtimeEagernessFromString(s string) (SubtitlingRealtimeEagerness, error) {
	switch s {
	case "low":
		return SubtitlingRealtimeEagernessLow, nil
	case "medium":
		return SubtitlingRealtimeEagernessMedium, nil
	case "high":
		return SubtitlingRealtimeEagernessHigh, nil
	}
	var t SubtitlingRealtimeEagerness
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SubtitlingRealtimeEagerness) Ptr() *SubtitlingRealtimeEagerness {
	return &e
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ListSpeechEnginesResponse struct {
	// The speech engines matching the query
	SpeechEngines []*SpeechEngineSummaryResponse `json:"speech_engines" xml:"speech_engines" required:"true"`
	// Cursor for fetching the next page
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether there are more results
	HasMore         bool `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListSpeechEnginesResponse) GetSpeechEngines() (value []*SpeechEngineSummaryResponse) {
	if l == nil {
		return
	}
	return l.SpeechEngines
}

func (l *ListSpeechEnginesResponse) GetNextCursor() (value *string) {
	if l == nil {
		return
	}
	return l.NextCursor
}

func (l *ListSpeechEnginesResponse) GetHasMore() (value bool) {
	if l == nil {
		return
	}
	return l.HasMore
}

func (l *ListSpeechEnginesResponse) SetSpeechEngines(speechEngines []*SpeechEngineSummaryResponse) {
	l.SpeechEngines = speechEngines
}

func (l *ListSpeechEnginesResponse) SetNextCursor(nextCursor *string) {
	l.NextCursor = nextCursor
}

func (l *ListSpeechEnginesResponse) SetHasMore(hasMore bool) {
	l.HasMore = hasMore
}

func (l *ListSpeechEnginesResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListSpeechEnginesResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListSpeechEnginesResponse to string"
	}
	return string(jsonData)
}

func (l *ListSpeechEnginesResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListSpeechEnginesResponse) MarshalJSON() ([]byte, error) {
	type embed ListSpeechEnginesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

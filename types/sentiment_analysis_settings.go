package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SentimentAnalysisSettings struct {
	extraProperties map[string]interface{}
}

func (s *SentimentAnalysisSettings) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SentimentAnalysisSettings) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SentimentAnalysisSettings to string"
	}
	return string(jsonData)
}

func (s *SentimentAnalysisSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler SentimentAnalysisSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SentimentAnalysisSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SentimentAnalysisSettings) MarshalJSON() ([]byte, error) {
	type embed SentimentAnalysisSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

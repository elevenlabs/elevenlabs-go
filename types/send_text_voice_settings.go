package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SendTextVoiceSettings struct {
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	Style           *float64 `json:"style,omitempty" xml:"style,omitempty"`
	UseSpeakerBoost *bool    `json:"use_speaker_boost,omitempty" xml:"use_speaker_boost,omitempty"`
	Speed           *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SendTextVoiceSettings) GetStability() (value *float64) {
	if s == nil {
		return
	}
	return s.Stability
}

func (s *SendTextVoiceSettings) GetSimilarityBoost() (value *float64) {
	if s == nil {
		return
	}
	return s.SimilarityBoost
}

func (s *SendTextVoiceSettings) GetStyle() (value *float64) {
	if s == nil {
		return
	}
	return s.Style
}

func (s *SendTextVoiceSettings) GetUseSpeakerBoost() (value *bool) {
	if s == nil {
		return
	}
	return s.UseSpeakerBoost
}

func (s *SendTextVoiceSettings) GetSpeed() (value *float64) {
	if s == nil {
		return
	}
	return s.Speed
}

func (s *SendTextVoiceSettings) SetStability(stability *float64) {
	s.Stability = stability
}

func (s *SendTextVoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	s.SimilarityBoost = similarityBoost
}

func (s *SendTextVoiceSettings) SetStyle(style *float64) {
	s.Style = style
}

func (s *SendTextVoiceSettings) SetUseSpeakerBoost(useSpeakerBoost *bool) {
	s.UseSpeakerBoost = useSpeakerBoost
}

func (s *SendTextVoiceSettings) SetSpeed(speed *float64) {
	s.Speed = speed
}

func (s *SendTextVoiceSettings) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SendTextVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendTextVoiceSettings to string"
	}
	return string(jsonData)
}

func (s *SendTextVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler SendTextVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SendTextVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SendTextVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed SendTextVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

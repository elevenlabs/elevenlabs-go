package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type InitializeConnectionMultiVoiceSettings struct {
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	Style           *float64 `json:"style,omitempty" xml:"style,omitempty"`
	UseSpeakerBoost *bool    `json:"use_speaker_boost,omitempty" xml:"use_speaker_boost,omitempty"`
	Speed           *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	extraProperties map[string]interface{}
}

func (i *InitializeConnectionMultiVoiceSettings) GetStability() (value *float64) {
	if i == nil {
		return
	}
	return i.Stability
}

func (i *InitializeConnectionMultiVoiceSettings) GetSimilarityBoost() (value *float64) {
	if i == nil {
		return
	}
	return i.SimilarityBoost
}

func (i *InitializeConnectionMultiVoiceSettings) GetStyle() (value *float64) {
	if i == nil {
		return
	}
	return i.Style
}

func (i *InitializeConnectionMultiVoiceSettings) GetUseSpeakerBoost() (value *bool) {
	if i == nil {
		return
	}
	return i.UseSpeakerBoost
}

func (i *InitializeConnectionMultiVoiceSettings) GetSpeed() (value *float64) {
	if i == nil {
		return
	}
	return i.Speed
}

func (i *InitializeConnectionMultiVoiceSettings) SetStability(stability *float64) {
	i.Stability = stability
}

func (i *InitializeConnectionMultiVoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	i.SimilarityBoost = similarityBoost
}

func (i *InitializeConnectionMultiVoiceSettings) SetStyle(style *float64) {
	i.Style = style
}

func (i *InitializeConnectionMultiVoiceSettings) SetUseSpeakerBoost(useSpeakerBoost *bool) {
	i.UseSpeakerBoost = useSpeakerBoost
}

func (i *InitializeConnectionMultiVoiceSettings) SetSpeed(speed *float64) {
	i.Speed = speed
}

func (i *InitializeConnectionMultiVoiceSettings) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionMultiVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionMultiVoiceSettings to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionMultiVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler InitializeConnectionMultiVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InitializeConnectionMultiVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnectionMultiVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionMultiVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}

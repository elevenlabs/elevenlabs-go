package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type InitializeConnectionVoiceSettings struct {
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	Style           *float64 `json:"style,omitempty" xml:"style,omitempty"`
	UseSpeakerBoost *bool    `json:"use_speaker_boost,omitempty" xml:"use_speaker_boost,omitempty"`
	Speed           *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	extraProperties map[string]interface{}
}

func (i *InitializeConnectionVoiceSettings) GetStability() (value *float64) {
	if i == nil {
		return
	}
	return i.Stability
}

func (i *InitializeConnectionVoiceSettings) GetSimilarityBoost() (value *float64) {
	if i == nil {
		return
	}
	return i.SimilarityBoost
}

func (i *InitializeConnectionVoiceSettings) GetStyle() (value *float64) {
	if i == nil {
		return
	}
	return i.Style
}

func (i *InitializeConnectionVoiceSettings) GetUseSpeakerBoost() (value *bool) {
	if i == nil {
		return
	}
	return i.UseSpeakerBoost
}

func (i *InitializeConnectionVoiceSettings) GetSpeed() (value *float64) {
	if i == nil {
		return
	}
	return i.Speed
}

func (i *InitializeConnectionVoiceSettings) SetStability(stability *float64) {
	i.Stability = stability
}

func (i *InitializeConnectionVoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	i.SimilarityBoost = similarityBoost
}

func (i *InitializeConnectionVoiceSettings) SetStyle(style *float64) {
	i.Style = style
}

func (i *InitializeConnectionVoiceSettings) SetUseSpeakerBoost(useSpeakerBoost *bool) {
	i.UseSpeakerBoost = useSpeakerBoost
}

func (i *InitializeConnectionVoiceSettings) SetSpeed(speed *float64) {
	i.Speed = speed
}

func (i *InitializeConnectionVoiceSettings) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionVoiceSettings to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler InitializeConnectionVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InitializeConnectionVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnectionVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}

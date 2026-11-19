package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type InitialiseContextVoiceSettings struct {
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	Style           *float64 `json:"style,omitempty" xml:"style,omitempty"`
	UseSpeakerBoost *bool    `json:"use_speaker_boost,omitempty" xml:"use_speaker_boost,omitempty"`
	Speed           *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	extraProperties map[string]interface{}
}

func (i *InitialiseContextVoiceSettings) GetStability() (value *float64) {
	if i == nil {
		return
	}
	return i.Stability
}

func (i *InitialiseContextVoiceSettings) GetSimilarityBoost() (value *float64) {
	if i == nil {
		return
	}
	return i.SimilarityBoost
}

func (i *InitialiseContextVoiceSettings) GetStyle() (value *float64) {
	if i == nil {
		return
	}
	return i.Style
}

func (i *InitialiseContextVoiceSettings) GetUseSpeakerBoost() (value *bool) {
	if i == nil {
		return
	}
	return i.UseSpeakerBoost
}

func (i *InitialiseContextVoiceSettings) GetSpeed() (value *float64) {
	if i == nil {
		return
	}
	return i.Speed
}

func (i *InitialiseContextVoiceSettings) SetStability(stability *float64) {
	i.Stability = stability
}

func (i *InitialiseContextVoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	i.SimilarityBoost = similarityBoost
}

func (i *InitialiseContextVoiceSettings) SetStyle(style *float64) {
	i.Style = style
}

func (i *InitialiseContextVoiceSettings) SetUseSpeakerBoost(useSpeakerBoost *bool) {
	i.UseSpeakerBoost = useSpeakerBoost
}

func (i *InitialiseContextVoiceSettings) SetSpeed(speed *float64) {
	i.Speed = speed
}

func (i *InitialiseContextVoiceSettings) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitialiseContextVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitialiseContextVoiceSettings to string"
	}
	return string(jsonData)
}

func (i *InitialiseContextVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler InitialiseContextVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InitialiseContextVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitialiseContextVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed InitialiseContextVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}

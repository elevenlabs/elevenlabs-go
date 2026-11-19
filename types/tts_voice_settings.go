package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Overrides for the voice's saved settings, applied to one generation.
type TtsVoiceSettings struct {
	// How consistent the voice stays across generations. Lower values give more expressive, varied speech.
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	// How closely the output adheres to the original voice.
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty" min:"0" max:"1"`
	// How strongly the speaking style is exaggerated.
	Style *float64 `json:"style,omitempty" xml:"style,omitempty" min:"0" max:"1"`
	// Whether to boost similarity to the original speaker, at some latency cost.
	UseSpeakerBoost *bool `json:"use_speaker_boost,omitempty" xml:"use_speaker_boost,omitempty"`
	// The speed of the generated speech, where 1.0 is the voice's natural pace.
	Speed           *float64 `json:"speed,omitempty" xml:"speed,omitempty" min:"0.7" max:"1.2"`
	extraProperties map[string]interface{}
}

func (t *TtsVoiceSettings) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TtsVoiceSettings) GetSimilarityBoost() (value *float64) {
	if t == nil {
		return
	}
	return t.SimilarityBoost
}

func (t *TtsVoiceSettings) GetStyle() (value *float64) {
	if t == nil {
		return
	}
	return t.Style
}

func (t *TtsVoiceSettings) GetUseSpeakerBoost() (value *bool) {
	if t == nil {
		return
	}
	return t.UseSpeakerBoost
}

func (t *TtsVoiceSettings) GetSpeed() (value *float64) {
	if t == nil {
		return
	}
	return t.Speed
}

func (t *TtsVoiceSettings) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *TtsVoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	t.SimilarityBoost = similarityBoost
}

func (t *TtsVoiceSettings) SetStyle(style *float64) {
	t.Style = style
}

func (t *TtsVoiceSettings) SetUseSpeakerBoost(useSpeakerBoost *bool) {
	t.UseSpeakerBoost = useSpeakerBoost
}

func (t *TtsVoiceSettings) SetSpeed(speed *float64) {
	t.Speed = speed
}

func (t *TtsVoiceSettings) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TtsVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TtsVoiceSettings to string"
	}
	return string(jsonData)
}

func (t *TtsVoiceSettings) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TtsVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed TtsVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

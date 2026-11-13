package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Overrides for the voice's saved settings, applied to one generation.
type ElevenFlashV25VoiceSettings struct {
	// How consistent the voice stays across generations. Lower values give more expressive, varied speech.
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	// How closely the output adheres to the original voice.
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty" min:"0" max:"1"`
	// The speed of the generated speech, where 1.0 is the voice's natural pace.
	Speed           *float64 `json:"speed,omitempty" xml:"speed,omitempty" min:"0.7" max:"1.2"`
	extraProperties map[string]interface{}
}

func (e *ElevenFlashV25VoiceSettings) GetStability() (value *float64) {
	if e == nil {
		return
	}
	return e.Stability
}

func (e *ElevenFlashV25VoiceSettings) GetSimilarityBoost() (value *float64) {
	if e == nil {
		return
	}
	return e.SimilarityBoost
}

func (e *ElevenFlashV25VoiceSettings) GetSpeed() (value *float64) {
	if e == nil {
		return
	}
	return e.Speed
}

func (e *ElevenFlashV25VoiceSettings) SetStability(stability *float64) {
	e.Stability = stability
}

func (e *ElevenFlashV25VoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	e.SimilarityBoost = similarityBoost
}

func (e *ElevenFlashV25VoiceSettings) SetSpeed(speed *float64) {
	e.Speed = speed
}

func (e *ElevenFlashV25VoiceSettings) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ElevenFlashV25VoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ElevenFlashV25VoiceSettings to string"
	}
	return string(jsonData)
}

func (e *ElevenFlashV25VoiceSettings) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ElevenFlashV25VoiceSettings) MarshalJSON() ([]byte, error) {
	type embed ElevenFlashV25VoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}

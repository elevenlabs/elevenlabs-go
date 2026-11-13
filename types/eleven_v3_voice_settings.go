package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Overrides for the voice's saved settings, applied to one generation.
type ElevenV3VoiceSettings struct {
	// How consistent the voice stays across generations. Lower values give more expressive, varied speech.
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	extraProperties map[string]interface{}
}

func (e *ElevenV3VoiceSettings) GetStability() (value *float64) {
	if e == nil {
		return
	}
	return e.Stability
}

func (e *ElevenV3VoiceSettings) SetStability(stability *float64) {
	e.Stability = stability
}

func (e *ElevenV3VoiceSettings) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ElevenV3VoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ElevenV3VoiceSettings to string"
	}
	return string(jsonData)
}

func (e *ElevenV3VoiceSettings) UnmarshalJSON(data []byte) error {
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

func (e *ElevenV3VoiceSettings) MarshalJSON() ([]byte, error) {
	type embed ElevenV3VoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Voice settings for the connection.
type TextToDialogueClientMessageVoiceSettings struct {
	// Determines how stable the voice is and the randomness between each generation. Lower values introduce broader emotional range for the voice. Higher values can result in a monotonous voice with limited emotion.
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	extraProperties map[string]interface{}
}

func (t *TextToDialogueClientMessageVoiceSettings) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TextToDialogueClientMessageVoiceSettings) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *TextToDialogueClientMessageVoiceSettings) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextToDialogueClientMessageVoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextToDialogueClientMessageVoiceSettings to string"
	}
	return string(jsonData)
}

func (t *TextToDialogueClientMessageVoiceSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler TextToDialogueClientMessageVoiceSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TextToDialogueClientMessageVoiceSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextToDialogueClientMessageVoiceSettings) MarshalJSON() ([]byte, error) {
	type embed TextToDialogueClientMessageVoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

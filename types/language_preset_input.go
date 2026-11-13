package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguagePresetInput struct {
	Overrides *ConversationConfigClientOverrideInput `json:"overrides" xml:"overrides" required:"true"`
	// The translation of the first message
	FirstMessageTranslation *LanguagePresetTranslation `json:"first_message_translation,omitempty" xml:"first_message_translation,omitempty"`
	// The translation of the soft timeout message
	SoftTimeoutTranslation *LanguagePresetTranslation `json:"soft_timeout_translation,omitempty" xml:"soft_timeout_translation,omitempty"`
	// The translation of the noisy environment message
	NoisyEnvironmentMessageTranslation *LanguagePresetTranslation `json:"noisy_environment_message_translation,omitempty" xml:"noisy_environment_message_translation,omitempty"`
	extraProperties                    map[string]interface{}
}

func (l *LanguagePresetInput) GetOverrides() (value *ConversationConfigClientOverrideInput) {
	if l == nil {
		return
	}
	return l.Overrides
}

func (l *LanguagePresetInput) GetFirstMessageTranslation() (value *LanguagePresetTranslation) {
	if l == nil {
		return
	}
	return l.FirstMessageTranslation
}

func (l *LanguagePresetInput) GetSoftTimeoutTranslation() (value *LanguagePresetTranslation) {
	if l == nil {
		return
	}
	return l.SoftTimeoutTranslation
}

func (l *LanguagePresetInput) GetNoisyEnvironmentMessageTranslation() (value *LanguagePresetTranslation) {
	if l == nil {
		return
	}
	return l.NoisyEnvironmentMessageTranslation
}

func (l *LanguagePresetInput) SetOverrides(overrides *ConversationConfigClientOverrideInput) {
	l.Overrides = overrides
}

func (l *LanguagePresetInput) SetFirstMessageTranslation(firstMessageTranslation *LanguagePresetTranslation) {
	l.FirstMessageTranslation = firstMessageTranslation
}

func (l *LanguagePresetInput) SetSoftTimeoutTranslation(softTimeoutTranslation *LanguagePresetTranslation) {
	l.SoftTimeoutTranslation = softTimeoutTranslation
}

func (l *LanguagePresetInput) SetNoisyEnvironmentMessageTranslation(noisyEnvironmentMessageTranslation *LanguagePresetTranslation) {
	l.NoisyEnvironmentMessageTranslation = noisyEnvironmentMessageTranslation
}

func (l *LanguagePresetInput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguagePresetInput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguagePresetInput to string"
	}
	return string(jsonData)
}

func (l *LanguagePresetInput) UnmarshalJSON(data []byte) error {
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

func (l *LanguagePresetInput) MarshalJSON() ([]byte, error) {
	type embed LanguagePresetInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

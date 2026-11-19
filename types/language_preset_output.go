package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguagePresetOutput struct {
	Overrides *ConversationConfigClientOverrideOutput `json:"overrides" xml:"overrides" required:"true"`
	// The translation of the first message
	FirstMessageTranslation *LanguagePresetTranslation `json:"first_message_translation,omitempty" xml:"first_message_translation,omitempty"`
	// The translation of the soft timeout message
	SoftTimeoutTranslation *LanguagePresetTranslation `json:"soft_timeout_translation,omitempty" xml:"soft_timeout_translation,omitempty"`
	// The translation of the noisy environment message
	NoisyEnvironmentMessageTranslation *LanguagePresetTranslation `json:"noisy_environment_message_translation,omitempty" xml:"noisy_environment_message_translation,omitempty"`
	extraProperties                    map[string]interface{}
}

func (l *LanguagePresetOutput) GetOverrides() (value *ConversationConfigClientOverrideOutput) {
	if l == nil {
		return
	}
	return l.Overrides
}

func (l *LanguagePresetOutput) GetFirstMessageTranslation() (value *LanguagePresetTranslation) {
	if l == nil {
		return
	}
	return l.FirstMessageTranslation
}

func (l *LanguagePresetOutput) GetSoftTimeoutTranslation() (value *LanguagePresetTranslation) {
	if l == nil {
		return
	}
	return l.SoftTimeoutTranslation
}

func (l *LanguagePresetOutput) GetNoisyEnvironmentMessageTranslation() (value *LanguagePresetTranslation) {
	if l == nil {
		return
	}
	return l.NoisyEnvironmentMessageTranslation
}

func (l *LanguagePresetOutput) SetOverrides(overrides *ConversationConfigClientOverrideOutput) {
	l.Overrides = overrides
}

func (l *LanguagePresetOutput) SetFirstMessageTranslation(firstMessageTranslation *LanguagePresetTranslation) {
	l.FirstMessageTranslation = firstMessageTranslation
}

func (l *LanguagePresetOutput) SetSoftTimeoutTranslation(softTimeoutTranslation *LanguagePresetTranslation) {
	l.SoftTimeoutTranslation = softTimeoutTranslation
}

func (l *LanguagePresetOutput) SetNoisyEnvironmentMessageTranslation(noisyEnvironmentMessageTranslation *LanguagePresetTranslation) {
	l.NoisyEnvironmentMessageTranslation = noisyEnvironmentMessageTranslation
}

func (l *LanguagePresetOutput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguagePresetOutput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguagePresetOutput to string"
	}
	return string(jsonData)
}

func (l *LanguagePresetOutput) UnmarshalJSON(data []byte) error {
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

func (l *LanguagePresetOutput) MarshalJSON() ([]byte, error) {
	type embed LanguagePresetOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

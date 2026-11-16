package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FlattenedConfigOverrides struct {
	Language        *string `json:"language,omitempty" xml:"language,omitempty"`
	FirstMessage    *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	Prompt          *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	VoiceId         *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FlattenedConfigOverrides) GetLanguage() (value *string) {
	if f == nil {
		return
	}
	return f.Language
}

func (f *FlattenedConfigOverrides) GetFirstMessage() (value *string) {
	if f == nil {
		return
	}
	return f.FirstMessage
}

func (f *FlattenedConfigOverrides) GetPrompt() (value *string) {
	if f == nil {
		return
	}
	return f.Prompt
}

func (f *FlattenedConfigOverrides) GetVoiceId() (value *string) {
	if f == nil {
		return
	}
	return f.VoiceId
}

func (f *FlattenedConfigOverrides) SetLanguage(language *string) {
	f.Language = language
}

func (f *FlattenedConfigOverrides) SetFirstMessage(firstMessage *string) {
	f.FirstMessage = firstMessage
}

func (f *FlattenedConfigOverrides) SetPrompt(prompt *string) {
	f.Prompt = prompt
}

func (f *FlattenedConfigOverrides) SetVoiceId(voiceId *string) {
	f.VoiceId = voiceId
}

func (f *FlattenedConfigOverrides) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FlattenedConfigOverrides) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FlattenedConfigOverrides to string"
	}
	return string(jsonData)
}

func (f *FlattenedConfigOverrides) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, f); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FlattenedConfigOverrides) MarshalJSON() ([]byte, error) {
	type embed FlattenedConfigOverrides
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}

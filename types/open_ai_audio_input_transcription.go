package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiAudioInputTranscription struct {
	// Language code of the conversation (e.g. en, es, pt-br), used for transcription, speech and the model's reply language. Defaults to en.
	Language        *string `json:"language,omitempty" xml:"language,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioInputTranscription) GetLanguage() (value *string) {
	if o == nil {
		return
	}
	return o.Language
}

func (o *OpenAiAudioInputTranscription) SetLanguage(language *string) {
	o.Language = language
}

func (o *OpenAiAudioInputTranscription) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioInputTranscription) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioInputTranscription to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioInputTranscription) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioInputTranscription) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioInputTranscription
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

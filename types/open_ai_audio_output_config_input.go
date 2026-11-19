package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiAudioOutputConfigInput struct {
	Format          *OpenAiAudioOutputFormat `json:"format,omitempty" xml:"format,omitempty"`
	Voice           *string                  `json:"voice,omitempty" xml:"voice,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioOutputConfigInput) GetFormat() (value *OpenAiAudioOutputFormat) {
	if o == nil {
		return
	}
	return o.Format
}

func (o *OpenAiAudioOutputConfigInput) GetVoice() (value *string) {
	if o == nil {
		return
	}
	return o.Voice
}

func (o *OpenAiAudioOutputConfigInput) SetFormat(format *OpenAiAudioOutputFormat) {
	o.Format = format
}

func (o *OpenAiAudioOutputConfigInput) SetVoice(voice *string) {
	o.Voice = voice
}

func (o *OpenAiAudioOutputConfigInput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioOutputConfigInput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioOutputConfigInput to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioOutputConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenAiAudioOutputConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenAiAudioOutputConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioOutputConfigInput) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioOutputConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

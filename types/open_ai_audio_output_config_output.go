package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiAudioOutputConfigOutput struct {
	Format          *OpenAiAudioOutputFormat `json:"format,omitempty" xml:"format,omitempty"`
	Voice           *string                  `json:"voice,omitempty" xml:"voice,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioOutputConfigOutput) GetFormat() (value *OpenAiAudioOutputFormat) {
	if o == nil {
		return
	}
	return o.Format
}

func (o *OpenAiAudioOutputConfigOutput) GetVoice() (value *string) {
	if o == nil {
		return
	}
	return o.Voice
}

func (o *OpenAiAudioOutputConfigOutput) SetFormat(format *OpenAiAudioOutputFormat) {
	o.Format = format
}

func (o *OpenAiAudioOutputConfigOutput) SetVoice(voice *string) {
	o.Voice = voice
}

func (o *OpenAiAudioOutputConfigOutput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioOutputConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioOutputConfigOutput to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioOutputConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenAiAudioOutputConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenAiAudioOutputConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioOutputConfigOutput) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioOutputConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

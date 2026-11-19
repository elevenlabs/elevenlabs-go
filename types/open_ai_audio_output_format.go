package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiAudioOutputFormat struct {
	Type            *OpenAiAudioFormat `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioOutputFormat) GetType() (value *OpenAiAudioFormat) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OpenAiAudioOutputFormat) SetType(type_ *OpenAiAudioFormat) {
	o.Type = type_
}

func (o *OpenAiAudioOutputFormat) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioOutputFormat) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioOutputFormat to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioOutputFormat) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenAiAudioOutputFormat
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenAiAudioOutputFormat(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioOutputFormat) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioOutputFormat
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

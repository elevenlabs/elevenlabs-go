package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiAudioInputFormat struct {
	Type *OpenAiAudioFormat `json:"type,omitempty" xml:"type,omitempty"`
	// Sample rate in Hz of `audio/pcm` input. Only 24000 is supported; G.711 formats are always 8000 and ignore this field.
	Rate            *int `json:"rate,omitempty" xml:"rate,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioInputFormat) GetType() (value *OpenAiAudioFormat) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OpenAiAudioInputFormat) GetRate() (value *int) {
	if o == nil {
		return
	}
	return o.Rate
}

func (o *OpenAiAudioInputFormat) SetType(type_ *OpenAiAudioFormat) {
	o.Type = type_
}

func (o *OpenAiAudioInputFormat) SetRate(rate *int) {
	o.Rate = rate
}

func (o *OpenAiAudioInputFormat) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioInputFormat) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioInputFormat to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioInputFormat) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenAiAudioInputFormat
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenAiAudioInputFormat(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioInputFormat) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioInputFormat
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiAudioConfigInput struct {
	Input           *OpenAiAudioInputConfigInput  `json:"input,omitempty" xml:"input,omitempty"`
	Output          *OpenAiAudioOutputConfigInput `json:"output,omitempty" xml:"output,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioConfigInput) GetInput() (value *OpenAiAudioInputConfigInput) {
	if o == nil {
		return
	}
	return o.Input
}

func (o *OpenAiAudioConfigInput) GetOutput() (value *OpenAiAudioOutputConfigInput) {
	if o == nil {
		return
	}
	return o.Output
}

func (o *OpenAiAudioConfigInput) SetInput(input *OpenAiAudioInputConfigInput) {
	o.Input = input
}

func (o *OpenAiAudioConfigInput) SetOutput(output *OpenAiAudioOutputConfigInput) {
	o.Output = output
}

func (o *OpenAiAudioConfigInput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioConfigInput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioConfigInput to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenAiAudioConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenAiAudioConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioConfigInput) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

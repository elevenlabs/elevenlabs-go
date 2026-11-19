package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiAudioConfigOutput struct {
	Input           *OpenAiAudioInputConfigOutput  `json:"input,omitempty" xml:"input,omitempty"`
	Output          *OpenAiAudioOutputConfigOutput `json:"output,omitempty" xml:"output,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioConfigOutput) GetInput() (value *OpenAiAudioInputConfigOutput) {
	if o == nil {
		return
	}
	return o.Input
}

func (o *OpenAiAudioConfigOutput) GetOutput() (value *OpenAiAudioOutputConfigOutput) {
	if o == nil {
		return
	}
	return o.Output
}

func (o *OpenAiAudioConfigOutput) SetInput(input *OpenAiAudioInputConfigOutput) {
	o.Input = input
}

func (o *OpenAiAudioConfigOutput) SetOutput(output *OpenAiAudioOutputConfigOutput) {
	o.Output = output
}

func (o *OpenAiAudioConfigOutput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioConfigOutput to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenAiAudioConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenAiAudioConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioConfigOutput) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

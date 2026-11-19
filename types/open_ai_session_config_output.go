package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiSessionConfigOutput struct {
	Type                    *string                                           `json:"type,omitempty" xml:"type,omitempty"`
	Model                   *RealtimeModel                                    `json:"model,omitempty" xml:"model,omitempty"`
	Instructions            *string                                           `json:"instructions,omitempty" xml:"instructions,omitempty"`
	OutputModalities        []string                                          `json:"output_modalities,omitempty" xml:"output_modalities,omitempty"`
	Audio                   *OpenAiAudioConfigOutput                          `json:"audio,omitempty" xml:"audio,omitempty"`
	Tools                   []*OpenAiSessionConfigOutputToolsItem             `json:"tools,omitempty" xml:"tools,omitempty"`
	ToolChoice              *OpenAiSessionConfigOutputToolChoice              `json:"tool_choice,omitempty" xml:"tool_choice,omitempty"`
	Temperature             *float64                                          `json:"temperature,omitempty" xml:"temperature,omitempty"`
	MaxResponseOutputTokens *OpenAiSessionConfigOutputMaxResponseOutputTokens `json:"max_response_output_tokens,omitempty" xml:"max_response_output_tokens,omitempty"`
	// Maximum wall-clock duration of the session in seconds. The session is closed with code 1000 once it is exceeded.
	MaxDurationSeconds *int                              `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty" min:"60" max:"7200"`
	Tracing            *OpenAiSessionConfigOutputTracing `json:"tracing,omitempty" xml:"tracing,omitempty"`
	extraProperties    map[string]interface{}
}

func (o *OpenAiSessionConfigOutput) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OpenAiSessionConfigOutput) GetModel() (value *RealtimeModel) {
	if o == nil {
		return
	}
	return o.Model
}

func (o *OpenAiSessionConfigOutput) GetInstructions() (value *string) {
	if o == nil {
		return
	}
	return o.Instructions
}

func (o *OpenAiSessionConfigOutput) GetOutputModalities() (value []string) {
	if o == nil {
		return
	}
	return o.OutputModalities
}

func (o *OpenAiSessionConfigOutput) GetAudio() (value *OpenAiAudioConfigOutput) {
	if o == nil {
		return
	}
	return o.Audio
}

func (o *OpenAiSessionConfigOutput) GetTools() (value []*OpenAiSessionConfigOutputToolsItem) {
	if o == nil {
		return
	}
	return o.Tools
}

func (o *OpenAiSessionConfigOutput) GetToolChoice() (value *OpenAiSessionConfigOutputToolChoice) {
	if o == nil {
		return
	}
	return o.ToolChoice
}

func (o *OpenAiSessionConfigOutput) GetTemperature() (value *float64) {
	if o == nil {
		return
	}
	return o.Temperature
}

func (o *OpenAiSessionConfigOutput) GetMaxResponseOutputTokens() (value *OpenAiSessionConfigOutputMaxResponseOutputTokens) {
	if o == nil {
		return
	}
	return o.MaxResponseOutputTokens
}

func (o *OpenAiSessionConfigOutput) GetMaxDurationSeconds() (value *int) {
	if o == nil {
		return
	}
	return o.MaxDurationSeconds
}

func (o *OpenAiSessionConfigOutput) GetTracing() (value *OpenAiSessionConfigOutputTracing) {
	if o == nil {
		return
	}
	return o.Tracing
}

func (o *OpenAiSessionConfigOutput) SetType(type_ *string) {
	o.Type = type_
}

func (o *OpenAiSessionConfigOutput) SetModel(model *RealtimeModel) {
	o.Model = model
}

func (o *OpenAiSessionConfigOutput) SetInstructions(instructions *string) {
	o.Instructions = instructions
}

func (o *OpenAiSessionConfigOutput) SetOutputModalities(outputModalities []string) {
	o.OutputModalities = outputModalities
}

func (o *OpenAiSessionConfigOutput) SetAudio(audio *OpenAiAudioConfigOutput) {
	o.Audio = audio
}

func (o *OpenAiSessionConfigOutput) SetTools(tools []*OpenAiSessionConfigOutputToolsItem) {
	o.Tools = tools
}

func (o *OpenAiSessionConfigOutput) SetToolChoice(toolChoice *OpenAiSessionConfigOutputToolChoice) {
	o.ToolChoice = toolChoice
}

func (o *OpenAiSessionConfigOutput) SetTemperature(temperature *float64) {
	o.Temperature = temperature
}

func (o *OpenAiSessionConfigOutput) SetMaxResponseOutputTokens(maxResponseOutputTokens *OpenAiSessionConfigOutputMaxResponseOutputTokens) {
	o.MaxResponseOutputTokens = maxResponseOutputTokens
}

func (o *OpenAiSessionConfigOutput) SetMaxDurationSeconds(maxDurationSeconds *int) {
	o.MaxDurationSeconds = maxDurationSeconds
}

func (o *OpenAiSessionConfigOutput) SetTracing(tracing *OpenAiSessionConfigOutputTracing) {
	o.Tracing = tracing
}

func (o *OpenAiSessionConfigOutput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiSessionConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiSessionConfigOutput to string"
	}
	return string(jsonData)
}

func (o *OpenAiSessionConfigOutput) UnmarshalJSON(data []byte) error {
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

func (o *OpenAiSessionConfigOutput) MarshalJSON() ([]byte, error) {
	type embed OpenAiSessionConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

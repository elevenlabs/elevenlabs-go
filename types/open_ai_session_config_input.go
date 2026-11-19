package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiSessionConfigInput struct {
	Type                    *string                                          `json:"type,omitempty" xml:"type,omitempty"`
	Model                   *RealtimeModel                                   `json:"model,omitempty" xml:"model,omitempty"`
	Instructions            *string                                          `json:"instructions,omitempty" xml:"instructions,omitempty"`
	OutputModalities        []string                                         `json:"output_modalities,omitempty" xml:"output_modalities,omitempty"`
	Audio                   *OpenAiAudioConfigInput                          `json:"audio,omitempty" xml:"audio,omitempty"`
	Tools                   []*OpenAiSessionConfigInputToolsItem             `json:"tools,omitempty" xml:"tools,omitempty"`
	ToolChoice              *OpenAiSessionConfigInputToolChoice              `json:"tool_choice,omitempty" xml:"tool_choice,omitempty"`
	Temperature             *float64                                         `json:"temperature,omitempty" xml:"temperature,omitempty"`
	MaxResponseOutputTokens *OpenAiSessionConfigInputMaxResponseOutputTokens `json:"max_response_output_tokens,omitempty" xml:"max_response_output_tokens,omitempty"`
	// Maximum wall-clock duration of the session in seconds. The session is closed with code 1000 once it is exceeded.
	MaxDurationSeconds *int                                   `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty" min:"60" max:"7200"`
	TurnDetection      *OpenAiSessionConfigInputTurnDetection `json:"turn_detection,omitempty" xml:"turn_detection,omitempty"`
	Tracing            *OpenAiSessionConfigInputTracing       `json:"tracing,omitempty" xml:"tracing,omitempty"`
	Include            []string                               `json:"include,omitempty" xml:"include,omitempty"`
	ParallelToolCalls  *bool                                  `json:"parallel_tool_calls,omitempty" xml:"parallel_tool_calls,omitempty"`
	Reasoning          map[string]any                         `json:"reasoning,omitempty"`
	Truncation         *OpenAiSessionConfigInputTruncation    `json:"truncation,omitempty" xml:"truncation,omitempty"`
	Prompt             map[string]any                         `json:"prompt,omitempty"`
	extraProperties    map[string]interface{}
}

func (o *OpenAiSessionConfigInput) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OpenAiSessionConfigInput) GetModel() (value *RealtimeModel) {
	if o == nil {
		return
	}
	return o.Model
}

func (o *OpenAiSessionConfigInput) GetInstructions() (value *string) {
	if o == nil {
		return
	}
	return o.Instructions
}

func (o *OpenAiSessionConfigInput) GetOutputModalities() (value []string) {
	if o == nil {
		return
	}
	return o.OutputModalities
}

func (o *OpenAiSessionConfigInput) GetAudio() (value *OpenAiAudioConfigInput) {
	if o == nil {
		return
	}
	return o.Audio
}

func (o *OpenAiSessionConfigInput) GetTools() (value []*OpenAiSessionConfigInputToolsItem) {
	if o == nil {
		return
	}
	return o.Tools
}

func (o *OpenAiSessionConfigInput) GetToolChoice() (value *OpenAiSessionConfigInputToolChoice) {
	if o == nil {
		return
	}
	return o.ToolChoice
}

func (o *OpenAiSessionConfigInput) GetTemperature() (value *float64) {
	if o == nil {
		return
	}
	return o.Temperature
}

func (o *OpenAiSessionConfigInput) GetMaxResponseOutputTokens() (value *OpenAiSessionConfigInputMaxResponseOutputTokens) {
	if o == nil {
		return
	}
	return o.MaxResponseOutputTokens
}

func (o *OpenAiSessionConfigInput) GetMaxDurationSeconds() (value *int) {
	if o == nil {
		return
	}
	return o.MaxDurationSeconds
}

func (o *OpenAiSessionConfigInput) GetTurnDetection() (value *OpenAiSessionConfigInputTurnDetection) {
	if o == nil {
		return
	}
	return o.TurnDetection
}

func (o *OpenAiSessionConfigInput) GetTracing() (value *OpenAiSessionConfigInputTracing) {
	if o == nil {
		return
	}
	return o.Tracing
}

func (o *OpenAiSessionConfigInput) GetInclude() (value []string) {
	if o == nil {
		return
	}
	return o.Include
}

func (o *OpenAiSessionConfigInput) GetParallelToolCalls() (value *bool) {
	if o == nil {
		return
	}
	return o.ParallelToolCalls
}

func (o *OpenAiSessionConfigInput) GetReasoning() (value map[string]any) {
	if o == nil {
		return
	}
	return o.Reasoning
}

func (o *OpenAiSessionConfigInput) GetTruncation() (value *OpenAiSessionConfigInputTruncation) {
	if o == nil {
		return
	}
	return o.Truncation
}

func (o *OpenAiSessionConfigInput) GetPrompt() (value map[string]any) {
	if o == nil {
		return
	}
	return o.Prompt
}

func (o *OpenAiSessionConfigInput) SetType(type_ *string) {
	o.Type = type_
}

func (o *OpenAiSessionConfigInput) SetModel(model *RealtimeModel) {
	o.Model = model
}

func (o *OpenAiSessionConfigInput) SetInstructions(instructions *string) {
	o.Instructions = instructions
}

func (o *OpenAiSessionConfigInput) SetOutputModalities(outputModalities []string) {
	o.OutputModalities = outputModalities
}

func (o *OpenAiSessionConfigInput) SetAudio(audio *OpenAiAudioConfigInput) {
	o.Audio = audio
}

func (o *OpenAiSessionConfigInput) SetTools(tools []*OpenAiSessionConfigInputToolsItem) {
	o.Tools = tools
}

func (o *OpenAiSessionConfigInput) SetToolChoice(toolChoice *OpenAiSessionConfigInputToolChoice) {
	o.ToolChoice = toolChoice
}

func (o *OpenAiSessionConfigInput) SetTemperature(temperature *float64) {
	o.Temperature = temperature
}

func (o *OpenAiSessionConfigInput) SetMaxResponseOutputTokens(maxResponseOutputTokens *OpenAiSessionConfigInputMaxResponseOutputTokens) {
	o.MaxResponseOutputTokens = maxResponseOutputTokens
}

func (o *OpenAiSessionConfigInput) SetMaxDurationSeconds(maxDurationSeconds *int) {
	o.MaxDurationSeconds = maxDurationSeconds
}

func (o *OpenAiSessionConfigInput) SetTurnDetection(turnDetection *OpenAiSessionConfigInputTurnDetection) {
	o.TurnDetection = turnDetection
}

func (o *OpenAiSessionConfigInput) SetTracing(tracing *OpenAiSessionConfigInputTracing) {
	o.Tracing = tracing
}

func (o *OpenAiSessionConfigInput) SetInclude(include []string) {
	o.Include = include
}

func (o *OpenAiSessionConfigInput) SetParallelToolCalls(parallelToolCalls *bool) {
	o.ParallelToolCalls = parallelToolCalls
}

func (o *OpenAiSessionConfigInput) SetReasoning(reasoning map[string]any) {
	o.Reasoning = reasoning
}

func (o *OpenAiSessionConfigInput) SetTruncation(truncation *OpenAiSessionConfigInputTruncation) {
	o.Truncation = truncation
}

func (o *OpenAiSessionConfigInput) SetPrompt(prompt map[string]any) {
	o.Prompt = prompt
}

func (o *OpenAiSessionConfigInput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiSessionConfigInput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiSessionConfigInput to string"
	}
	return string(jsonData)
}

func (o *OpenAiSessionConfigInput) UnmarshalJSON(data []byte) error {
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

func (o *OpenAiSessionConfigInput) MarshalJSON() ([]byte, error) {
	type embed OpenAiSessionConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

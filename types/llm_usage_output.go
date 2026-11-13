package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmUsageOutput struct {
	// Token usage and USD price per LLM, keyed by model ID and broken down by token category.
	ModelUsage      map[string]*LlmInputOutputTokensUsage `json:"model_usage,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmUsageOutput) GetModelUsage() (value map[string]*LlmInputOutputTokensUsage) {
	if l == nil {
		return
	}
	return l.ModelUsage
}

func (l *LlmUsageOutput) SetModelUsage(modelUsage map[string]*LlmInputOutputTokensUsage) {
	l.ModelUsage = modelUsage
}

func (l *LlmUsageOutput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmUsageOutput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmUsageOutput to string"
	}
	return string(jsonData)
}

func (l *LlmUsageOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmUsageOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmUsageOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmUsageOutput) MarshalJSON() ([]byte, error) {
	type embed LlmUsageOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

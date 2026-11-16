package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmUsageInput struct {
	// Token usage and USD price per LLM, keyed by model ID and broken down by token category.
	ModelUsage      map[string]*LlmInputOutputTokensUsage `json:"model_usage,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmUsageInput) GetModelUsage() (value map[string]*LlmInputOutputTokensUsage) {
	if l == nil {
		return
	}
	return l.ModelUsage
}

func (l *LlmUsageInput) SetModelUsage(modelUsage map[string]*LlmInputOutputTokensUsage) {
	l.ModelUsage = modelUsage
}

func (l *LlmUsageInput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmUsageInput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmUsageInput to string"
	}
	return string(jsonData)
}

func (l *LlmUsageInput) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmUsageInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmUsageInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmUsageInput) MarshalJSON() ([]byte, error) {
	type embed LlmUsageInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

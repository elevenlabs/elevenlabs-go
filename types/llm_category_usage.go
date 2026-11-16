package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// LLM usage of a conversation, split by whether the generated response reached the user.
type LlmCategoryUsage struct {
	IrreversibleGeneration *LlmUsageOutput `json:"irreversible_generation,omitempty" xml:"irreversible_generation,omitempty"`
	InitiatedGeneration    *LlmUsageOutput `json:"initiated_generation,omitempty" xml:"initiated_generation,omitempty"`
	extraProperties        map[string]interface{}
}

func (l *LlmCategoryUsage) GetIrreversibleGeneration() (value *LlmUsageOutput) {
	if l == nil {
		return
	}
	return l.IrreversibleGeneration
}

func (l *LlmCategoryUsage) GetInitiatedGeneration() (value *LlmUsageOutput) {
	if l == nil {
		return
	}
	return l.InitiatedGeneration
}

func (l *LlmCategoryUsage) SetIrreversibleGeneration(irreversibleGeneration *LlmUsageOutput) {
	l.IrreversibleGeneration = irreversibleGeneration
}

func (l *LlmCategoryUsage) SetInitiatedGeneration(initiatedGeneration *LlmUsageOutput) {
	l.InitiatedGeneration = initiatedGeneration
}

func (l *LlmCategoryUsage) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmCategoryUsage) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmCategoryUsage to string"
	}
	return string(jsonData)
}

func (l *LlmCategoryUsage) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmCategoryUsage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmCategoryUsage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmCategoryUsage) MarshalJSON() ([]byte, error) {
	type embed LlmCategoryUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

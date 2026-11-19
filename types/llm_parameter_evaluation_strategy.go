package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmParameterEvaluationStrategy struct {
	type_ string
	// A description of the evaluation strategy to use for the test.
	Description     string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LlmParameterEvaluationStrategy) GetDescription() (value string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmParameterEvaluationStrategy) SetDescription(description string) {
	l.Description = description
}

func (l *LlmParameterEvaluationStrategy) Type() string {
	return l.type_
}

func (l *LlmParameterEvaluationStrategy) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmParameterEvaluationStrategy) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmParameterEvaluationStrategy to string"
	}
	return string(jsonData)
}

func (l *LlmParameterEvaluationStrategy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LlmParameterEvaluationStrategy
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LlmParameterEvaluationStrategy(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "llm" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", l, "llm", *literals.Type)
	}
	l.type_ = "llm"
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmParameterEvaluationStrategy) MarshalJSON() ([]byte, error) {
	type embed LlmParameterEvaluationStrategy
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*l),
		Type:  "llm",
	}
	return json.Marshal(marshaler)
}

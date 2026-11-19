package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowLlmConditionModelOutput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label *string `json:"label" xml:"label" required:"true"`
	type_ string
	// Condition to evaluate
	Condition       string `json:"condition" xml:"condition" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowLlmConditionModelOutput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowLlmConditionModelOutput) GetCondition() (value string) {
	if w == nil {
		return
	}
	return w.Condition
}

func (w *WorkflowLlmConditionModelOutput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowLlmConditionModelOutput) SetCondition(condition string) {
	w.Condition = condition
}

func (w *WorkflowLlmConditionModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowLlmConditionModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowLlmConditionModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowLlmConditionModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowLlmConditionModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "llm" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "llm", *literals.Type)
	}
	w.type_ = "llm"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowLlmConditionModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowLlmConditionModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "llm",
	}
	return json.Marshal(marshaler)
}

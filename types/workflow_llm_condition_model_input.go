package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowLlmConditionModelInput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// Condition to evaluate
	Condition       string `json:"condition" xml:"condition" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowLlmConditionModelInput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowLlmConditionModelInput) GetCondition() (value string) {
	if w == nil {
		return
	}
	return w.Condition
}

func (w *WorkflowLlmConditionModelInput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowLlmConditionModelInput) SetCondition(condition string) {
	w.Condition = condition
}

func (w *WorkflowLlmConditionModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowLlmConditionModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowLlmConditionModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowLlmConditionModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowLlmConditionModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowLlmConditionModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

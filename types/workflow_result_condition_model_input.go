package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowResultConditionModelInput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// Whether all tools in the previously executed tool node were executed successfully.
	Successful      bool `json:"successful" xml:"successful" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowResultConditionModelInput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowResultConditionModelInput) GetSuccessful() (value bool) {
	if w == nil {
		return
	}
	return w.Successful
}

func (w *WorkflowResultConditionModelInput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowResultConditionModelInput) SetSuccessful(successful bool) {
	w.Successful = successful
}

func (w *WorkflowResultConditionModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowResultConditionModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowResultConditionModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowResultConditionModelInput) UnmarshalJSON(data []byte) error {
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

func (w *WorkflowResultConditionModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowResultConditionModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowResultConditionModelOutput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label *string `json:"label" xml:"label" required:"true"`
	type_ string
	// Whether all tools in the previously executed tool node were executed successfully.
	Successful      bool `json:"successful" xml:"successful" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowResultConditionModelOutput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowResultConditionModelOutput) GetSuccessful() (value bool) {
	if w == nil {
		return
	}
	return w.Successful
}

func (w *WorkflowResultConditionModelOutput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowResultConditionModelOutput) SetSuccessful(successful bool) {
	w.Successful = successful
}

func (w *WorkflowResultConditionModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowResultConditionModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowResultConditionModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowResultConditionModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowResultConditionModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "result" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "result", *literals.Type)
	}
	w.type_ = "result"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowResultConditionModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowResultConditionModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "result",
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowUnconditionalModelOutput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label           *string `json:"label" xml:"label" required:"true"`
	type_           string
	extraProperties map[string]interface{}
}

func (w *WorkflowUnconditionalModelOutput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowUnconditionalModelOutput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowUnconditionalModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowUnconditionalModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowUnconditionalModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowUnconditionalModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowUnconditionalModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "unconditional" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "unconditional", *literals.Type)
	}
	w.type_ = "unconditional"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowUnconditionalModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowUnconditionalModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "unconditional",
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowUnconditionalModelInput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label           *string `json:"label,omitempty" xml:"label,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowUnconditionalModelInput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowUnconditionalModelInput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowUnconditionalModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowUnconditionalModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowUnconditionalModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowUnconditionalModelInput) UnmarshalJSON(data []byte) error {
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

func (w *WorkflowUnconditionalModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowUnconditionalModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

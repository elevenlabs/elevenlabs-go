package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// A common model for workflow tool responses.
type WorkflowToolResponseModelInput struct {
	Steps           []*WorkflowToolResponseModelInputStepsItem `json:"steps,omitempty" xml:"steps,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolResponseModelInput) GetSteps() (value []*WorkflowToolResponseModelInputStepsItem) {
	if w == nil {
		return
	}
	return w.Steps
}

func (w *WorkflowToolResponseModelInput) SetSteps(steps []*WorkflowToolResponseModelInputStepsItem) {
	w.Steps = steps
}

func (w *WorkflowToolResponseModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolResponseModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolResponseModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolResponseModelInput) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkflowToolResponseModelInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkflowToolResponseModelInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolResponseModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolResponseModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

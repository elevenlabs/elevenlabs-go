package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// A common model for workflow tool responses.
type WorkflowToolResponseModelOutput struct {
	Steps           []*WorkflowToolResponseModelOutputStepsItem `json:"steps,omitempty" xml:"steps,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolResponseModelOutput) GetSteps() (value []*WorkflowToolResponseModelOutputStepsItem) {
	if w == nil {
		return
	}
	return w.Steps
}

func (w *WorkflowToolResponseModelOutput) SetSteps(steps []*WorkflowToolResponseModelOutputStepsItem) {
	w.Steps = steps
}

func (w *WorkflowToolResponseModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolResponseModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolResponseModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolResponseModelOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkflowToolResponseModelOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkflowToolResponseModelOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolResponseModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolResponseModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

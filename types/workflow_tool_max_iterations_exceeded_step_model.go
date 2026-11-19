package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowToolMaxIterationsExceededStepModel struct {
	StepLatencySecs float64 `json:"step_latency_secs" xml:"step_latency_secs" required:"true"`
	MaxIterations   int     `json:"max_iterations" xml:"max_iterations" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolMaxIterationsExceededStepModel) GetStepLatencySecs() (value float64) {
	if w == nil {
		return
	}
	return w.StepLatencySecs
}

func (w *WorkflowToolMaxIterationsExceededStepModel) GetMaxIterations() (value int) {
	if w == nil {
		return
	}
	return w.MaxIterations
}

func (w *WorkflowToolMaxIterationsExceededStepModel) SetStepLatencySecs(stepLatencySecs float64) {
	w.StepLatencySecs = stepLatencySecs
}

func (w *WorkflowToolMaxIterationsExceededStepModel) SetMaxIterations(maxIterations int) {
	w.MaxIterations = maxIterations
}

func (w *WorkflowToolMaxIterationsExceededStepModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolMaxIterationsExceededStepModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolMaxIterationsExceededStepModel to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolMaxIterationsExceededStepModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowToolMaxIterationsExceededStepModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowToolMaxIterationsExceededStepModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolMaxIterationsExceededStepModel) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolMaxIterationsExceededStepModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

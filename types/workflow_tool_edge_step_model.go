package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowToolEdgeStepModel struct {
	StepLatencySecs float64 `json:"step_latency_secs" xml:"step_latency_secs" required:"true"`
	EdgeId          string  `json:"edge_id" xml:"edge_id" required:"true"`
	TargetNodeId    string  `json:"target_node_id" xml:"target_node_id" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolEdgeStepModel) GetStepLatencySecs() (value float64) {
	if w == nil {
		return
	}
	return w.StepLatencySecs
}

func (w *WorkflowToolEdgeStepModel) GetEdgeId() (value string) {
	if w == nil {
		return
	}
	return w.EdgeId
}

func (w *WorkflowToolEdgeStepModel) GetTargetNodeId() (value string) {
	if w == nil {
		return
	}
	return w.TargetNodeId
}

func (w *WorkflowToolEdgeStepModel) SetStepLatencySecs(stepLatencySecs float64) {
	w.StepLatencySecs = stepLatencySecs
}

func (w *WorkflowToolEdgeStepModel) SetEdgeId(edgeId string) {
	w.EdgeId = edgeId
}

func (w *WorkflowToolEdgeStepModel) SetTargetNodeId(targetNodeId string) {
	w.TargetNodeId = targetNodeId
}

func (w *WorkflowToolEdgeStepModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolEdgeStepModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolEdgeStepModel to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolEdgeStepModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowToolEdgeStepModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowToolEdgeStepModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolEdgeStepModel) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolEdgeStepModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

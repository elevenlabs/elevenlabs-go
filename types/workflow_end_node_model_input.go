package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowEndNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	// Whether reaching an End node completes successfully or fails.
	Outcome         *WorkflowEndNodeOutcome `json:"outcome,omitempty" xml:"outcome,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowEndNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowEndNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowEndNodeModelInput) GetOutcome() (value *WorkflowEndNodeOutcome) {
	if w == nil {
		return
	}
	return w.Outcome
}

func (w *WorkflowEndNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowEndNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowEndNodeModelInput) SetOutcome(outcome *WorkflowEndNodeOutcome) {
	w.Outcome = outcome
}

func (w *WorkflowEndNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowEndNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowEndNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowEndNodeModelInput) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkflowEndNodeModelInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkflowEndNodeModelInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowEndNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowEndNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowEdgeModelInput struct {
	// ID of the source node.
	Source string `json:"source" xml:"source" required:"true"`
	// ID of the target node.
	Target string `json:"target" xml:"target" required:"true"`
	// Condition that must be met for the edge to be traversed in the forward direction (source to target).
	ForwardCondition *WorkflowEdgeModelInputForwardCondition `json:"forward_condition,omitempty" xml:"forward_condition,omitempty"`
	// Condition that must be met for the edge to be traversed in the backward direction (target to source).
	BackwardCondition *WorkflowEdgeModelInputBackwardCondition `json:"backward_condition,omitempty" xml:"backward_condition,omitempty"`
	// Points the edge is routed through between its source and target, in order. Layout only, no effect on execution.
	Reroutes        []*PositionInput `json:"reroutes,omitempty" xml:"reroutes,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowEdgeModelInput) GetSource() (value string) {
	if w == nil {
		return
	}
	return w.Source
}

func (w *WorkflowEdgeModelInput) GetTarget() (value string) {
	if w == nil {
		return
	}
	return w.Target
}

func (w *WorkflowEdgeModelInput) GetForwardCondition() (value *WorkflowEdgeModelInputForwardCondition) {
	if w == nil {
		return
	}
	return w.ForwardCondition
}

func (w *WorkflowEdgeModelInput) GetBackwardCondition() (value *WorkflowEdgeModelInputBackwardCondition) {
	if w == nil {
		return
	}
	return w.BackwardCondition
}

func (w *WorkflowEdgeModelInput) GetReroutes() (value []*PositionInput) {
	if w == nil {
		return
	}
	return w.Reroutes
}

func (w *WorkflowEdgeModelInput) SetSource(source string) {
	w.Source = source
}

func (w *WorkflowEdgeModelInput) SetTarget(target string) {
	w.Target = target
}

func (w *WorkflowEdgeModelInput) SetForwardCondition(forwardCondition *WorkflowEdgeModelInputForwardCondition) {
	w.ForwardCondition = forwardCondition
}

func (w *WorkflowEdgeModelInput) SetBackwardCondition(backwardCondition *WorkflowEdgeModelInputBackwardCondition) {
	w.BackwardCondition = backwardCondition
}

func (w *WorkflowEdgeModelInput) SetReroutes(reroutes []*PositionInput) {
	w.Reroutes = reroutes
}

func (w *WorkflowEdgeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowEdgeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowEdgeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowEdgeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowEdgeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowEdgeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

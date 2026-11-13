package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowEdgeModelOutput struct {
	// ID of the source node.
	Source string `json:"source" xml:"source" required:"true"`
	// ID of the target node.
	Target string `json:"target" xml:"target" required:"true"`
	// Condition that must be met for the edge to be traversed in the forward direction (source to target).
	ForwardCondition *WorkflowEdgeModelOutputForwardCondition `json:"forward_condition" xml:"forward_condition" required:"true"`
	// Condition that must be met for the edge to be traversed in the backward direction (target to source).
	BackwardCondition *WorkflowEdgeModelOutputBackwardCondition `json:"backward_condition" xml:"backward_condition" required:"true"`
	// Points the edge is routed through between its source and target, in order. Layout only, no effect on execution.
	Reroutes        []*PositionOutput `json:"reroutes" xml:"reroutes" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowEdgeModelOutput) GetSource() (value string) {
	if w == nil {
		return
	}
	return w.Source
}

func (w *WorkflowEdgeModelOutput) GetTarget() (value string) {
	if w == nil {
		return
	}
	return w.Target
}

func (w *WorkflowEdgeModelOutput) GetForwardCondition() (value *WorkflowEdgeModelOutputForwardCondition) {
	if w == nil {
		return
	}
	return w.ForwardCondition
}

func (w *WorkflowEdgeModelOutput) GetBackwardCondition() (value *WorkflowEdgeModelOutputBackwardCondition) {
	if w == nil {
		return
	}
	return w.BackwardCondition
}

func (w *WorkflowEdgeModelOutput) GetReroutes() (value []*PositionOutput) {
	if w == nil {
		return
	}
	return w.Reroutes
}

func (w *WorkflowEdgeModelOutput) SetSource(source string) {
	w.Source = source
}

func (w *WorkflowEdgeModelOutput) SetTarget(target string) {
	w.Target = target
}

func (w *WorkflowEdgeModelOutput) SetForwardCondition(forwardCondition *WorkflowEdgeModelOutputForwardCondition) {
	w.ForwardCondition = forwardCondition
}

func (w *WorkflowEdgeModelOutput) SetBackwardCondition(backwardCondition *WorkflowEdgeModelOutputBackwardCondition) {
	w.BackwardCondition = backwardCondition
}

func (w *WorkflowEdgeModelOutput) SetReroutes(reroutes []*PositionOutput) {
	w.Reroutes = reroutes
}

func (w *WorkflowEdgeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowEdgeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowEdgeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowEdgeModelOutput) UnmarshalJSON(data []byte) error {
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

func (w *WorkflowEdgeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowEdgeModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

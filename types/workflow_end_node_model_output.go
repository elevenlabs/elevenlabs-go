package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowEndNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order" xml:"edge_order" required:"true"`
	// Whether reaching an End node completes successfully or fails.
	Outcome         WorkflowEndNodeOutcome `json:"outcome" xml:"outcome" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowEndNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowEndNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowEndNodeModelOutput) GetOutcome() (value WorkflowEndNodeOutcome) {
	if w == nil {
		return
	}
	return w.Outcome
}

func (w *WorkflowEndNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowEndNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowEndNodeModelOutput) SetOutcome(outcome WorkflowEndNodeOutcome) {
	w.Outcome = outcome
}

func (w *WorkflowEndNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowEndNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowEndNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowEndNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowEndNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowEndNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowEndNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "end" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "end", *literals.Type)
	}
	w.type_ = "end"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowEndNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowEndNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "end",
	}
	return json.Marshal(marshaler)
}

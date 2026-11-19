package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowUpdateStateNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order" xml:"edge_order" required:"true"`
	// Human-readable label for the node used throughout the UI.
	Label string `json:"label" xml:"label" required:"true"`
	// State updates applied when the workflow reaches this node. A node accepts at most 50 updates.
	Updates         []*WorkflowStateUpdateOutput `json:"updates" xml:"updates" required:"true" maxItems:"50"`
	extraProperties map[string]interface{}
}

func (w *WorkflowUpdateStateNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowUpdateStateNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowUpdateStateNodeModelOutput) GetLabel() (value string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowUpdateStateNodeModelOutput) GetUpdates() (value []*WorkflowStateUpdateOutput) {
	if w == nil {
		return
	}
	return w.Updates
}

func (w *WorkflowUpdateStateNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowUpdateStateNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowUpdateStateNodeModelOutput) SetLabel(label string) {
	w.Label = label
}

func (w *WorkflowUpdateStateNodeModelOutput) SetUpdates(updates []*WorkflowStateUpdateOutput) {
	w.Updates = updates
}

func (w *WorkflowUpdateStateNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowUpdateStateNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowUpdateStateNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowUpdateStateNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowUpdateStateNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowUpdateStateNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowUpdateStateNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "update_state" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "update_state", *literals.Type)
	}
	w.type_ = "update_state"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowUpdateStateNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowUpdateStateNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "update_state",
	}
	return json.Marshal(marshaler)
}

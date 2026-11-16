package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowUpdateStateNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	// Human-readable label for the node used throughout the UI.
	Label string `json:"label" xml:"label" required:"true"`
	// State updates applied when the workflow reaches this node. A node accepts at most 50 updates.
	Updates         []*WorkflowStateUpdateInput `json:"updates" xml:"updates" required:"true" maxItems:"50"`
	extraProperties map[string]interface{}
}

func (w *WorkflowUpdateStateNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowUpdateStateNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowUpdateStateNodeModelInput) GetLabel() (value string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowUpdateStateNodeModelInput) GetUpdates() (value []*WorkflowStateUpdateInput) {
	if w == nil {
		return
	}
	return w.Updates
}

func (w *WorkflowUpdateStateNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowUpdateStateNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowUpdateStateNodeModelInput) SetLabel(label string) {
	w.Label = label
}

func (w *WorkflowUpdateStateNodeModelInput) SetUpdates(updates []*WorkflowStateUpdateInput) {
	w.Updates = updates
}

func (w *WorkflowUpdateStateNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowUpdateStateNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowUpdateStateNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowUpdateStateNodeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowUpdateStateNodeModelInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowUpdateStateNodeModelInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowUpdateStateNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowUpdateStateNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

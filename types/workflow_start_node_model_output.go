package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowStartNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder       []string `json:"edge_order" xml:"edge_order" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowStartNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowStartNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowStartNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowStartNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowStartNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowStartNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStartNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStartNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStartNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowStartNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowStartNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "start" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "start", *literals.Type)
	}
	w.type_ = "start"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStartNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStartNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "start",
	}
	return json.Marshal(marshaler)
}

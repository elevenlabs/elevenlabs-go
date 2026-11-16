package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowToolNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order" xml:"edge_order" required:"true"`
	// List of tools to execute in parallel. The entire node is considered successful if all tools are executed successfully.
	Tools           []*WorkflowToolLocator `json:"tools" xml:"tools" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowToolNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowToolNodeModelOutput) GetTools() (value []*WorkflowToolLocator) {
	if w == nil {
		return
	}
	return w.Tools
}

func (w *WorkflowToolNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowToolNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowToolNodeModelOutput) SetTools(tools []*WorkflowToolLocator) {
	w.Tools = tools
}

func (w *WorkflowToolNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowToolNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowToolNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowToolNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "tool" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "tool", *literals.Type)
	}
	w.type_ = "tool"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "tool",
	}
	return json.Marshal(marshaler)
}

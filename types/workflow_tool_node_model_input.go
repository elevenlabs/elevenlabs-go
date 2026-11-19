package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowToolNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	// List of tools to execute in parallel. The entire node is considered successful if all tools are executed successfully.
	Tools           []*WorkflowToolLocator `json:"tools,omitempty" xml:"tools,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowToolNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowToolNodeModelInput) GetTools() (value []*WorkflowToolLocator) {
	if w == nil {
		return
	}
	return w.Tools
}

func (w *WorkflowToolNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowToolNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowToolNodeModelInput) SetTools(tools []*WorkflowToolLocator) {
	w.Tools = tools
}

func (w *WorkflowToolNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolNodeModelInput) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkflowToolNodeModelInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkflowToolNodeModelInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

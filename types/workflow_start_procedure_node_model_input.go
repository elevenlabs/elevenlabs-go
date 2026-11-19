package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Starts one procedure when the workflow reaches it.
//
// A freeform procedure is injected and traversal continues along the node's
// outgoing edges in the same turn. A structured procedure transfers into its
// compiled island with this node as the return node; when it ends, the
// workflow runner hosts this node and evaluates its outgoing edges.
type WorkflowStartProcedureNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	// ID of the procedure to start when the workflow reaches this node.
	ProcedureId     string `json:"procedure_id" xml:"procedure_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (w *WorkflowStartProcedureNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowStartProcedureNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowStartProcedureNodeModelInput) GetProcedureId() (value string) {
	if w == nil {
		return
	}
	return w.ProcedureId
}

func (w *WorkflowStartProcedureNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowStartProcedureNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowStartProcedureNodeModelInput) SetProcedureId(procedureId string) {
	w.ProcedureId = procedureId
}

func (w *WorkflowStartProcedureNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStartProcedureNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStartProcedureNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStartProcedureNodeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowStartProcedureNodeModelInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowStartProcedureNodeModelInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStartProcedureNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStartProcedureNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

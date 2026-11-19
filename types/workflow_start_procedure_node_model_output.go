package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Starts one procedure when the workflow reaches it.
//
// A freeform procedure is injected and traversal continues along the node's
// outgoing edges in the same turn. A structured procedure transfers into its
// compiled island with this node as the return node; when it ends, the
// workflow runner hosts this node and evaluates its outgoing edges.
type WorkflowStartProcedureNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order" xml:"edge_order" required:"true"`
	// ID of the procedure to start when the workflow reaches this node.
	ProcedureId     string `json:"procedure_id" xml:"procedure_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (w *WorkflowStartProcedureNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowStartProcedureNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowStartProcedureNodeModelOutput) GetProcedureId() (value string) {
	if w == nil {
		return
	}
	return w.ProcedureId
}

func (w *WorkflowStartProcedureNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowStartProcedureNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowStartProcedureNodeModelOutput) SetProcedureId(procedureId string) {
	w.ProcedureId = procedureId
}

func (w *WorkflowStartProcedureNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowStartProcedureNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStartProcedureNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStartProcedureNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStartProcedureNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowStartProcedureNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowStartProcedureNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "start_procedure" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "start_procedure", *literals.Type)
	}
	w.type_ = "start_procedure"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStartProcedureNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStartProcedureNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "start_procedure",
	}
	return json.Marshal(marshaler)
}

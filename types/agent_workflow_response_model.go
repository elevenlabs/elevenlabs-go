package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentWorkflowResponseModel struct {
	Edges map[string]*WorkflowEdgeModelOutput              `json:"edges" required:"true"`
	Nodes map[string]*AgentWorkflowResponseModelNodesValue `json:"nodes" required:"true"`
	// Whether to prevent loops in the workflow execution.
	PreventSubagentLoops bool `json:"prevent_subagent_loops" xml:"prevent_subagent_loops" required:"true"`
	extraProperties      map[string]interface{}
}

func (a *AgentWorkflowResponseModel) GetEdges() (value map[string]*WorkflowEdgeModelOutput) {
	if a == nil {
		return
	}
	return a.Edges
}

func (a *AgentWorkflowResponseModel) GetNodes() (value map[string]*AgentWorkflowResponseModelNodesValue) {
	if a == nil {
		return
	}
	return a.Nodes
}

func (a *AgentWorkflowResponseModel) GetPreventSubagentLoops() (value bool) {
	if a == nil {
		return
	}
	return a.PreventSubagentLoops
}

func (a *AgentWorkflowResponseModel) SetEdges(edges map[string]*WorkflowEdgeModelOutput) {
	a.Edges = edges
}

func (a *AgentWorkflowResponseModel) SetNodes(nodes map[string]*AgentWorkflowResponseModelNodesValue) {
	a.Nodes = nodes
}

func (a *AgentWorkflowResponseModel) SetPreventSubagentLoops(preventSubagentLoops bool) {
	a.PreventSubagentLoops = preventSubagentLoops
}

func (a *AgentWorkflowResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentWorkflowResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentWorkflowResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentWorkflowResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentWorkflowResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentWorkflowResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentWorkflowResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentWorkflowResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

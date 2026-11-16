package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentWorkflowRequestModel struct {
	Edges map[string]*WorkflowEdgeModelInput              `json:"edges,omitempty"`
	Nodes map[string]*AgentWorkflowRequestModelNodesValue `json:"nodes,omitempty"`
	// Whether to prevent loops in the workflow execution.
	PreventSubagentLoops *bool `json:"prevent_subagent_loops,omitempty" xml:"prevent_subagent_loops,omitempty"`
	extraProperties      map[string]interface{}
}

func (a *AgentWorkflowRequestModel) GetEdges() (value map[string]*WorkflowEdgeModelInput) {
	if a == nil {
		return
	}
	return a.Edges
}

func (a *AgentWorkflowRequestModel) GetNodes() (value map[string]*AgentWorkflowRequestModelNodesValue) {
	if a == nil {
		return
	}
	return a.Nodes
}

func (a *AgentWorkflowRequestModel) GetPreventSubagentLoops() (value *bool) {
	if a == nil {
		return
	}
	return a.PreventSubagentLoops
}

func (a *AgentWorkflowRequestModel) SetEdges(edges map[string]*WorkflowEdgeModelInput) {
	a.Edges = edges
}

func (a *AgentWorkflowRequestModel) SetNodes(nodes map[string]*AgentWorkflowRequestModelNodesValue) {
	a.Nodes = nodes
}

func (a *AgentWorkflowRequestModel) SetPreventSubagentLoops(preventSubagentLoops *bool) {
	a.PreventSubagentLoops = preventSubagentLoops
}

func (a *AgentWorkflowRequestModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentWorkflowRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentWorkflowRequestModel to string"
	}
	return string(jsonData)
}

func (a *AgentWorkflowRequestModel) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentWorkflowRequestModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentWorkflowRequestModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentWorkflowRequestModel) MarshalJSON() ([]byte, error) {
	type embed AgentWorkflowRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

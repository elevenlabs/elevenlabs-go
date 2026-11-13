package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentMetadata struct {
	AgentId         string  `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId        *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	WorkflowNodeId  *string `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	VersionId       *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentMetadata) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentMetadata) GetBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AgentMetadata) GetWorkflowNodeId() (value *string) {
	if a == nil {
		return
	}
	return a.WorkflowNodeId
}

func (a *AgentMetadata) GetVersionId() (value *string) {
	if a == nil {
		return
	}
	return a.VersionId
}

func (a *AgentMetadata) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentMetadata) SetBranchId(branchId *string) {
	a.BranchId = branchId
}

func (a *AgentMetadata) SetWorkflowNodeId(workflowNodeId *string) {
	a.WorkflowNodeId = workflowNodeId
}

func (a *AgentMetadata) SetVersionId(versionId *string) {
	a.VersionId = versionId
}

func (a *AgentMetadata) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentMetadata) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentMetadata to string"
	}
	return string(jsonData)
}

func (a *AgentMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentMetadata) MarshalJSON() ([]byte, error) {
	type embed AgentMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

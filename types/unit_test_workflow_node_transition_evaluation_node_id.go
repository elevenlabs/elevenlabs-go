package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UnitTestWorkflowNodeTransitionEvaluationNodeId struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the agent whose workflow contains the target node.
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The ID of the workflow node that the agent should transition to.
	TargetNodeId    string `json:"target_node_id" xml:"target_node_id" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) GetAgentId() (value string) {
	if u == nil {
		return
	}
	return u.AgentId
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) GetTargetNodeId() (value string) {
	if u == nil {
		return
	}
	return u.TargetNodeId
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) SetType(type_ *string) {
	u.Type = type_
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) SetAgentId(agentId string) {
	u.AgentId = agentId
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) SetTargetNodeId(targetNodeId string) {
	u.TargetNodeId = targetNodeId
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UnitTestWorkflowNodeTransitionEvaluationNodeId) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnitTestWorkflowNodeTransitionEvaluationNodeId to string"
	}
	return string(jsonData)
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UnitTestWorkflowNodeTransitionEvaluationNodeId
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UnitTestWorkflowNodeTransitionEvaluationNodeId(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UnitTestWorkflowNodeTransitionEvaluationNodeId) MarshalJSON() ([]byte, error) {
	type embed UnitTestWorkflowNodeTransitionEvaluationNodeId
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

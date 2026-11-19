package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An agent (and optional branch) that participated in the call, in first-seen transcript order.
type VisitedAgentRef struct {
	AgentId         string  `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId        *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VisitedAgentRef) GetAgentId() (value string) {
	if v == nil {
		return
	}
	return v.AgentId
}

func (v *VisitedAgentRef) GetBranchId() (value *string) {
	if v == nil {
		return
	}
	return v.BranchId
}

func (v *VisitedAgentRef) SetAgentId(agentId string) {
	v.AgentId = agentId
}

func (v *VisitedAgentRef) SetBranchId(branchId *string) {
	v.BranchId = branchId
}

func (v *VisitedAgentRef) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VisitedAgentRef) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VisitedAgentRef to string"
	}
	return string(jsonData)
}

func (v *VisitedAgentRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VisitedAgentRef) MarshalJSON() ([]byte, error) {
	type embed VisitedAgentRef
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}

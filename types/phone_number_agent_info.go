package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneNumberAgentInfo struct {
	// The ID of the agent
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The name of the agent
	AgentName string `json:"agent_name" xml:"agent_name" required:"true"`
	// Environment to use for resolving environment variables on calls to this number.
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// Agent branch to use for calls to this number.
	BranchId        *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PhoneNumberAgentInfo) GetAgentId() (value string) {
	if p == nil {
		return
	}
	return p.AgentId
}

func (p *PhoneNumberAgentInfo) GetAgentName() (value string) {
	if p == nil {
		return
	}
	return p.AgentName
}

func (p *PhoneNumberAgentInfo) GetEnvironment() (value *string) {
	if p == nil {
		return
	}
	return p.Environment
}

func (p *PhoneNumberAgentInfo) GetBranchId() (value *string) {
	if p == nil {
		return
	}
	return p.BranchId
}

func (p *PhoneNumberAgentInfo) SetAgentId(agentId string) {
	p.AgentId = agentId
}

func (p *PhoneNumberAgentInfo) SetAgentName(agentName string) {
	p.AgentName = agentName
}

func (p *PhoneNumberAgentInfo) SetEnvironment(environment *string) {
	p.Environment = environment
}

func (p *PhoneNumberAgentInfo) SetBranchId(branchId *string) {
	p.BranchId = branchId
}

func (p *PhoneNumberAgentInfo) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneNumberAgentInfo) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneNumberAgentInfo to string"
	}
	return string(jsonData)
}

func (p *PhoneNumberAgentInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneNumberAgentInfo) MarshalJSON() ([]byte, error) {
	type embed PhoneNumberAgentInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}

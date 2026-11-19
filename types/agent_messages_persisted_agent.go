package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentMessagesPersistedAgent struct {
	type_   string
	AgentId string `json:"agent_id" xml:"agent_id" required:"true" minLength:"1"`
	// ID of the agent branch to use for this conversation
	BranchId        *string                 `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	Overrides       *AgentMessagesOverrides `json:"overrides,omitempty" xml:"overrides,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentMessagesPersistedAgent) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentMessagesPersistedAgent) GetBranchId() (value *string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AgentMessagesPersistedAgent) GetOverrides() (value *AgentMessagesOverrides) {
	if a == nil {
		return
	}
	return a.Overrides
}

func (a *AgentMessagesPersistedAgent) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentMessagesPersistedAgent) SetBranchId(branchId *string) {
	a.BranchId = branchId
}

func (a *AgentMessagesPersistedAgent) SetOverrides(overrides *AgentMessagesOverrides) {
	a.Overrides = overrides
}

func (a *AgentMessagesPersistedAgent) Type() string {
	return a.type_
}

func (a *AgentMessagesPersistedAgent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentMessagesPersistedAgent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentMessagesPersistedAgent to string"
	}
	return string(jsonData)
}

func (a *AgentMessagesPersistedAgent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "persisted" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "persisted", *literals.Type)
	}
	a.type_ = "persisted"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentMessagesPersistedAgent) MarshalJSON() ([]byte, error) {
	type embed AgentMessagesPersistedAgent
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "persisted",
	}
	return json.Marshal(marshaler)
}

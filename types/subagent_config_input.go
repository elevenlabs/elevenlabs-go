package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubagentConfigInput struct {
	// Agents this agent can hand a single turn to.
	Agents []*SubAgentInput `json:"agents,omitempty" xml:"agents,omitempty" maxItems:"10"`
	// Lets this agent hand a single turn to a fresh copy of itself.
	SelfSubagent    *SelfSubagent `json:"self_subagent,omitempty" xml:"self_subagent,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SubagentConfigInput) GetAgents() (value []*SubAgentInput) {
	if s == nil {
		return
	}
	return s.Agents
}

func (s *SubagentConfigInput) GetSelfSubagent() (value *SelfSubagent) {
	if s == nil {
		return
	}
	return s.SelfSubagent
}

func (s *SubagentConfigInput) SetAgents(agents []*SubAgentInput) {
	s.Agents = agents
}

func (s *SubagentConfigInput) SetSelfSubagent(selfSubagent *SelfSubagent) {
	s.SelfSubagent = selfSubagent
}

func (s *SubagentConfigInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubagentConfigInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubagentConfigInput to string"
	}
	return string(jsonData)
}

func (s *SubagentConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubagentConfigInput) MarshalJSON() ([]byte, error) {
	type embed SubagentConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

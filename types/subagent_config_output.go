package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubagentConfigOutput struct {
	// Agents this agent can hand a single turn to.
	Agents []*SubAgentOutput `json:"agents,omitempty" xml:"agents,omitempty" maxItems:"10"`
	// Lets this agent hand a single turn to a fresh copy of itself.
	SelfSubagent    *SelfSubagent `json:"self_subagent,omitempty" xml:"self_subagent,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SubagentConfigOutput) GetAgents() (value []*SubAgentOutput) {
	if s == nil {
		return
	}
	return s.Agents
}

func (s *SubagentConfigOutput) GetSelfSubagent() (value *SelfSubagent) {
	if s == nil {
		return
	}
	return s.SelfSubagent
}

func (s *SubagentConfigOutput) SetAgents(agents []*SubAgentOutput) {
	s.Agents = agents
}

func (s *SubagentConfigOutput) SetSelfSubagent(selfSubagent *SelfSubagent) {
	s.SelfSubagent = selfSubagent
}

func (s *SubagentConfigOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubagentConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubagentConfigOutput to string"
	}
	return string(jsonData)
}

func (s *SubagentConfigOutput) UnmarshalJSON(data []byte) error {
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

func (s *SubagentConfigOutput) MarshalJSON() ([]byte, error) {
	type embed SubagentConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

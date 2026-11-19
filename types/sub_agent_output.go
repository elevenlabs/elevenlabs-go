package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubAgentOutput struct {
	// Whether to wait for the reply, or keep going and get it later in the conversation.
	ExecutionMode *SubAgentOutputExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// How long the agent may run before it's stopped.
	TimeoutSecs *int `json:"timeout_secs,omitempty" xml:"timeout_secs,omitempty" min:"10" max:"3600"`
	// The agent to run. Must be in the same workspace.
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The branch to run. Defaults to the agent's live deployment.
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// What the agent is for. Shown to the LLM when it picks an agent.
	Description string `json:"description" xml:"description" required:"true"`
	// Values passed to the agent as dynamic variables.
	Parameters      *ObjectJsonSchemaPropertyOutput `json:"parameters,omitempty" xml:"parameters,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SubAgentOutput) GetExecutionMode() (value *SubAgentOutputExecutionMode) {
	if s == nil {
		return
	}
	return s.ExecutionMode
}

func (s *SubAgentOutput) GetTimeoutSecs() (value *int) {
	if s == nil {
		return
	}
	return s.TimeoutSecs
}

func (s *SubAgentOutput) GetAgentId() (value string) {
	if s == nil {
		return
	}
	return s.AgentId
}

func (s *SubAgentOutput) GetBranchId() (value *string) {
	if s == nil {
		return
	}
	return s.BranchId
}

func (s *SubAgentOutput) GetDescription() (value string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SubAgentOutput) GetParameters() (value *ObjectJsonSchemaPropertyOutput) {
	if s == nil {
		return
	}
	return s.Parameters
}

func (s *SubAgentOutput) SetExecutionMode(executionMode *SubAgentOutputExecutionMode) {
	s.ExecutionMode = executionMode
}

func (s *SubAgentOutput) SetTimeoutSecs(timeoutSecs *int) {
	s.TimeoutSecs = timeoutSecs
}

func (s *SubAgentOutput) SetAgentId(agentId string) {
	s.AgentId = agentId
}

func (s *SubAgentOutput) SetBranchId(branchId *string) {
	s.BranchId = branchId
}

func (s *SubAgentOutput) SetDescription(description string) {
	s.Description = description
}

func (s *SubAgentOutput) SetParameters(parameters *ObjectJsonSchemaPropertyOutput) {
	s.Parameters = parameters
}

func (s *SubAgentOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubAgentOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubAgentOutput to string"
	}
	return string(jsonData)
}

func (s *SubAgentOutput) UnmarshalJSON(data []byte) error {
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

func (s *SubAgentOutput) MarshalJSON() ([]byte, error) {
	type embed SubAgentOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

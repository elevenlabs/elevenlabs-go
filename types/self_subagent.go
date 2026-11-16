package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The agent itself, run on the same branch with none of the conversation.
type SelfSubagent struct {
	// Whether to wait for the reply, or keep going and get it later in the conversation.
	ExecutionMode *SubAgentInputExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// How long the agent may run before it's stopped.
	TimeoutSecs *int `json:"timeout_secs,omitempty" xml:"timeout_secs,omitempty" min:"10" max:"3600"`
	// When the agent should hand work to itself. Defaults to a built-in description.
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SelfSubagent) GetExecutionMode() (value *SubAgentInputExecutionMode) {
	if s == nil {
		return
	}
	return s.ExecutionMode
}

func (s *SelfSubagent) GetTimeoutSecs() (value *int) {
	if s == nil {
		return
	}
	return s.TimeoutSecs
}

func (s *SelfSubagent) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SelfSubagent) SetExecutionMode(executionMode *SubAgentInputExecutionMode) {
	s.ExecutionMode = executionMode
}

func (s *SelfSubagent) SetTimeoutSecs(timeoutSecs *int) {
	s.TimeoutSecs = timeoutSecs
}

func (s *SelfSubagent) SetDescription(description *string) {
	s.Description = description
}

func (s *SelfSubagent) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SelfSubagent) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SelfSubagent to string"
	}
	return string(jsonData)
}

func (s *SelfSubagent) UnmarshalJSON(data []byte) error {
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

func (s *SelfSubagent) MarshalJSON() ([]byte, error) {
	type embed SelfSubagent
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

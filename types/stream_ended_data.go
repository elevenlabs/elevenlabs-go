package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StreamEndedData struct {
	Status AgentOutputStatus `json:"status" xml:"status" required:"true"`
	Error  *AgentOutputError `json:"error,omitempty" xml:"error,omitempty"`
	// The actions the agent waits on. Answer each `client_tool_call` with a `client_tool_result` item in `input` to continue.
	RequiredActions []*ClientToolCallRequiredAction `json:"required_actions,omitempty" xml:"required_actions,omitempty"`
	extraProperties map[string]interface{}
}

func (s *StreamEndedData) GetStatus() (value AgentOutputStatus) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *StreamEndedData) GetError() (value *AgentOutputError) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *StreamEndedData) GetRequiredActions() (value []*ClientToolCallRequiredAction) {
	if s == nil {
		return
	}
	return s.RequiredActions
}

func (s *StreamEndedData) SetStatus(status AgentOutputStatus) {
	s.Status = status
}

func (s *StreamEndedData) SetError(error *AgentOutputError) {
	s.Error = error
}

func (s *StreamEndedData) SetRequiredActions(requiredActions []*ClientToolCallRequiredAction) {
	s.RequiredActions = requiredActions
}

func (s *StreamEndedData) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StreamEndedData) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StreamEndedData to string"
	}
	return string(jsonData)
}

func (s *StreamEndedData) UnmarshalJSON(data []byte) error {
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

func (s *StreamEndedData) MarshalJSON() ([]byte, error) {
	type embed StreamEndedData
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

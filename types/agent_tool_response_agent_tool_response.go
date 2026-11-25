package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolResponseAgentToolResponse struct {
	ToolName   string `json:"tool_name" xml:"tool_name" required:"true"`
	ToolCallId string `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolType   string `json:"tool_type" xml:"tool_type" required:"true"`
	IsError    bool   `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked  *bool  `json:"is_blocked,omitempty" xml:"is_blocked,omitempty"`
	EventId    int    `json:"event_id" xml:"event_id" required:"true"`
	IsCalled   bool   `json:"is_called" xml:"is_called" required:"true"`
	// Tool call status derived from execution flags.
	Status          *AgentToolResponseAgentToolResponseStatus `json:"status,omitempty" xml:"status,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentToolResponseAgentToolResponse) GetToolName() (value string) {
	if a == nil {
		return
	}
	return a.ToolName
}

func (a *AgentToolResponseAgentToolResponse) GetToolCallId() (value string) {
	if a == nil {
		return
	}
	return a.ToolCallId
}

func (a *AgentToolResponseAgentToolResponse) GetToolType() (value string) {
	if a == nil {
		return
	}
	return a.ToolType
}

func (a *AgentToolResponseAgentToolResponse) GetIsError() (value bool) {
	if a == nil {
		return
	}
	return a.IsError
}

func (a *AgentToolResponseAgentToolResponse) GetIsBlocked() (value *bool) {
	if a == nil {
		return
	}
	return a.IsBlocked
}

func (a *AgentToolResponseAgentToolResponse) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentToolResponseAgentToolResponse) GetIsCalled() (value bool) {
	if a == nil {
		return
	}
	return a.IsCalled
}

func (a *AgentToolResponseAgentToolResponse) GetStatus() (value *AgentToolResponseAgentToolResponseStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AgentToolResponseAgentToolResponse) SetToolName(toolName string) {
	a.ToolName = toolName
}

func (a *AgentToolResponseAgentToolResponse) SetToolCallId(toolCallId string) {
	a.ToolCallId = toolCallId
}

func (a *AgentToolResponseAgentToolResponse) SetToolType(toolType string) {
	a.ToolType = toolType
}

func (a *AgentToolResponseAgentToolResponse) SetIsError(isError bool) {
	a.IsError = isError
}

func (a *AgentToolResponseAgentToolResponse) SetIsBlocked(isBlocked *bool) {
	a.IsBlocked = isBlocked
}

func (a *AgentToolResponseAgentToolResponse) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentToolResponseAgentToolResponse) SetIsCalled(isCalled bool) {
	a.IsCalled = isCalled
}

func (a *AgentToolResponseAgentToolResponse) SetStatus(status *AgentToolResponseAgentToolResponseStatus) {
	a.Status = status
}

func (a *AgentToolResponseAgentToolResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolResponseAgentToolResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolResponseAgentToolResponse to string"
	}
	return string(jsonData)
}

func (a *AgentToolResponseAgentToolResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolResponseAgentToolResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolResponseAgentToolResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolResponseAgentToolResponse) MarshalJSON() ([]byte, error) {
	type embed AgentToolResponseAgentToolResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

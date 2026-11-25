package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolResponseFullPayloadAgentToolResponseFullPayload struct {
	ToolName       string `json:"tool_name" xml:"tool_name" required:"true"`
	ToolCallId     string `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolType       string `json:"tool_type" xml:"tool_type" required:"true"`
	IsError        bool   `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked      *bool  `json:"is_blocked,omitempty" xml:"is_blocked,omitempty"`
	EventId        int    `json:"event_id" xml:"event_id" required:"true"`
	IsCalled       bool   `json:"is_called" xml:"is_called" required:"true"`
	FullToolResult string `json:"full_tool_result" xml:"full_tool_result" required:"true"`
	Truncated      *bool  `json:"truncated,omitempty" xml:"truncated,omitempty"`
	// Tool call status derived from execution flags.
	Status          *AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus `json:"status,omitempty" xml:"status,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetToolName() (value string) {
	if a == nil {
		return
	}
	return a.ToolName
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetToolCallId() (value string) {
	if a == nil {
		return
	}
	return a.ToolCallId
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetToolType() (value string) {
	if a == nil {
		return
	}
	return a.ToolType
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetIsError() (value bool) {
	if a == nil {
		return
	}
	return a.IsError
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetIsBlocked() (value *bool) {
	if a == nil {
		return
	}
	return a.IsBlocked
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetIsCalled() (value bool) {
	if a == nil {
		return
	}
	return a.IsCalled
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetFullToolResult() (value string) {
	if a == nil {
		return
	}
	return a.FullToolResult
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetTruncated() (value *bool) {
	if a == nil {
		return
	}
	return a.Truncated
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetStatus() (value *AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetToolName(toolName string) {
	a.ToolName = toolName
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetToolCallId(toolCallId string) {
	a.ToolCallId = toolCallId
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetToolType(toolType string) {
	a.ToolType = toolType
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetIsError(isError bool) {
	a.IsError = isError
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetIsBlocked(isBlocked *bool) {
	a.IsBlocked = isBlocked
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetIsCalled(isCalled bool) {
	a.IsCalled = isCalled
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetFullToolResult(fullToolResult string) {
	a.FullToolResult = fullToolResult
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetTruncated(truncated *bool) {
	a.Truncated = truncated
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) SetStatus(status *AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus) {
	a.Status = status
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolResponseFullPayloadAgentToolResponseFullPayload) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolResponseFullPayloadAgentToolResponseFullPayload to string"
	}
	return string(jsonData)
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolResponseFullPayloadAgentToolResponseFullPayload
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolResponseFullPayloadAgentToolResponseFullPayload(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolResponseFullPayloadAgentToolResponseFullPayload) MarshalJSON() ([]byte, error) {
	type embed AgentToolResponseFullPayloadAgentToolResponseFullPayload
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

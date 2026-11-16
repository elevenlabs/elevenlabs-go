package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolResponseClientEventData struct {
	ToolName        string                   `json:"tool_name" xml:"tool_name" required:"true"`
	ToolCallId      string                   `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolType        string                   `json:"tool_type" xml:"tool_type" required:"true"`
	IsError         bool                     `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked       *bool                    `json:"is_blocked,omitempty" xml:"is_blocked,omitempty"`
	EventId         int                      `json:"event_id" xml:"event_id" required:"true"`
	IsCalled        bool                     `json:"is_called" xml:"is_called" required:"true"`
	Status          *AgentToolResponseStatus `json:"status,omitempty" xml:"status,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentToolResponseClientEventData) GetToolName() (value string) {
	if a == nil {
		return
	}
	return a.ToolName
}

func (a *AgentToolResponseClientEventData) GetToolCallId() (value string) {
	if a == nil {
		return
	}
	return a.ToolCallId
}

func (a *AgentToolResponseClientEventData) GetToolType() (value string) {
	if a == nil {
		return
	}
	return a.ToolType
}

func (a *AgentToolResponseClientEventData) GetIsError() (value bool) {
	if a == nil {
		return
	}
	return a.IsError
}

func (a *AgentToolResponseClientEventData) GetIsBlocked() (value *bool) {
	if a == nil {
		return
	}
	return a.IsBlocked
}

func (a *AgentToolResponseClientEventData) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentToolResponseClientEventData) GetIsCalled() (value bool) {
	if a == nil {
		return
	}
	return a.IsCalled
}

func (a *AgentToolResponseClientEventData) GetStatus() (value *AgentToolResponseStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AgentToolResponseClientEventData) SetToolName(toolName string) {
	a.ToolName = toolName
}

func (a *AgentToolResponseClientEventData) SetToolCallId(toolCallId string) {
	a.ToolCallId = toolCallId
}

func (a *AgentToolResponseClientEventData) SetToolType(toolType string) {
	a.ToolType = toolType
}

func (a *AgentToolResponseClientEventData) SetIsError(isError bool) {
	a.IsError = isError
}

func (a *AgentToolResponseClientEventData) SetIsBlocked(isBlocked *bool) {
	a.IsBlocked = isBlocked
}

func (a *AgentToolResponseClientEventData) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentToolResponseClientEventData) SetIsCalled(isCalled bool) {
	a.IsCalled = isCalled
}

func (a *AgentToolResponseClientEventData) SetStatus(status *AgentToolResponseStatus) {
	a.Status = status
}

func (a *AgentToolResponseClientEventData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolResponseClientEventData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolResponseClientEventData to string"
	}
	return string(jsonData)
}

func (a *AgentToolResponseClientEventData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolResponseClientEventData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolResponseClientEventData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolResponseClientEventData) MarshalJSON() ([]byte, error) {
	type embed AgentToolResponseClientEventData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

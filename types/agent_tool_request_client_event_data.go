package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolRequestClientEventData struct {
	ToolName             string `json:"tool_name" xml:"tool_name" required:"true"`
	ToolCallId           string `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolType             string `json:"tool_type" xml:"tool_type" required:"true"`
	EventId              int    `json:"event_id" xml:"event_id" required:"true"`
	ExpectsResponse      bool   `json:"expects_response" xml:"expects_response" required:"true"`
	DisableInterruptions bool   `json:"disable_interruptions" xml:"disable_interruptions" required:"true"`
	ResponseTimeoutSecs  int    `json:"response_timeout_secs" xml:"response_timeout_secs" required:"true"`
	ExecutionMode        string `json:"execution_mode" xml:"execution_mode" required:"true"`
	extraProperties      map[string]interface{}
}

func (a *AgentToolRequestClientEventData) GetToolName() (value string) {
	if a == nil {
		return
	}
	return a.ToolName
}

func (a *AgentToolRequestClientEventData) GetToolCallId() (value string) {
	if a == nil {
		return
	}
	return a.ToolCallId
}

func (a *AgentToolRequestClientEventData) GetToolType() (value string) {
	if a == nil {
		return
	}
	return a.ToolType
}

func (a *AgentToolRequestClientEventData) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentToolRequestClientEventData) GetExpectsResponse() (value bool) {
	if a == nil {
		return
	}
	return a.ExpectsResponse
}

func (a *AgentToolRequestClientEventData) GetDisableInterruptions() (value bool) {
	if a == nil {
		return
	}
	return a.DisableInterruptions
}

func (a *AgentToolRequestClientEventData) GetResponseTimeoutSecs() (value int) {
	if a == nil {
		return
	}
	return a.ResponseTimeoutSecs
}

func (a *AgentToolRequestClientEventData) GetExecutionMode() (value string) {
	if a == nil {
		return
	}
	return a.ExecutionMode
}

func (a *AgentToolRequestClientEventData) SetToolName(toolName string) {
	a.ToolName = toolName
}

func (a *AgentToolRequestClientEventData) SetToolCallId(toolCallId string) {
	a.ToolCallId = toolCallId
}

func (a *AgentToolRequestClientEventData) SetToolType(toolType string) {
	a.ToolType = toolType
}

func (a *AgentToolRequestClientEventData) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentToolRequestClientEventData) SetExpectsResponse(expectsResponse bool) {
	a.ExpectsResponse = expectsResponse
}

func (a *AgentToolRequestClientEventData) SetDisableInterruptions(disableInterruptions bool) {
	a.DisableInterruptions = disableInterruptions
}

func (a *AgentToolRequestClientEventData) SetResponseTimeoutSecs(responseTimeoutSecs int) {
	a.ResponseTimeoutSecs = responseTimeoutSecs
}

func (a *AgentToolRequestClientEventData) SetExecutionMode(executionMode string) {
	a.ExecutionMode = executionMode
}

func (a *AgentToolRequestClientEventData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolRequestClientEventData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolRequestClientEventData to string"
	}
	return string(jsonData)
}

func (a *AgentToolRequestClientEventData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolRequestClientEventData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolRequestClientEventData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolRequestClientEventData) MarshalJSON() ([]byte, error) {
	type embed AgentToolRequestClientEventData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

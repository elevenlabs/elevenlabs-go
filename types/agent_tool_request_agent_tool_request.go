package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolRequestAgentToolRequest struct {
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

func (a *AgentToolRequestAgentToolRequest) GetToolName() (value string) {
	if a == nil {
		return
	}
	return a.ToolName
}

func (a *AgentToolRequestAgentToolRequest) GetToolCallId() (value string) {
	if a == nil {
		return
	}
	return a.ToolCallId
}

func (a *AgentToolRequestAgentToolRequest) GetToolType() (value string) {
	if a == nil {
		return
	}
	return a.ToolType
}

func (a *AgentToolRequestAgentToolRequest) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentToolRequestAgentToolRequest) GetExpectsResponse() (value bool) {
	if a == nil {
		return
	}
	return a.ExpectsResponse
}

func (a *AgentToolRequestAgentToolRequest) GetDisableInterruptions() (value bool) {
	if a == nil {
		return
	}
	return a.DisableInterruptions
}

func (a *AgentToolRequestAgentToolRequest) GetResponseTimeoutSecs() (value int) {
	if a == nil {
		return
	}
	return a.ResponseTimeoutSecs
}

func (a *AgentToolRequestAgentToolRequest) GetExecutionMode() (value string) {
	if a == nil {
		return
	}
	return a.ExecutionMode
}

func (a *AgentToolRequestAgentToolRequest) SetToolName(toolName string) {
	a.ToolName = toolName
}

func (a *AgentToolRequestAgentToolRequest) SetToolCallId(toolCallId string) {
	a.ToolCallId = toolCallId
}

func (a *AgentToolRequestAgentToolRequest) SetToolType(toolType string) {
	a.ToolType = toolType
}

func (a *AgentToolRequestAgentToolRequest) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentToolRequestAgentToolRequest) SetExpectsResponse(expectsResponse bool) {
	a.ExpectsResponse = expectsResponse
}

func (a *AgentToolRequestAgentToolRequest) SetDisableInterruptions(disableInterruptions bool) {
	a.DisableInterruptions = disableInterruptions
}

func (a *AgentToolRequestAgentToolRequest) SetResponseTimeoutSecs(responseTimeoutSecs int) {
	a.ResponseTimeoutSecs = responseTimeoutSecs
}

func (a *AgentToolRequestAgentToolRequest) SetExecutionMode(executionMode string) {
	a.ExecutionMode = executionMode
}

func (a *AgentToolRequestAgentToolRequest) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolRequestAgentToolRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolRequestAgentToolRequest to string"
	}
	return string(jsonData)
}

func (a *AgentToolRequestAgentToolRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolRequestAgentToolRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolRequestAgentToolRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolRequestAgentToolRequest) MarshalJSON() ([]byte, error) {
	type embed AgentToolRequestAgentToolRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

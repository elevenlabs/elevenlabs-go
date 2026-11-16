package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptOtherToolsResultCommonModel struct {
	RequestId string `json:"request_id" xml:"request_id" required:"true"`
	ToolName  string `json:"tool_name" xml:"tool_name" required:"true"`
	// ID of the tool that produced this result: a workspace tool ID, 'mcp:<server_id>:<tool_name>' for MCP tools, or '__system_<tool_name>' for system tools. Not set for conversations recorded before tool IDs were tracked.
	ToolId                 *string                                                       `json:"tool_id,omitempty" xml:"tool_id,omitempty"`
	ResultValue            string                                                        `json:"result_value" xml:"result_value" required:"true"`
	IsError                bool                                                          `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked              *bool                                                         `json:"is_blocked,omitempty" xml:"is_blocked,omitempty"`
	ToolHasBeenCalled      bool                                                          `json:"tool_has_been_called" xml:"tool_has_been_called" required:"true"`
	ToolLatencySecs        *float64                                                      `json:"tool_latency_secs,omitempty" xml:"tool_latency_secs,omitempty"`
	ErrorType              *string                                                       `json:"error_type,omitempty" xml:"error_type,omitempty"`
	RawErrorMessage        *string                                                       `json:"raw_error_message,omitempty" xml:"raw_error_message,omitempty"`
	DynamicVariableUpdates []*DynamicVariableUpdateCommonModel                           `json:"dynamic_variable_updates,omitempty" xml:"dynamic_variable_updates,omitempty"`
	Type                   *ConversationHistoryTranscriptOtherToolsResultCommonModelType `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties        map[string]interface{}
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetResultValue() (value string) {
	if c == nil {
		return
	}
	return c.ResultValue
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetIsBlocked() (value *bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetToolLatencySecs() (value *float64) {
	if c == nil {
		return
	}
	return c.ToolLatencySecs
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetErrorType() (value *string) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetRawErrorMessage() (value *string) {
	if c == nil {
		return
	}
	return c.RawErrorMessage
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetDynamicVariableUpdates() (value []*DynamicVariableUpdateCommonModel) {
	if c == nil {
		return
	}
	return c.DynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetType() (value *ConversationHistoryTranscriptOtherToolsResultCommonModelType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetResultValue(resultValue string) {
	c.ResultValue = resultValue
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetIsBlocked(isBlocked *bool) {
	c.IsBlocked = isBlocked
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetToolLatencySecs(toolLatencySecs *float64) {
	c.ToolLatencySecs = toolLatencySecs
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetErrorType(errorType *string) {
	c.ErrorType = errorType
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetRawErrorMessage(rawErrorMessage *string) {
	c.RawErrorMessage = rawErrorMessage
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetDynamicVariableUpdates(dynamicVariableUpdates []*DynamicVariableUpdateCommonModel) {
	c.DynamicVariableUpdates = dynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) SetType(type_ *ConversationHistoryTranscriptOtherToolsResultCommonModelType) {
	c.Type = type_
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptOtherToolsResultCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptOtherToolsResultCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptOtherToolsResultCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptOtherToolsResultCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

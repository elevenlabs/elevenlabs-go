package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptSystemToolResultCommonModelOutput struct {
	RequestId string `json:"request_id" xml:"request_id" required:"true"`
	ToolName  string `json:"tool_name" xml:"tool_name" required:"true"`
	// ID of the tool that produced this result: a workspace tool ID, 'mcp:<server_id>:<tool_name>' for MCP tools, or '__system_<tool_name>' for system tools. Not set for conversations recorded before tool IDs were tracked.
	ToolId                 *string                             `json:"tool_id,omitempty" xml:"tool_id,omitempty"`
	ResultValue            string                              `json:"result_value" xml:"result_value" required:"true"`
	IsError                bool                                `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked              *bool                               `json:"is_blocked,omitempty" xml:"is_blocked,omitempty"`
	ToolHasBeenCalled      bool                                `json:"tool_has_been_called" xml:"tool_has_been_called" required:"true"`
	ToolLatencySecs        *float64                            `json:"tool_latency_secs,omitempty" xml:"tool_latency_secs,omitempty"`
	ErrorType              *string                             `json:"error_type,omitempty" xml:"error_type,omitempty"`
	RawErrorMessage        *string                             `json:"raw_error_message,omitempty" xml:"raw_error_message,omitempty"`
	DynamicVariableUpdates []*DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates,omitempty" xml:"dynamic_variable_updates,omitempty"`
	type_                  string
	Result                 *ConversationHistoryTranscriptSystemToolResultCommonModelOutputResult `json:"result,omitempty" xml:"result,omitempty"`
	extraProperties        map[string]interface{}
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetResultValue() (value string) {
	if c == nil {
		return
	}
	return c.ResultValue
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetIsBlocked() (value *bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetToolLatencySecs() (value *float64) {
	if c == nil {
		return
	}
	return c.ToolLatencySecs
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetErrorType() (value *string) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetRawErrorMessage() (value *string) {
	if c == nil {
		return
	}
	return c.RawErrorMessage
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetDynamicVariableUpdates() (value []*DynamicVariableUpdateCommonModel) {
	if c == nil {
		return
	}
	return c.DynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetResult() (value *ConversationHistoryTranscriptSystemToolResultCommonModelOutputResult) {
	if c == nil {
		return
	}
	return c.Result
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetResultValue(resultValue string) {
	c.ResultValue = resultValue
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetIsBlocked(isBlocked *bool) {
	c.IsBlocked = isBlocked
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetToolLatencySecs(toolLatencySecs *float64) {
	c.ToolLatencySecs = toolLatencySecs
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetErrorType(errorType *string) {
	c.ErrorType = errorType
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetRawErrorMessage(rawErrorMessage *string) {
	c.RawErrorMessage = rawErrorMessage
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetDynamicVariableUpdates(dynamicVariableUpdates []*DynamicVariableUpdateCommonModel) {
	c.DynamicVariableUpdates = dynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) SetResult(result *ConversationHistoryTranscriptSystemToolResultCommonModelOutputResult) {
	c.Result = result
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) Type() string {
	return c.type_
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptSystemToolResultCommonModelOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptSystemToolResultCommonModelOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "system" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "system", *literals.Type)
	}
	c.type_ = "system"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptSystemToolResultCommonModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "system",
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptSystemToolResultCommonModelInput struct {
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
	Result                 *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult `json:"result,omitempty" xml:"result,omitempty"`
	extraProperties        map[string]interface{}
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetResultValue() (value string) {
	if c == nil {
		return
	}
	return c.ResultValue
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetIsBlocked() (value *bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetToolLatencySecs() (value *float64) {
	if c == nil {
		return
	}
	return c.ToolLatencySecs
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetErrorType() (value *string) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetRawErrorMessage() (value *string) {
	if c == nil {
		return
	}
	return c.RawErrorMessage
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetDynamicVariableUpdates() (value []*DynamicVariableUpdateCommonModel) {
	if c == nil {
		return
	}
	return c.DynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetResult() (value *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) {
	if c == nil {
		return
	}
	return c.Result
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetResultValue(resultValue string) {
	c.ResultValue = resultValue
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetIsBlocked(isBlocked *bool) {
	c.IsBlocked = isBlocked
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetToolLatencySecs(toolLatencySecs *float64) {
	c.ToolLatencySecs = toolLatencySecs
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetErrorType(errorType *string) {
	c.ErrorType = errorType
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetRawErrorMessage(rawErrorMessage *string) {
	c.RawErrorMessage = rawErrorMessage
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetDynamicVariableUpdates(dynamicVariableUpdates []*DynamicVariableUpdateCommonModel) {
	c.DynamicVariableUpdates = dynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) SetResult(result *ConversationHistoryTranscriptSystemToolResultCommonModelInputResult) {
	c.Result = result
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) Type() string {
	return c.type_
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptSystemToolResultCommonModelInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptSystemToolResultCommonModelInput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryTranscriptSystemToolResultCommonModelInput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptSystemToolResultCommonModelInput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "system",
	}
	return json.Marshal(marshaler)
}

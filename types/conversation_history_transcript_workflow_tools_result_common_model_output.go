package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput struct {
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
	Result                 *WorkflowToolResponseModelOutput `json:"result,omitempty" xml:"result,omitempty"`
	extraProperties        map[string]interface{}
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetResultValue() (value string) {
	if c == nil {
		return
	}
	return c.ResultValue
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetIsBlocked() (value *bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetToolLatencySecs() (value *float64) {
	if c == nil {
		return
	}
	return c.ToolLatencySecs
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetErrorType() (value *string) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetRawErrorMessage() (value *string) {
	if c == nil {
		return
	}
	return c.RawErrorMessage
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetDynamicVariableUpdates() (value []*DynamicVariableUpdateCommonModel) {
	if c == nil {
		return
	}
	return c.DynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetResult() (value *WorkflowToolResponseModelOutput) {
	if c == nil {
		return
	}
	return c.Result
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetResultValue(resultValue string) {
	c.ResultValue = resultValue
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetIsBlocked(isBlocked *bool) {
	c.IsBlocked = isBlocked
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetToolLatencySecs(toolLatencySecs *float64) {
	c.ToolLatencySecs = toolLatencySecs
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetErrorType(errorType *string) {
	c.ErrorType = errorType
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetRawErrorMessage(rawErrorMessage *string) {
	c.RawErrorMessage = rawErrorMessage
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetDynamicVariableUpdates(dynamicVariableUpdates []*DynamicVariableUpdateCommonModel) {
	c.DynamicVariableUpdates = dynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) SetResult(result *WorkflowToolResponseModelOutput) {
	c.Result = result
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) Type() string {
	return c.type_
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "workflow" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "workflow", *literals.Type)
	}
	c.type_ = "workflow"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "workflow",
	}
	return json.Marshal(marshaler)
}

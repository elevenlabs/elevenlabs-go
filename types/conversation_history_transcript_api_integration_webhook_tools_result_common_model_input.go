package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput struct {
	RequestId string `json:"request_id" xml:"request_id" required:"true"`
	ToolName  string `json:"tool_name" xml:"tool_name" required:"true"`
	// ID of the tool that produced this result: a workspace tool ID, 'mcp:<server_id>:<tool_name>' for MCP tools, or '__system_<tool_name>' for system tools. Not set for conversations recorded before tool IDs were tracked.
	ToolId                  *string                             `json:"tool_id,omitempty" xml:"tool_id,omitempty"`
	ResultValue             string                              `json:"result_value" xml:"result_value" required:"true"`
	IsError                 bool                                `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked               *bool                               `json:"is_blocked,omitempty" xml:"is_blocked,omitempty"`
	ToolHasBeenCalled       bool                                `json:"tool_has_been_called" xml:"tool_has_been_called" required:"true"`
	ToolLatencySecs         *float64                            `json:"tool_latency_secs,omitempty" xml:"tool_latency_secs,omitempty"`
	ErrorType               *string                             `json:"error_type,omitempty" xml:"error_type,omitempty"`
	RawErrorMessage         *string                             `json:"raw_error_message,omitempty" xml:"raw_error_message,omitempty"`
	DynamicVariableUpdates  []*DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates,omitempty" xml:"dynamic_variable_updates,omitempty"`
	type_                   string
	IntegrationId           *string `json:"integration_id,omitempty" xml:"integration_id,omitempty"`
	CredentialId            *string `json:"credential_id,omitempty" xml:"credential_id,omitempty"`
	IntegrationConnectionId *string `json:"integration_connection_id,omitempty" xml:"integration_connection_id,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetResultValue() (value string) {
	if c == nil {
		return
	}
	return c.ResultValue
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetIsBlocked() (value *bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetToolLatencySecs() (value *float64) {
	if c == nil {
		return
	}
	return c.ToolLatencySecs
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetErrorType() (value *string) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetRawErrorMessage() (value *string) {
	if c == nil {
		return
	}
	return c.RawErrorMessage
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetDynamicVariableUpdates() (value []*DynamicVariableUpdateCommonModel) {
	if c == nil {
		return
	}
	return c.DynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetIntegrationId() (value *string) {
	if c == nil {
		return
	}
	return c.IntegrationId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetCredentialId() (value *string) {
	if c == nil {
		return
	}
	return c.CredentialId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetIntegrationConnectionId() (value *string) {
	if c == nil {
		return
	}
	return c.IntegrationConnectionId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetResultValue(resultValue string) {
	c.ResultValue = resultValue
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetIsBlocked(isBlocked *bool) {
	c.IsBlocked = isBlocked
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetToolLatencySecs(toolLatencySecs *float64) {
	c.ToolLatencySecs = toolLatencySecs
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetErrorType(errorType *string) {
	c.ErrorType = errorType
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetRawErrorMessage(rawErrorMessage *string) {
	c.RawErrorMessage = rawErrorMessage
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetDynamicVariableUpdates(dynamicVariableUpdates []*DynamicVariableUpdateCommonModel) {
	c.DynamicVariableUpdates = dynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetIntegrationId(integrationId *string) {
	c.IntegrationId = integrationId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetCredentialId(credentialId *string) {
	c.CredentialId = credentialId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) SetIntegrationConnectionId(integrationConnectionId *string) {
	c.IntegrationConnectionId = integrationConnectionId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) Type() string {
	return c.type_
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "api_integration_webhook" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "api_integration_webhook", *literals.Type)
	}
	c.type_ = "api_integration_webhook"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "api_integration_webhook",
	}
	return json.Marshal(marshaler)
}

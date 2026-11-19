package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput struct {
	RequestId string `json:"request_id" xml:"request_id" required:"true"`
	ToolName  string `json:"tool_name" xml:"tool_name" required:"true"`
	// ID of the tool that produced this result: a workspace tool ID, 'mcp:<server_id>:<tool_name>' for MCP tools, or '__system_<tool_name>' for system tools. Not set for conversations recorded before tool IDs were tracked.
	ToolId                  *string                             `json:"tool_id" xml:"tool_id" required:"true"`
	ResultValue             string                              `json:"result_value" xml:"result_value" required:"true"`
	IsError                 bool                                `json:"is_error" xml:"is_error" required:"true"`
	IsBlocked               bool                                `json:"is_blocked" xml:"is_blocked" required:"true"`
	ToolHasBeenCalled       bool                                `json:"tool_has_been_called" xml:"tool_has_been_called" required:"true"`
	ToolLatencySecs         float64                             `json:"tool_latency_secs" xml:"tool_latency_secs" required:"true"`
	ErrorType               string                              `json:"error_type" xml:"error_type" required:"true"`
	RawErrorMessage         string                              `json:"raw_error_message" xml:"raw_error_message" required:"true"`
	DynamicVariableUpdates  []*DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates" xml:"dynamic_variable_updates" required:"true"`
	type_                   string
	IntegrationId           string `json:"integration_id" xml:"integration_id" required:"true"`
	CredentialId            string `json:"credential_id" xml:"credential_id" required:"true"`
	IntegrationConnectionId string `json:"integration_connection_id" xml:"integration_connection_id" required:"true"`
	extraProperties         map[string]interface{}
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetResultValue() (value string) {
	if c == nil {
		return
	}
	return c.ResultValue
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetIsBlocked() (value bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetToolLatencySecs() (value float64) {
	if c == nil {
		return
	}
	return c.ToolLatencySecs
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetErrorType() (value string) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetRawErrorMessage() (value string) {
	if c == nil {
		return
	}
	return c.RawErrorMessage
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetDynamicVariableUpdates() (value []*DynamicVariableUpdateCommonModel) {
	if c == nil {
		return
	}
	return c.DynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetIntegrationId() (value string) {
	if c == nil {
		return
	}
	return c.IntegrationId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetCredentialId() (value string) {
	if c == nil {
		return
	}
	return c.CredentialId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetIntegrationConnectionId() (value string) {
	if c == nil {
		return
	}
	return c.IntegrationConnectionId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetResultValue(resultValue string) {
	c.ResultValue = resultValue
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetIsBlocked(isBlocked bool) {
	c.IsBlocked = isBlocked
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetToolLatencySecs(toolLatencySecs float64) {
	c.ToolLatencySecs = toolLatencySecs
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetErrorType(errorType string) {
	c.ErrorType = errorType
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetRawErrorMessage(rawErrorMessage string) {
	c.RawErrorMessage = rawErrorMessage
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetDynamicVariableUpdates(dynamicVariableUpdates []*DynamicVariableUpdateCommonModel) {
	c.DynamicVariableUpdates = dynamicVariableUpdates
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetIntegrationId(integrationId string) {
	c.IntegrationId = integrationId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetCredentialId(credentialId string) {
	c.CredentialId = credentialId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) SetIntegrationConnectionId(integrationConnectionId string) {
	c.IntegrationConnectionId = integrationConnectionId
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) Type() string {
	return c.type_
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "api_integration_webhook",
	}
	return json.Marshal(marshaler)
}

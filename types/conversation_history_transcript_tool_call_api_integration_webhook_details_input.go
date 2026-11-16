package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput struct {
	IntegrationId           *string                                                          `json:"integration_id,omitempty" xml:"integration_id,omitempty"`
	CredentialId            *string                                                          `json:"credential_id,omitempty" xml:"credential_id,omitempty"`
	IntegrationConnectionId *string                                                          `json:"integration_connection_id,omitempty" xml:"integration_connection_id,omitempty"`
	WebhookDetails          *ConversationHistoryTranscriptToolCallWebhookDetails             `json:"webhook_details" xml:"webhook_details" required:"true"`
	FunctionDetails         *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails `json:"function_details,omitempty" xml:"function_details,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) GetIntegrationId() (value *string) {
	if c == nil {
		return
	}
	return c.IntegrationId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) GetCredentialId() (value *string) {
	if c == nil {
		return
	}
	return c.CredentialId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) GetIntegrationConnectionId() (value *string) {
	if c == nil {
		return
	}
	return c.IntegrationConnectionId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) GetWebhookDetails() (value *ConversationHistoryTranscriptToolCallWebhookDetails) {
	if c == nil {
		return
	}
	return c.WebhookDetails
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) GetFunctionDetails() (value *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) {
	if c == nil {
		return
	}
	return c.FunctionDetails
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) SetIntegrationId(integrationId *string) {
	c.IntegrationId = integrationId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) SetCredentialId(credentialId *string) {
	c.CredentialId = credentialId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) SetIntegrationConnectionId(integrationConnectionId *string) {
	c.IntegrationConnectionId = integrationConnectionId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) SetWebhookDetails(webhookDetails *ConversationHistoryTranscriptToolCallWebhookDetails) {
	c.WebhookDetails = webhookDetails
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) SetFunctionDetails(functionDetails *ConversationHistoryTranscriptToolCallIntegrationFunctionDetails) {
	c.FunctionDetails = functionDetails
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

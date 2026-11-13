package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput struct {
	type_                   string
	IntegrationId           string                                               `json:"integration_id" xml:"integration_id" required:"true"`
	CredentialId            string                                               `json:"credential_id" xml:"credential_id" required:"true"`
	IntegrationConnectionId string                                               `json:"integration_connection_id" xml:"integration_connection_id" required:"true"`
	WebhookDetails          *ConversationHistoryTranscriptToolCallWebhookDetails `json:"webhook_details" xml:"webhook_details" required:"true"`
	extraProperties         map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) GetIntegrationId() (value string) {
	if c == nil {
		return
	}
	return c.IntegrationId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) GetCredentialId() (value string) {
	if c == nil {
		return
	}
	return c.CredentialId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) GetIntegrationConnectionId() (value string) {
	if c == nil {
		return
	}
	return c.IntegrationConnectionId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) GetWebhookDetails() (value *ConversationHistoryTranscriptToolCallWebhookDetails) {
	if c == nil {
		return
	}
	return c.WebhookDetails
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) SetIntegrationId(integrationId string) {
	c.IntegrationId = integrationId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) SetCredentialId(credentialId string) {
	c.CredentialId = credentialId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) SetIntegrationConnectionId(integrationConnectionId string) {
	c.IntegrationConnectionId = integrationConnectionId
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) SetWebhookDetails(webhookDetails *ConversationHistoryTranscriptToolCallWebhookDetails) {
	c.WebhookDetails = webhookDetails
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) Type() string {
	return c.type_
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput(tmp)
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

func (c *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "api_integration_webhook",
	}
	return json.Marshal(marshaler)
}

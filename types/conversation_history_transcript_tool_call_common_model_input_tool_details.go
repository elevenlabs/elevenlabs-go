package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationHistoryTranscriptToolCallCommonModelInputToolDetails struct {
	Type                  string
	ApiIntegrationWebhook *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput
	Client                *ConversationHistoryTranscriptToolCallClientDetails
	Mcp                   *ConversationHistoryTranscriptToolCallMcpDetails
	Webhook               *ConversationHistoryTranscriptToolCallWebhookDetails
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) GetApiIntegrationWebhook() *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput {
	if c == nil {
		return nil
	}
	if c.ApiIntegrationWebhook == nil {
		return nil
	}
	return c.ApiIntegrationWebhook
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) GetClient() *ConversationHistoryTranscriptToolCallClientDetails {
	if c == nil {
		return nil
	}
	if c.Client == nil {
		return nil
	}
	return c.Client
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) GetMcp() *ConversationHistoryTranscriptToolCallMcpDetails {
	if c == nil {
		return nil
	}
	if c.Mcp == nil {
		return nil
	}
	return c.Mcp
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) GetWebhook() *ConversationHistoryTranscriptToolCallWebhookDetails {
	if c == nil {
		return nil
	}
	if c.Webhook == nil {
		return nil
	}
	return c.Webhook
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", c)
	}
	switch unmarshaler.Type {
	case "api_integration_webhook":
		value := new(ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.ApiIntegrationWebhook = value
	case "client":
		value := new(ConversationHistoryTranscriptToolCallClientDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Client = value
	case "mcp":
		value := new(ConversationHistoryTranscriptToolCallMcpDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Mcp = value
	case "webhook":
		value := new(ConversationHistoryTranscriptToolCallWebhookDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Webhook = value
	}
	return nil
}

func (c ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.ApiIntegrationWebhook != nil {
		return internal.MarshalJSONWithDiscriminant(c.ApiIntegrationWebhook, "type", "api_integration_webhook")
	}
	if c.Client != nil {
		return internal.MarshalJSONWithDiscriminant(c.Client, "type", "client")
	}
	if c.Mcp != nil {
		return internal.MarshalJSONWithDiscriminant(c.Mcp, "type", "mcp")
	}
	if c.Webhook != nil {
		return internal.MarshalJSONWithDiscriminant(c.Webhook, "type", "webhook")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ConversationHistoryTranscriptToolCallCommonModelInputToolDetailsVisitor interface {
	VisitApiIntegrationWebhook(*ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput) error
	VisitClient(*ConversationHistoryTranscriptToolCallClientDetails) error
	VisitMcp(*ConversationHistoryTranscriptToolCallMcpDetails) error
	VisitWebhook(*ConversationHistoryTranscriptToolCallWebhookDetails) error
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) Accept(visitor ConversationHistoryTranscriptToolCallCommonModelInputToolDetailsVisitor) error {
	if c.ApiIntegrationWebhook != nil {
		return visitor.VisitApiIntegrationWebhook(c.ApiIntegrationWebhook)
	}
	if c.Client != nil {
		return visitor.VisitClient(c.Client)
	}
	if c.Mcp != nil {
		return visitor.VisitMcp(c.Mcp)
	}
	if c.Webhook != nil {
		return visitor.VisitWebhook(c.Webhook)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.ApiIntegrationWebhook != nil {
		fields = append(fields, "api_integration_webhook")
	}
	if c.Client != nil {
		fields = append(fields, "client")
	}
	if c.Mcp != nil {
		fields = append(fields, "mcp")
	}
	if c.Webhook != nil {
		fields = append(fields, "webhook")
	}
	if len(fields) == 0 {
		if c.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Type)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Type != "" {
		field := fields[0]
		if c.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Type,
				c,
			)
		}
	}
	return nil
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ToolExecutionResponseModelToolCallDetails struct {
	Type                  string
	ApiIntegrationWebhook *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput
	Client                *ConversationHistoryTranscriptToolCallClientDetails
	Mcp                   *ConversationHistoryTranscriptToolCallMcpDetails
	Webhook               *ConversationHistoryTranscriptToolCallWebhookDetails
}

func (t *ToolExecutionResponseModelToolCallDetails) GetApiIntegrationWebhook() *ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput {
	if t == nil {
		return nil
	}
	if t.ApiIntegrationWebhook == nil {
		return nil
	}
	return t.ApiIntegrationWebhook
}

func (t *ToolExecutionResponseModelToolCallDetails) GetClient() *ConversationHistoryTranscriptToolCallClientDetails {
	if t == nil {
		return nil
	}
	if t.Client == nil {
		return nil
	}
	return t.Client
}

func (t *ToolExecutionResponseModelToolCallDetails) GetMcp() *ConversationHistoryTranscriptToolCallMcpDetails {
	if t == nil {
		return nil
	}
	if t.Mcp == nil {
		return nil
	}
	return t.Mcp
}

func (t *ToolExecutionResponseModelToolCallDetails) GetWebhook() *ConversationHistoryTranscriptToolCallWebhookDetails {
	if t == nil {
		return nil
	}
	if t.Webhook == nil {
		return nil
	}
	return t.Webhook
}

func (t *ToolExecutionResponseModelToolCallDetails) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *ToolExecutionResponseModelToolCallDetails) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", t)
	}
	switch unmarshaler.Type {
	case "api_integration_webhook":
		value := new(ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ApiIntegrationWebhook = value
	case "client":
		value := new(ConversationHistoryTranscriptToolCallClientDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Client = value
	case "mcp":
		value := new(ConversationHistoryTranscriptToolCallMcpDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Mcp = value
	case "webhook":
		value := new(ConversationHistoryTranscriptToolCallWebhookDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Webhook = value
	}
	return nil
}

func (t ToolExecutionResponseModelToolCallDetails) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.ApiIntegrationWebhook != nil {
		return internal.MarshalJSONWithDiscriminant(t.ApiIntegrationWebhook, "type", "api_integration_webhook")
	}
	if t.Client != nil {
		return internal.MarshalJSONWithDiscriminant(t.Client, "type", "client")
	}
	if t.Mcp != nil {
		return internal.MarshalJSONWithDiscriminant(t.Mcp, "type", "mcp")
	}
	if t.Webhook != nil {
		return internal.MarshalJSONWithDiscriminant(t.Webhook, "type", "webhook")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type ToolExecutionResponseModelToolCallDetailsVisitor interface {
	VisitApiIntegrationWebhook(*ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsOutput) error
	VisitClient(*ConversationHistoryTranscriptToolCallClientDetails) error
	VisitMcp(*ConversationHistoryTranscriptToolCallMcpDetails) error
	VisitWebhook(*ConversationHistoryTranscriptToolCallWebhookDetails) error
}

func (t *ToolExecutionResponseModelToolCallDetails) Accept(visitor ToolExecutionResponseModelToolCallDetailsVisitor) error {
	if t.ApiIntegrationWebhook != nil {
		return visitor.VisitApiIntegrationWebhook(t.ApiIntegrationWebhook)
	}
	if t.Client != nil {
		return visitor.VisitClient(t.Client)
	}
	if t.Mcp != nil {
		return visitor.VisitMcp(t.Mcp)
	}
	if t.Webhook != nil {
		return visitor.VisitWebhook(t.Webhook)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *ToolExecutionResponseModelToolCallDetails) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.ApiIntegrationWebhook != nil {
		fields = append(fields, "api_integration_webhook")
	}
	if t.Client != nil {
		fields = append(fields, "client")
	}
	if t.Mcp != nil {
		fields = append(fields, "mcp")
	}
	if t.Webhook != nil {
		fields = append(fields, "webhook")
	}
	if len(fields) == 0 {
		if t.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.Type)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.Type != "" {
		field := fields[0]
		if t.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.Type,
				t,
			)
		}
	}
	return nil
}

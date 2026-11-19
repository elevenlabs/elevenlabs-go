package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The type of tool
type PromptAgentApiModelInputToolsItem struct {
	Type                  string
	ApiIntegrationWebhook *ApiIntegrationWebhookToolConfigInput
	Client                *ClientToolConfigInput
	Smb                   *SmbToolConfig
	System                *SystemToolConfigInput
	Webhook               *WebhookToolConfigInput
}

func (p *PromptAgentApiModelInputToolsItem) GetApiIntegrationWebhook() *ApiIntegrationWebhookToolConfigInput {
	if p == nil {
		return nil
	}
	if p.ApiIntegrationWebhook == nil {
		return nil
	}
	return p.ApiIntegrationWebhook
}

func (p *PromptAgentApiModelInputToolsItem) GetClient() *ClientToolConfigInput {
	if p == nil {
		return nil
	}
	if p.Client == nil {
		return nil
	}
	return p.Client
}

func (p *PromptAgentApiModelInputToolsItem) GetSmb() *SmbToolConfig {
	if p == nil {
		return nil
	}
	if p.Smb == nil {
		return nil
	}
	return p.Smb
}

func (p *PromptAgentApiModelInputToolsItem) GetSystem() *SystemToolConfigInput {
	if p == nil {
		return nil
	}
	if p.System == nil {
		return nil
	}
	return p.System
}

func (p *PromptAgentApiModelInputToolsItem) GetWebhook() *WebhookToolConfigInput {
	if p == nil {
		return nil
	}
	if p.Webhook == nil {
		return nil
	}
	return p.Webhook
}

func (p *PromptAgentApiModelInputToolsItem) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *PromptAgentApiModelInputToolsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", p)
	}
	switch unmarshaler.Type {
	case "api_integration_webhook":
		value := new(ApiIntegrationWebhookToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.ApiIntegrationWebhook = value
	case "client":
		value := new(ClientToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Client = value
	case "smb":
		value := new(SmbToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Smb = value
	case "system":
		value := new(SystemToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.System = value
	case "webhook":
		value := new(WebhookToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Webhook = value
	}
	return nil
}

func (p PromptAgentApiModelInputToolsItem) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.ApiIntegrationWebhook != nil {
		return internal.MarshalJSONWithDiscriminant(p.ApiIntegrationWebhook, "type", "api_integration_webhook")
	}
	if p.Client != nil {
		return internal.MarshalJSONWithDiscriminant(p.Client, "type", "client")
	}
	if p.Smb != nil {
		return internal.MarshalJSONWithDiscriminant(p.Smb, "type", "smb")
	}
	if p.System != nil {
		return internal.MarshalJSONWithDiscriminant(p.System, "type", "system")
	}
	if p.Webhook != nil {
		return internal.MarshalJSONWithDiscriminant(p.Webhook, "type", "webhook")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PromptAgentApiModelInputToolsItemVisitor interface {
	VisitApiIntegrationWebhook(*ApiIntegrationWebhookToolConfigInput) error
	VisitClient(*ClientToolConfigInput) error
	VisitSmb(*SmbToolConfig) error
	VisitSystem(*SystemToolConfigInput) error
	VisitWebhook(*WebhookToolConfigInput) error
}

func (p *PromptAgentApiModelInputToolsItem) Accept(visitor PromptAgentApiModelInputToolsItemVisitor) error {
	if p.ApiIntegrationWebhook != nil {
		return visitor.VisitApiIntegrationWebhook(p.ApiIntegrationWebhook)
	}
	if p.Client != nil {
		return visitor.VisitClient(p.Client)
	}
	if p.Smb != nil {
		return visitor.VisitSmb(p.Smb)
	}
	if p.System != nil {
		return visitor.VisitSystem(p.System)
	}
	if p.Webhook != nil {
		return visitor.VisitWebhook(p.Webhook)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *PromptAgentApiModelInputToolsItem) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.ApiIntegrationWebhook != nil {
		fields = append(fields, "api_integration_webhook")
	}
	if p.Client != nil {
		fields = append(fields, "client")
	}
	if p.Smb != nil {
		fields = append(fields, "smb")
	}
	if p.System != nil {
		fields = append(fields, "system")
	}
	if p.Webhook != nil {
		fields = append(fields, "webhook")
	}
	if len(fields) == 0 {
		if p.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Type)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Type != "" {
		field := fields[0]
		if p.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Type,
				p,
			)
		}
	}
	return nil
}

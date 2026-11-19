package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The type of tool
type PromptAgentApiModelOutputToolsItem struct {
	Type                  string
	ApiIntegrationWebhook *ApiIntegrationWebhookToolConfigOutput
	Client                *ClientToolConfigOutput
	Mcp                   *McpToolConfigOutput
	Smb                   *SmbToolConfig
	System                *SystemToolConfigOutput
	Webhook               *WebhookToolConfigOutput
}

func (p *PromptAgentApiModelOutputToolsItem) GetApiIntegrationWebhook() *ApiIntegrationWebhookToolConfigOutput {
	if p == nil {
		return nil
	}
	if p.ApiIntegrationWebhook == nil {
		return nil
	}
	return p.ApiIntegrationWebhook
}

func (p *PromptAgentApiModelOutputToolsItem) GetClient() *ClientToolConfigOutput {
	if p == nil {
		return nil
	}
	if p.Client == nil {
		return nil
	}
	return p.Client
}

func (p *PromptAgentApiModelOutputToolsItem) GetMcp() *McpToolConfigOutput {
	if p == nil {
		return nil
	}
	if p.Mcp == nil {
		return nil
	}
	return p.Mcp
}

func (p *PromptAgentApiModelOutputToolsItem) GetSmb() *SmbToolConfig {
	if p == nil {
		return nil
	}
	if p.Smb == nil {
		return nil
	}
	return p.Smb
}

func (p *PromptAgentApiModelOutputToolsItem) GetSystem() *SystemToolConfigOutput {
	if p == nil {
		return nil
	}
	if p.System == nil {
		return nil
	}
	return p.System
}

func (p *PromptAgentApiModelOutputToolsItem) GetWebhook() *WebhookToolConfigOutput {
	if p == nil {
		return nil
	}
	if p.Webhook == nil {
		return nil
	}
	return p.Webhook
}

func (p *PromptAgentApiModelOutputToolsItem) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *PromptAgentApiModelOutputToolsItem) UnmarshalJSON(data []byte) error {
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
		value := new(ApiIntegrationWebhookToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.ApiIntegrationWebhook = value
	case "client":
		value := new(ClientToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Client = value
	case "mcp":
		value := new(McpToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Mcp = value
	case "smb":
		value := new(SmbToolConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Smb = value
	case "system":
		value := new(SystemToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.System = value
	case "webhook":
		value := new(WebhookToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Webhook = value
	}
	return nil
}

func (p PromptAgentApiModelOutputToolsItem) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.ApiIntegrationWebhook != nil {
		return internal.MarshalJSONWithDiscriminant(p.ApiIntegrationWebhook, "type", "api_integration_webhook")
	}
	if p.Client != nil {
		return internal.MarshalJSONWithDiscriminant(p.Client, "type", "client")
	}
	if p.Mcp != nil {
		return internal.MarshalJSONWithDiscriminant(p.Mcp, "type", "mcp")
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

type PromptAgentApiModelOutputToolsItemVisitor interface {
	VisitApiIntegrationWebhook(*ApiIntegrationWebhookToolConfigOutput) error
	VisitClient(*ClientToolConfigOutput) error
	VisitMcp(*McpToolConfigOutput) error
	VisitSmb(*SmbToolConfig) error
	VisitSystem(*SystemToolConfigOutput) error
	VisitWebhook(*WebhookToolConfigOutput) error
}

func (p *PromptAgentApiModelOutputToolsItem) Accept(visitor PromptAgentApiModelOutputToolsItemVisitor) error {
	if p.ApiIntegrationWebhook != nil {
		return visitor.VisitApiIntegrationWebhook(p.ApiIntegrationWebhook)
	}
	if p.Client != nil {
		return visitor.VisitClient(p.Client)
	}
	if p.Mcp != nil {
		return visitor.VisitMcp(p.Mcp)
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

func (p *PromptAgentApiModelOutputToolsItem) validate() error {
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
	if p.Mcp != nil {
		fields = append(fields, "mcp")
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

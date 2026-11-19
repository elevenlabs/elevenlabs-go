package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The type of tool
type ToolResponseModelToolConfig struct {
	Type    string
	Client  *ClientToolConfigOutput
	Mcp     *McpToolConfigOutput
	System  *SystemToolConfigOutput
	Webhook *WebhookToolConfigOutput
}

func (t *ToolResponseModelToolConfig) GetClient() *ClientToolConfigOutput {
	if t == nil {
		return nil
	}
	if t.Client == nil {
		return nil
	}
	return t.Client
}

func (t *ToolResponseModelToolConfig) GetMcp() *McpToolConfigOutput {
	if t == nil {
		return nil
	}
	if t.Mcp == nil {
		return nil
	}
	return t.Mcp
}

func (t *ToolResponseModelToolConfig) GetSystem() *SystemToolConfigOutput {
	if t == nil {
		return nil
	}
	if t.System == nil {
		return nil
	}
	return t.System
}

func (t *ToolResponseModelToolConfig) GetWebhook() *WebhookToolConfigOutput {
	if t == nil {
		return nil
	}
	if t.Webhook == nil {
		return nil
	}
	return t.Webhook
}

func (t *ToolResponseModelToolConfig) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *ToolResponseModelToolConfig) UnmarshalJSON(data []byte) error {
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
	case "client":
		value := new(ClientToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Client = value
	case "mcp":
		value := new(McpToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Mcp = value
	case "system":
		value := new(SystemToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.System = value
	case "webhook":
		value := new(WebhookToolConfigOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Webhook = value
	}
	return nil
}

func (t ToolResponseModelToolConfig) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Client != nil {
		return internal.MarshalJSONWithDiscriminant(t.Client, "type", "client")
	}
	if t.Mcp != nil {
		return internal.MarshalJSONWithDiscriminant(t.Mcp, "type", "mcp")
	}
	if t.System != nil {
		return internal.MarshalJSONWithDiscriminant(t.System, "type", "system")
	}
	if t.Webhook != nil {
		return internal.MarshalJSONWithDiscriminant(t.Webhook, "type", "webhook")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type ToolResponseModelToolConfigVisitor interface {
	VisitClient(*ClientToolConfigOutput) error
	VisitMcp(*McpToolConfigOutput) error
	VisitSystem(*SystemToolConfigOutput) error
	VisitWebhook(*WebhookToolConfigOutput) error
}

func (t *ToolResponseModelToolConfig) Accept(visitor ToolResponseModelToolConfigVisitor) error {
	if t.Client != nil {
		return visitor.VisitClient(t.Client)
	}
	if t.Mcp != nil {
		return visitor.VisitMcp(t.Mcp)
	}
	if t.System != nil {
		return visitor.VisitSystem(t.System)
	}
	if t.Webhook != nil {
		return visitor.VisitWebhook(t.Webhook)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *ToolResponseModelToolConfig) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Client != nil {
		fields = append(fields, "client")
	}
	if t.Mcp != nil {
		fields = append(fields, "mcp")
	}
	if t.System != nil {
		fields = append(fields, "system")
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

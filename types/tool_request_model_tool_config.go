package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for the tool
type ToolRequestModelToolConfig struct {
	Type    string
	Client  *ClientToolConfigInput
	System  *SystemToolConfigInput
	Webhook *WebhookToolConfigInput
}

func (t *ToolRequestModelToolConfig) GetClient() *ClientToolConfigInput {
	if t == nil {
		return nil
	}
	if t.Client == nil {
		return nil
	}
	return t.Client
}

func (t *ToolRequestModelToolConfig) GetSystem() *SystemToolConfigInput {
	if t == nil {
		return nil
	}
	if t.System == nil {
		return nil
	}
	return t.System
}

func (t *ToolRequestModelToolConfig) GetWebhook() *WebhookToolConfigInput {
	if t == nil {
		return nil
	}
	if t.Webhook == nil {
		return nil
	}
	return t.Webhook
}

func (t *ToolRequestModelToolConfig) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *ToolRequestModelToolConfig) UnmarshalJSON(data []byte) error {
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
		value := new(ClientToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Client = value
	case "system":
		value := new(SystemToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.System = value
	case "webhook":
		value := new(WebhookToolConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Webhook = value
	}
	return nil
}

func (t ToolRequestModelToolConfig) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Client != nil {
		return internal.MarshalJSONWithDiscriminant(t.Client, "type", "client")
	}
	if t.System != nil {
		return internal.MarshalJSONWithDiscriminant(t.System, "type", "system")
	}
	if t.Webhook != nil {
		return internal.MarshalJSONWithDiscriminant(t.Webhook, "type", "webhook")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type ToolRequestModelToolConfigVisitor interface {
	VisitClient(*ClientToolConfigInput) error
	VisitSystem(*SystemToolConfigInput) error
	VisitWebhook(*WebhookToolConfigInput) error
}

func (t *ToolRequestModelToolConfig) Accept(visitor ToolRequestModelToolConfigVisitor) error {
	if t.Client != nil {
		return visitor.VisitClient(t.Client)
	}
	if t.System != nil {
		return visitor.VisitSystem(t.System)
	}
	if t.Webhook != nil {
		return visitor.VisitWebhook(t.Webhook)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *ToolRequestModelToolConfig) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Client != nil {
		fields = append(fields, "client")
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

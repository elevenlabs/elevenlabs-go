package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AlertingSettingsResponseNotifiersItem struct {
	Type        string
	Integration *AlertingIntegrationNotifierResponse
	Webhook     *AlertingWebhookNotifierResponse
}

func (a *AlertingSettingsResponseNotifiersItem) GetIntegration() *AlertingIntegrationNotifierResponse {
	if a == nil {
		return nil
	}
	if a.Integration == nil {
		return nil
	}
	return a.Integration
}

func (a *AlertingSettingsResponseNotifiersItem) GetWebhook() *AlertingWebhookNotifierResponse {
	if a == nil {
		return nil
	}
	if a.Webhook == nil {
		return nil
	}
	return a.Webhook
}

func (a *AlertingSettingsResponseNotifiersItem) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AlertingSettingsResponseNotifiersItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", a)
	}
	switch unmarshaler.Type {
	case "integration":
		value := new(AlertingIntegrationNotifierResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Integration = value
	case "webhook":
		value := new(AlertingWebhookNotifierResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Webhook = value
	}
	return nil
}

func (a AlertingSettingsResponseNotifiersItem) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.Integration != nil {
		return internal.MarshalJSONWithDiscriminant(a.Integration, "type", "integration")
	}
	if a.Webhook != nil {
		return internal.MarshalJSONWithDiscriminant(a.Webhook, "type", "webhook")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AlertingSettingsResponseNotifiersItemVisitor interface {
	VisitIntegration(*AlertingIntegrationNotifierResponse) error
	VisitWebhook(*AlertingWebhookNotifierResponse) error
}

func (a *AlertingSettingsResponseNotifiersItem) Accept(visitor AlertingSettingsResponseNotifiersItemVisitor) error {
	if a.Integration != nil {
		return visitor.VisitIntegration(a.Integration)
	}
	if a.Webhook != nil {
		return visitor.VisitWebhook(a.Webhook)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AlertingSettingsResponseNotifiersItem) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.Integration != nil {
		fields = append(fields, "integration")
	}
	if a.Webhook != nil {
		fields = append(fields, "webhook")
	}
	if len(fields) == 0 {
		if a.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Type)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Type != "" {
		field := fields[0]
		if a.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Type,
				a,
			)
		}
	}
	return nil
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingWebhookHeader struct {
	Name     string `json:"name" xml:"name" required:"true" maxLength:"256"`
	IsSecret *bool  `json:"is_secret,omitempty" xml:"is_secret,omitempty"`
	// Header value. May be null only for a secret header on update.
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AlertingWebhookHeader) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AlertingWebhookHeader) GetIsSecret() (value *bool) {
	if a == nil {
		return
	}
	return a.IsSecret
}

func (a *AlertingWebhookHeader) GetValue() (value *string) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AlertingWebhookHeader) SetName(name string) {
	a.Name = name
}

func (a *AlertingWebhookHeader) SetIsSecret(isSecret *bool) {
	a.IsSecret = isSecret
}

func (a *AlertingWebhookHeader) SetValue(value *string) {
	a.Value = value
}

func (a *AlertingWebhookHeader) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingWebhookHeader) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingWebhookHeader to string"
	}
	return string(jsonData)
}

func (a *AlertingWebhookHeader) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingWebhookHeader) MarshalJSON() ([]byte, error) {
	type embed AlertingWebhookHeader
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

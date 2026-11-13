package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A custom webhook header value, encrypted at rest.
type AlertingWebhookSecretHeaderValue struct {
	EncryptedValue  string `json:"encrypted_value" xml:"encrypted_value" required:"true"`
	Nonce           string `json:"nonce" xml:"nonce" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AlertingWebhookSecretHeaderValue) GetEncryptedValue() (value string) {
	if a == nil {
		return
	}
	return a.EncryptedValue
}

func (a *AlertingWebhookSecretHeaderValue) GetNonce() (value string) {
	if a == nil {
		return
	}
	return a.Nonce
}

func (a *AlertingWebhookSecretHeaderValue) SetEncryptedValue(encryptedValue string) {
	a.EncryptedValue = encryptedValue
}

func (a *AlertingWebhookSecretHeaderValue) SetNonce(nonce string) {
	a.Nonce = nonce
}

func (a *AlertingWebhookSecretHeaderValue) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingWebhookSecretHeaderValue) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingWebhookSecretHeaderValue to string"
	}
	return string(jsonData)
}

func (a *AlertingWebhookSecretHeaderValue) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingWebhookSecretHeaderValue
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingWebhookSecretHeaderValue(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingWebhookSecretHeaderValue) MarshalJSON() ([]byte, error) {
	type embed AlertingWebhookSecretHeaderValue
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

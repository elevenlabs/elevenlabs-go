package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingWebhookNotifier struct {
	// ID of the workspace webhook to deliver alert lifecycle notifications to.
	WebhookId       string `json:"webhook_id" xml:"webhook_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AlertingWebhookNotifier) GetWebhookId() (value string) {
	if a == nil {
		return
	}
	return a.WebhookId
}

func (a *AlertingWebhookNotifier) SetWebhookId(webhookId string) {
	a.WebhookId = webhookId
}

func (a *AlertingWebhookNotifier) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingWebhookNotifier) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingWebhookNotifier to string"
	}
	return string(jsonData)
}

func (a *AlertingWebhookNotifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingWebhookNotifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingWebhookNotifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingWebhookNotifier) MarshalJSON() ([]byte, error) {
	type embed AlertingWebhookNotifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

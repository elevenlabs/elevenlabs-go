package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingWebhookNotifierResponse struct {
	WebhookId       string `json:"webhook_id" xml:"webhook_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AlertingWebhookNotifierResponse) GetWebhookId() (value string) {
	if a == nil {
		return
	}
	return a.WebhookId
}

func (a *AlertingWebhookNotifierResponse) SetWebhookId(webhookId string) {
	a.WebhookId = webhookId
}

func (a *AlertingWebhookNotifierResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingWebhookNotifierResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingWebhookNotifierResponse to string"
	}
	return string(jsonData)
}

func (a *AlertingWebhookNotifierResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingWebhookNotifierResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingWebhookNotifierResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingWebhookNotifierResponse) MarshalJSON() ([]byte, error) {
	type embed AlertingWebhookNotifierResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

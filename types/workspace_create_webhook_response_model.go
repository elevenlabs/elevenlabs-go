package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceCreateWebhookResponseModel struct {
	WebhookId       string  `json:"webhook_id" xml:"webhook_id" required:"true"`
	WebhookSecret   *string `json:"webhook_secret,omitempty" xml:"webhook_secret,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceCreateWebhookResponseModel) GetWebhookId() (value string) {
	if w == nil {
		return
	}
	return w.WebhookId
}

func (w *WorkspaceCreateWebhookResponseModel) GetWebhookSecret() (value *string) {
	if w == nil {
		return
	}
	return w.WebhookSecret
}

func (w *WorkspaceCreateWebhookResponseModel) SetWebhookId(webhookId string) {
	w.WebhookId = webhookId
}

func (w *WorkspaceCreateWebhookResponseModel) SetWebhookSecret(webhookSecret *string) {
	w.WebhookSecret = webhookSecret
}

func (w *WorkspaceCreateWebhookResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceCreateWebhookResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceCreateWebhookResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceCreateWebhookResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceCreateWebhookResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceCreateWebhookResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

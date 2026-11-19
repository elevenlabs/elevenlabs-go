package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost struct {
	// The ID of the webhook.
	WebhookId   string       `json:"-" url:"-"`
	WebhookType WebhookTypes `json:"webhook_type" xml:"webhook_type" required:"true"`
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) GetWebhookId() (value string) {
	if b == nil {
		return
	}
	return b.WebhookId
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) GetWebhookType() (value WebhookTypes) {
	if b == nil {
		return
	}
	return b.WebhookType
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) SetWebhookId(webhookId string) {
	b.WebhookId = webhookId
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) SetWebhookType(webhookType WebhookTypes) {
	b.WebhookType = webhookType
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost to string"
	}
	return string(jsonData)
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost(tmp)
	return nil
}

func (b *BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost) MarshalJSON() ([]byte, error) {
	type embed BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

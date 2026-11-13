package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost struct {
	// Settings for creating an HMAC-authenticated webhook
	Settings *WebhookHmacSettings `json:"settings" xml:"settings" required:"true"`
}

func (b *BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost) GetSettings() (value *WebhookHmacSettings) {
	if b == nil {
		return
	}
	return b.Settings
}

func (b *BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost) SetSettings(settings *WebhookHmacSettings) {
	b.Settings = settings
}

func (b *BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost(tmp)
	return nil
}

func (b *BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

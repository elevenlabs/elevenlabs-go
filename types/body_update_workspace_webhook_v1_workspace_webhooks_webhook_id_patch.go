package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch struct {
	// The unique ID for the webhook
	WebhookId string `json:"-" url:"-"`
	// Whether to disable or enable the webhook
	IsDisabled bool `json:"is_disabled" xml:"is_disabled" required:"true"`
	// The display name of the webhook (used for display purposes only).
	Name string `json:"name" xml:"name" required:"true"`
	// Whether to enable automatic retries for transient failures (5xx, 429, timeout)
	RetryEnabled *bool `json:"retry_enabled,omitempty" xml:"retry_enabled,omitempty"`
	// A list of request headers to include with the webhook delivery (optional)
	RequestHeaders map[string]*string `json:"request_headers,omitempty"`
	// The complete set of workspace-level events this webhook should be subscribed to. The webhook is added to the events in the list and removed from any not in the list. Omit to leave the current event subscriptions unchanged.
	Events []WorkspaceWebhookEventType `json:"events,omitempty" xml:"events,omitempty"`
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetWebhookId() (value string) {
	if b == nil {
		return
	}
	return b.WebhookId
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetIsDisabled() (value bool) {
	if b == nil {
		return
	}
	return b.IsDisabled
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetRetryEnabled() (value *bool) {
	if b == nil {
		return
	}
	return b.RetryEnabled
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetRequestHeaders() (value map[string]*string) {
	if b == nil {
		return
	}
	return b.RequestHeaders
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetEvents() (value []WorkspaceWebhookEventType) {
	if b == nil {
		return
	}
	return b.Events
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) SetWebhookId(webhookId string) {
	b.WebhookId = webhookId
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) SetIsDisabled(isDisabled bool) {
	b.IsDisabled = isDisabled
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) SetName(name string) {
	b.Name = name
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) SetRetryEnabled(retryEnabled *bool) {
	b.RetryEnabled = retryEnabled
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) SetRequestHeaders(requestHeaders map[string]*string) {
	b.RequestHeaders = requestHeaders
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) SetEvents(events []WorkspaceWebhookEventType) {
	b.Events = events
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

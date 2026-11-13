package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceWebhookResponseModel struct {
	// The display name for this webhook.
	Name string `json:"name" xml:"name" required:"true"`
	// The unique ID for this webhook.
	WebhookId string `json:"webhook_id" xml:"webhook_id" required:"true"`
	// The HTTPS callback URL that is called when this webhook is triggered in the platform.
	WebhookUrl string `json:"webhook_url" xml:"webhook_url" required:"true"`
	// Whether the webhook has been manually disabled by a user.
	IsDisabled bool `json:"is_disabled" xml:"is_disabled" required:"true"`
	// Whether the webhook has been automatically disabled due to repeated consecutive failures over a long period of time.
	IsAutoDisabled bool `json:"is_auto_disabled" xml:"is_auto_disabled" required:"true"`
	// Original creation time of the webhook.
	CreatedAtUnix int                   `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	AuthType      WebhookAuthMethodType `json:"auth_type" xml:"auth_type" required:"true"`
	// The list of products that are currently configured to trigger this webhook.
	Usage []*WorkspaceWebhookUsageResponseModel `json:"usage,omitempty" xml:"usage,omitempty"`
	// The workspace-level events this webhook is currently subscribed to. Only populated when usages are requested.
	Events []WorkspaceWebhookEventType `json:"events,omitempty" xml:"events,omitempty"`
	// The most recent error code returned from the callback URL.
	MostRecentFailureErrorCode *int `json:"most_recent_failure_error_code,omitempty" xml:"most_recent_failure_error_code,omitempty"`
	// The most recent time the webhook failed, failures are any non-200 codes returned by the callback URL.
	MostRecentFailureTimestamp *int `json:"most_recent_failure_timestamp,omitempty" xml:"most_recent_failure_timestamp,omitempty"`
	extraProperties            map[string]interface{}
}

func (w *WorkspaceWebhookResponseModel) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WorkspaceWebhookResponseModel) GetWebhookId() (value string) {
	if w == nil {
		return
	}
	return w.WebhookId
}

func (w *WorkspaceWebhookResponseModel) GetWebhookUrl() (value string) {
	if w == nil {
		return
	}
	return w.WebhookUrl
}

func (w *WorkspaceWebhookResponseModel) GetIsDisabled() (value bool) {
	if w == nil {
		return
	}
	return w.IsDisabled
}

func (w *WorkspaceWebhookResponseModel) GetIsAutoDisabled() (value bool) {
	if w == nil {
		return
	}
	return w.IsAutoDisabled
}

func (w *WorkspaceWebhookResponseModel) GetCreatedAtUnix() (value int) {
	if w == nil {
		return
	}
	return w.CreatedAtUnix
}

func (w *WorkspaceWebhookResponseModel) GetAuthType() (value WebhookAuthMethodType) {
	if w == nil {
		return
	}
	return w.AuthType
}

func (w *WorkspaceWebhookResponseModel) GetUsage() (value []*WorkspaceWebhookUsageResponseModel) {
	if w == nil {
		return
	}
	return w.Usage
}

func (w *WorkspaceWebhookResponseModel) GetEvents() (value []WorkspaceWebhookEventType) {
	if w == nil {
		return
	}
	return w.Events
}

func (w *WorkspaceWebhookResponseModel) GetMostRecentFailureErrorCode() (value *int) {
	if w == nil {
		return
	}
	return w.MostRecentFailureErrorCode
}

func (w *WorkspaceWebhookResponseModel) GetMostRecentFailureTimestamp() (value *int) {
	if w == nil {
		return
	}
	return w.MostRecentFailureTimestamp
}

func (w *WorkspaceWebhookResponseModel) SetName(name string) {
	w.Name = name
}

func (w *WorkspaceWebhookResponseModel) SetWebhookId(webhookId string) {
	w.WebhookId = webhookId
}

func (w *WorkspaceWebhookResponseModel) SetWebhookUrl(webhookUrl string) {
	w.WebhookUrl = webhookUrl
}

func (w *WorkspaceWebhookResponseModel) SetIsDisabled(isDisabled bool) {
	w.IsDisabled = isDisabled
}

func (w *WorkspaceWebhookResponseModel) SetIsAutoDisabled(isAutoDisabled bool) {
	w.IsAutoDisabled = isAutoDisabled
}

func (w *WorkspaceWebhookResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	w.CreatedAtUnix = createdAtUnix
}

func (w *WorkspaceWebhookResponseModel) SetAuthType(authType WebhookAuthMethodType) {
	w.AuthType = authType
}

func (w *WorkspaceWebhookResponseModel) SetUsage(usage []*WorkspaceWebhookUsageResponseModel) {
	w.Usage = usage
}

func (w *WorkspaceWebhookResponseModel) SetEvents(events []WorkspaceWebhookEventType) {
	w.Events = events
}

func (w *WorkspaceWebhookResponseModel) SetMostRecentFailureErrorCode(mostRecentFailureErrorCode *int) {
	w.MostRecentFailureErrorCode = mostRecentFailureErrorCode
}

func (w *WorkspaceWebhookResponseModel) SetMostRecentFailureTimestamp(mostRecentFailureTimestamp *int) {
	w.MostRecentFailureTimestamp = mostRecentFailureTimestamp
}

func (w *WorkspaceWebhookResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceWebhookResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceWebhookResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceWebhookResponseModel) UnmarshalJSON(data []byte) error {
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

func (w *WorkspaceWebhookResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceWebhookResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Settings for creating an HMAC-authenticated webhook
type WebhookHmacSettings struct {
	// The authentication type for this webhook
	authType string
	// The display name for this webhook
	Name string `json:"name" xml:"name" required:"true"`
	// The HTTPS callback URL that will be called when this webhook is triggered
	WebhookUrl string `json:"webhook_url" xml:"webhook_url" required:"true"`
	// Optional custom request headers to include with each webhook delivery
	RequestHeaders map[string]*string `json:"request_headers,omitempty"`
	// Field names whose values the delivery log hides when it shows the request payload, matched at any depth, ignoring case and punctuation such as _ or -. Fields with secret-looking names, such as tokens and passwords, are always hidden. The stored log and redeliveries keep the original values. At most 50 names, each 1 to 128 characters, with no whitespace, dots or wildcards.
	LogRedactedFields []string `json:"log_redacted_fields,omitempty" xml:"log_redacted_fields,omitempty"`
	extraProperties   map[string]interface{}
}

func (w *WebhookHmacSettings) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WebhookHmacSettings) GetWebhookUrl() (value string) {
	if w == nil {
		return
	}
	return w.WebhookUrl
}

func (w *WebhookHmacSettings) GetRequestHeaders() (value map[string]*string) {
	if w == nil {
		return
	}
	return w.RequestHeaders
}

func (w *WebhookHmacSettings) GetLogRedactedFields() (value []string) {
	if w == nil {
		return
	}
	return w.LogRedactedFields
}

func (w *WebhookHmacSettings) SetName(name string) {
	w.Name = name
}

func (w *WebhookHmacSettings) SetWebhookUrl(webhookUrl string) {
	w.WebhookUrl = webhookUrl
}

func (w *WebhookHmacSettings) SetRequestHeaders(requestHeaders map[string]*string) {
	w.RequestHeaders = requestHeaders
}

func (w *WebhookHmacSettings) SetLogRedactedFields(logRedactedFields []string) {
	w.LogRedactedFields = logRedactedFields
}

func (w *WebhookHmacSettings) AuthType() string {
	return w.authType
}

func (w *WebhookHmacSettings) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookHmacSettings) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookHmacSettings to string"
	}
	return string(jsonData)
}

func (w *WebhookHmacSettings) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		AuthType *string `json:"auth_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.AuthType != nil && *literals.AuthType != "hmac" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "hmac", *literals.AuthType)
	}
	w.authType = "hmac"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "auth_type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebhookHmacSettings) MarshalJSON() ([]byte, error) {
	type embed WebhookHmacSettings
	var marshaler = struct {
		embed
		AuthType string `json:"auth_type"`
	}{
		embed:    embed(*w),
		AuthType: "hmac",
	}
	return json.Marshal(marshaler)
}

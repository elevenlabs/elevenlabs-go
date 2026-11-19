package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InternalAlertingWebhookNotifier struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The URL to send alert lifecycle notifications to.
	Url    string                 `json:"url" xml:"url" required:"true" maxLength:"2048"`
	Method *AlertingWebhookMethod `json:"method,omitempty" xml:"method,omitempty"`
	// Custom request headers sent with every notification. Secret header values are encrypted at rest and never returned by the API.
	Headers         []*AlertingWebhookHeader `json:"headers,omitempty" xml:"headers,omitempty" maxItems:"25"`
	extraProperties map[string]interface{}
}

func (i *InternalAlertingWebhookNotifier) GetType() (value *string) {
	if i == nil {
		return
	}
	return i.Type
}

func (i *InternalAlertingWebhookNotifier) GetUrl() (value string) {
	if i == nil {
		return
	}
	return i.Url
}

func (i *InternalAlertingWebhookNotifier) GetMethod() (value *AlertingWebhookMethod) {
	if i == nil {
		return
	}
	return i.Method
}

func (i *InternalAlertingWebhookNotifier) GetHeaders() (value []*AlertingWebhookHeader) {
	if i == nil {
		return
	}
	return i.Headers
}

func (i *InternalAlertingWebhookNotifier) SetType(type_ *string) {
	i.Type = type_
}

func (i *InternalAlertingWebhookNotifier) SetUrl(url string) {
	i.Url = url
}

func (i *InternalAlertingWebhookNotifier) SetMethod(method *AlertingWebhookMethod) {
	i.Method = method
}

func (i *InternalAlertingWebhookNotifier) SetHeaders(headers []*AlertingWebhookHeader) {
	i.Headers = headers
}

func (i *InternalAlertingWebhookNotifier) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InternalAlertingWebhookNotifier) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InternalAlertingWebhookNotifier to string"
	}
	return string(jsonData)
}

func (i *InternalAlertingWebhookNotifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InternalAlertingWebhookNotifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InternalAlertingWebhookNotifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InternalAlertingWebhookNotifier) MarshalJSON() ([]byte, error) {
	type embed InternalAlertingWebhookNotifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}

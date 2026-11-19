package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationClientDataWebhook struct {
	// The URL to send the webhook to
	Url string `json:"url" xml:"url" required:"true"`
	// The headers to send with the webhook request
	RequestHeaders  map[string]*ConversationInitiationClientDataWebhookRequestHeadersValue `json:"request_headers" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataWebhook) GetUrl() (value string) {
	if c == nil {
		return
	}
	return c.Url
}

func (c *ConversationInitiationClientDataWebhook) GetRequestHeaders() (value map[string]*ConversationInitiationClientDataWebhookRequestHeadersValue) {
	if c == nil {
		return
	}
	return c.RequestHeaders
}

func (c *ConversationInitiationClientDataWebhook) SetUrl(url string) {
	c.Url = url
}

func (c *ConversationInitiationClientDataWebhook) SetRequestHeaders(requestHeaders map[string]*ConversationInitiationClientDataWebhookRequestHeadersValue) {
	c.RequestHeaders = requestHeaders
}

func (c *ConversationInitiationClientDataWebhook) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataWebhook) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataWebhook to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataWebhook) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationInitiationClientDataWebhook
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataWebhook(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataWebhook) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataWebhook
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

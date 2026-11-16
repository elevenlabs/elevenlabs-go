package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallWebhookDetails struct {
	Type            *string           `json:"type,omitempty" xml:"type,omitempty"`
	Method          string            `json:"method" xml:"method" required:"true"`
	Url             string            `json:"url" xml:"url" required:"true"`
	Headers         map[string]string `json:"headers,omitempty"`
	PathParams      map[string]string `json:"path_params,omitempty"`
	QueryParams     map[string]string `json:"query_params,omitempty"`
	Body            *string           `json:"body,omitempty" xml:"body,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetMethod() (value string) {
	if c == nil {
		return
	}
	return c.Method
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetUrl() (value string) {
	if c == nil {
		return
	}
	return c.Url
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetHeaders() (value map[string]string) {
	if c == nil {
		return
	}
	return c.Headers
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetPathParams() (value map[string]string) {
	if c == nil {
		return
	}
	return c.PathParams
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetQueryParams() (value map[string]string) {
	if c == nil {
		return
	}
	return c.QueryParams
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetBody() (value *string) {
	if c == nil {
		return
	}
	return c.Body
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetMethod(method string) {
	c.Method = method
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetUrl(url string) {
	c.Url = url
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetHeaders(headers map[string]string) {
	c.Headers = headers
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetPathParams(pathParams map[string]string) {
	c.PathParams = pathParams
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetQueryParams(queryParams map[string]string) {
	c.QueryParams = queryParams
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) SetBody(body *string) {
	c.Body = body
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallWebhookDetails) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallWebhookDetails to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptToolCallWebhookDetails) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallWebhookDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

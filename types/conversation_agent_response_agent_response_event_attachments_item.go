package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationAgentResponseAgentResponseEventAttachmentsItem struct {
	Url             string  `json:"url" xml:"url" required:"true"`
	Name            string  `json:"name" xml:"name" required:"true"`
	MimeType        *string `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) GetUrl() (value string) {
	if c == nil {
		return
	}
	return c.Url
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) GetMimeType() (value *string) {
	if c == nil {
		return
	}
	return c.MimeType
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) SetUrl(url string) {
	c.Url = url
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) SetName(name string) {
	c.Name = name
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) SetMimeType(mimeType *string) {
	c.MimeType = mimeType
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationAgentResponseAgentResponseEventAttachmentsItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationAgentResponseAgentResponseEventAttachmentsItem to string"
	}
	return string(jsonData)
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationAgentResponseAgentResponseEventAttachmentsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationAgentResponseAgentResponseEventAttachmentsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationAgentResponseAgentResponseEventAttachmentsItem) MarshalJSON() ([]byte, error) {
	type embed ConversationAgentResponseAgentResponseEventAttachmentsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

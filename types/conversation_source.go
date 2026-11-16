package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationSource struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	ConversationId  string  `json:"conversation_id" xml:"conversation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationSource) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationSource) GetConversationId() (value string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *ConversationSource) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationSource) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
}

func (c *ConversationSource) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationSource) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationSource to string"
	}
	return string(jsonData)
}

func (c *ConversationSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationSource(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationSource) MarshalJSON() ([]byte, error) {
	type embed ConversationSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

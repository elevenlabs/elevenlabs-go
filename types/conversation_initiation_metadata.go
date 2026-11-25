package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationMetadata struct {
	ConversationInitiationMetadataEvent *ConversationInitiationMetadataConversationInitiationMetadataEvent `json:"conversation_initiation_metadata_event" xml:"conversation_initiation_metadata_event" required:"true"`
	Type                                *string                                                            `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties                     map[string]interface{}
}

func (c *ConversationInitiationMetadata) GetConversationInitiationMetadataEvent() (value *ConversationInitiationMetadataConversationInitiationMetadataEvent) {
	if c == nil {
		return
	}
	return c.ConversationInitiationMetadataEvent
}

func (c *ConversationInitiationMetadata) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationInitiationMetadata) SetConversationInitiationMetadataEvent(conversationInitiationMetadataEvent *ConversationInitiationMetadataConversationInitiationMetadataEvent) {
	c.ConversationInitiationMetadataEvent = conversationInitiationMetadataEvent
}

func (c *ConversationInitiationMetadata) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationInitiationMetadata) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationMetadata) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationMetadata to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationInitiationMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationInitiationMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationMetadata) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

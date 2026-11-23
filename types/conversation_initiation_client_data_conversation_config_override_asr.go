package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for conversational transcription
type ConversationInitiationClientDataConversationConfigOverrideAsr struct {
	// Keywords to boost prediction probability for
	Keywords        []string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAsr) GetKeywords() (value []string) {
	if c == nil {
		return
	}
	return c.Keywords
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAsr) SetKeywords(keywords []string) {
	c.Keywords = keywords
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAsr) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideAsr) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideAsr to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAsr) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideAsr
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideAsr(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAsr) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideAsr
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationHistoryElevenAssistantCommonModel struct {
	IsElevenAssistant *bool `json:"is_eleven_assistant,omitempty" xml:"is_eleven_assistant,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationHistoryElevenAssistantCommonModel) GetIsElevenAssistant() (value *bool) {
	if c == nil {
		return
	}
	return c.IsElevenAssistant
}

func (c *ConversationHistoryElevenAssistantCommonModel) SetIsElevenAssistant(isElevenAssistant *bool) {
	c.IsElevenAssistant = isElevenAssistant
}

func (c *ConversationHistoryElevenAssistantCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryElevenAssistantCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryElevenAssistantCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryElevenAssistantCommonModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationHistoryElevenAssistantCommonModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationHistoryElevenAssistantCommonModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryElevenAssistantCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryElevenAssistantCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Represents a message from a multi-voice agent.
type ConversationHistoryMultivoiceMessageModel struct {
	Parts           []*ConversationHistoryMultivoiceMessagePartModel `json:"parts" xml:"parts" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryMultivoiceMessageModel) GetParts() (value []*ConversationHistoryMultivoiceMessagePartModel) {
	if c == nil {
		return
	}
	return c.Parts
}

func (c *ConversationHistoryMultivoiceMessageModel) SetParts(parts []*ConversationHistoryMultivoiceMessagePartModel) {
	c.Parts = parts
}

func (c *ConversationHistoryMultivoiceMessageModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryMultivoiceMessageModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryMultivoiceMessageModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryMultivoiceMessageModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationHistoryMultivoiceMessageModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationHistoryMultivoiceMessageModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryMultivoiceMessageModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryMultivoiceMessageModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

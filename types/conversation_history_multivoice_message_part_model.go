package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a single voice part of a multi-voice message.
type ConversationHistoryMultivoiceMessagePartModel struct {
	Text            string  `json:"text" xml:"text" required:"true"`
	VoiceLabel      *string `json:"voice_label" xml:"voice_label" required:"true"`
	TimeInCallSecs  *int    `json:"time_in_call_secs" xml:"time_in_call_secs" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryMultivoiceMessagePartModel) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *ConversationHistoryMultivoiceMessagePartModel) GetVoiceLabel() (value *string) {
	if c == nil {
		return
	}
	return c.VoiceLabel
}

func (c *ConversationHistoryMultivoiceMessagePartModel) GetTimeInCallSecs() (value *int) {
	if c == nil {
		return
	}
	return c.TimeInCallSecs
}

func (c *ConversationHistoryMultivoiceMessagePartModel) SetText(text string) {
	c.Text = text
}

func (c *ConversationHistoryMultivoiceMessagePartModel) SetVoiceLabel(voiceLabel *string) {
	c.VoiceLabel = voiceLabel
}

func (c *ConversationHistoryMultivoiceMessagePartModel) SetTimeInCallSecs(timeInCallSecs *int) {
	c.TimeInCallSecs = timeInCallSecs
}

func (c *ConversationHistoryMultivoiceMessagePartModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryMultivoiceMessagePartModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryMultivoiceMessagePartModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryMultivoiceMessagePartModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryMultivoiceMessagePartModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryMultivoiceMessagePartModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

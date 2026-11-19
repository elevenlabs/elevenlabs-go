package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppOutboundCallResponse struct {
	Success         bool    `json:"success" xml:"success" required:"true"`
	Message         string  `json:"message" xml:"message" required:"true"`
	ConversationId  *string `json:"conversation_id" xml:"conversation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppOutboundCallResponse) GetSuccess() (value bool) {
	if w == nil {
		return
	}
	return w.Success
}

func (w *WhatsAppOutboundCallResponse) GetMessage() (value string) {
	if w == nil {
		return
	}
	return w.Message
}

func (w *WhatsAppOutboundCallResponse) GetConversationId() (value *string) {
	if w == nil {
		return
	}
	return w.ConversationId
}

func (w *WhatsAppOutboundCallResponse) SetSuccess(success bool) {
	w.Success = success
}

func (w *WhatsAppOutboundCallResponse) SetMessage(message string) {
	w.Message = message
}

func (w *WhatsAppOutboundCallResponse) SetConversationId(conversationId *string) {
	w.ConversationId = conversationId
}

func (w *WhatsAppOutboundCallResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppOutboundCallResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppOutboundCallResponse to string"
	}
	return string(jsonData)
}

func (w *WhatsAppOutboundCallResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppOutboundCallResponse) MarshalJSON() ([]byte, error) {
	type embed WhatsAppOutboundCallResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

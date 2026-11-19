package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppOutboundMessageResponse struct {
	ConversationId  string `json:"conversation_id" xml:"conversation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppOutboundMessageResponse) GetConversationId() (value string) {
	if w == nil {
		return
	}
	return w.ConversationId
}

func (w *WhatsAppOutboundMessageResponse) SetConversationId(conversationId string) {
	w.ConversationId = conversationId
}

func (w *WhatsAppOutboundMessageResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppOutboundMessageResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppOutboundMessageResponse to string"
	}
	return string(jsonData)
}

func (w *WhatsAppOutboundMessageResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppOutboundMessageResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppOutboundMessageResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppOutboundMessageResponse) MarshalJSON() ([]byte, error) {
	type embed WhatsAppOutboundMessageResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

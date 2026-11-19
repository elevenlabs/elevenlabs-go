package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppConversationInfo struct {
	Direction                *WhatsAppConversationInfoDirection `json:"direction,omitempty" xml:"direction,omitempty"`
	WhatsappPhoneNumberId    *string                            `json:"whatsapp_phone_number_id,omitempty" xml:"whatsapp_phone_number_id,omitempty"`
	WhatsappUserId           string                             `json:"whatsapp_user_id" xml:"whatsapp_user_id" required:"true"`
	AwaitingFirstUserMessage *bool                              `json:"awaiting_first_user_message,omitempty" xml:"awaiting_first_user_message,omitempty"`
	Recipient                *WhatsAppRecipient                 `json:"recipient,omitempty" xml:"recipient,omitempty"`
	CallPermissionRequest    *WhatsAppCallPermissionRequest     `json:"call_permission_request,omitempty" xml:"call_permission_request,omitempty"`
	CallResult               *CallResult                        `json:"call_result,omitempty" xml:"call_result,omitempty"`
	OpeningMessage           *WhatsAppOpeningMessage            `json:"opening_message,omitempty" xml:"opening_message,omitempty"`
	extraProperties          map[string]interface{}
}

func (w *WhatsAppConversationInfo) GetDirection() (value *WhatsAppConversationInfoDirection) {
	if w == nil {
		return
	}
	return w.Direction
}

func (w *WhatsAppConversationInfo) GetWhatsappPhoneNumberId() (value *string) {
	if w == nil {
		return
	}
	return w.WhatsappPhoneNumberId
}

func (w *WhatsAppConversationInfo) GetWhatsappUserId() (value string) {
	if w == nil {
		return
	}
	return w.WhatsappUserId
}

func (w *WhatsAppConversationInfo) GetAwaitingFirstUserMessage() (value *bool) {
	if w == nil {
		return
	}
	return w.AwaitingFirstUserMessage
}

func (w *WhatsAppConversationInfo) GetRecipient() (value *WhatsAppRecipient) {
	if w == nil {
		return
	}
	return w.Recipient
}

func (w *WhatsAppConversationInfo) GetCallPermissionRequest() (value *WhatsAppCallPermissionRequest) {
	if w == nil {
		return
	}
	return w.CallPermissionRequest
}

func (w *WhatsAppConversationInfo) GetCallResult() (value *CallResult) {
	if w == nil {
		return
	}
	return w.CallResult
}

func (w *WhatsAppConversationInfo) GetOpeningMessage() (value *WhatsAppOpeningMessage) {
	if w == nil {
		return
	}
	return w.OpeningMessage
}

func (w *WhatsAppConversationInfo) SetDirection(direction *WhatsAppConversationInfoDirection) {
	w.Direction = direction
}

func (w *WhatsAppConversationInfo) SetWhatsappPhoneNumberId(whatsappPhoneNumberId *string) {
	w.WhatsappPhoneNumberId = whatsappPhoneNumberId
}

func (w *WhatsAppConversationInfo) SetWhatsappUserId(whatsappUserId string) {
	w.WhatsappUserId = whatsappUserId
}

func (w *WhatsAppConversationInfo) SetAwaitingFirstUserMessage(awaitingFirstUserMessage *bool) {
	w.AwaitingFirstUserMessage = awaitingFirstUserMessage
}

func (w *WhatsAppConversationInfo) SetRecipient(recipient *WhatsAppRecipient) {
	w.Recipient = recipient
}

func (w *WhatsAppConversationInfo) SetCallPermissionRequest(callPermissionRequest *WhatsAppCallPermissionRequest) {
	w.CallPermissionRequest = callPermissionRequest
}

func (w *WhatsAppConversationInfo) SetCallResult(callResult *CallResult) {
	w.CallResult = callResult
}

func (w *WhatsAppConversationInfo) SetOpeningMessage(openingMessage *WhatsAppOpeningMessage) {
	w.OpeningMessage = openingMessage
}

func (w *WhatsAppConversationInfo) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppConversationInfo) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppConversationInfo to string"
	}
	return string(jsonData)
}

func (w *WhatsAppConversationInfo) UnmarshalJSON(data []byte) error {
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

func (w *WhatsAppConversationInfo) MarshalJSON() ([]byte, error) {
	type embed WhatsAppConversationInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

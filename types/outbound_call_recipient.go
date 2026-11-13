package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OutboundCallRecipient struct {
	Id                               *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	PhoneNumber                      *string                                       `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	WhatsappUserId                   *string                                       `json:"whatsapp_user_id,omitempty" xml:"whatsapp_user_id,omitempty"`
	ConversationInitiationClientData *ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	extraProperties                  map[string]interface{}
}

func (o *OutboundCallRecipient) GetId() (value *string) {
	if o == nil {
		return
	}
	return o.Id
}

func (o *OutboundCallRecipient) GetPhoneNumber() (value *string) {
	if o == nil {
		return
	}
	return o.PhoneNumber
}

func (o *OutboundCallRecipient) GetWhatsappUserId() (value *string) {
	if o == nil {
		return
	}
	return o.WhatsappUserId
}

func (o *OutboundCallRecipient) GetConversationInitiationClientData() (value *ConversationInitiationClientDataRequestInput) {
	if o == nil {
		return
	}
	return o.ConversationInitiationClientData
}

func (o *OutboundCallRecipient) SetId(id *string) {
	o.Id = id
}

func (o *OutboundCallRecipient) SetPhoneNumber(phoneNumber *string) {
	o.PhoneNumber = phoneNumber
}

func (o *OutboundCallRecipient) SetWhatsappUserId(whatsappUserId *string) {
	o.WhatsappUserId = whatsappUserId
}

func (o *OutboundCallRecipient) SetConversationInitiationClientData(conversationInitiationClientData *ConversationInitiationClientDataRequestInput) {
	o.ConversationInitiationClientData = conversationInitiationClientData
}

func (o *OutboundCallRecipient) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OutboundCallRecipient) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OutboundCallRecipient to string"
	}
	return string(jsonData)
}

func (o *OutboundCallRecipient) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OutboundCallRecipient) MarshalJSON() ([]byte, error) {
	type embed OutboundCallRecipient
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OutboundCallRecipientResponseModel struct {
	Id                               string                                    `json:"id" xml:"id" required:"true"`
	PhoneNumber                      *string                                   `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	WhatsappUserId                   *string                                   `json:"whatsapp_user_id,omitempty" xml:"whatsapp_user_id,omitempty"`
	Status                           BatchCallRecipientStatus                  `json:"status" xml:"status" required:"true"`
	CreatedAtUnix                    int                                       `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	UpdatedAtUnix                    int                                       `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	ConversationId                   *string                                   `json:"conversation_id" xml:"conversation_id" required:"true"`
	ConversationInitiationClientData *ConversationInitiationClientDataInternal `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	extraProperties                  map[string]interface{}
}

func (o *OutboundCallRecipientResponseModel) GetId() (value string) {
	if o == nil {
		return
	}
	return o.Id
}

func (o *OutboundCallRecipientResponseModel) GetPhoneNumber() (value *string) {
	if o == nil {
		return
	}
	return o.PhoneNumber
}

func (o *OutboundCallRecipientResponseModel) GetWhatsappUserId() (value *string) {
	if o == nil {
		return
	}
	return o.WhatsappUserId
}

func (o *OutboundCallRecipientResponseModel) GetStatus() (value BatchCallRecipientStatus) {
	if o == nil {
		return
	}
	return o.Status
}

func (o *OutboundCallRecipientResponseModel) GetCreatedAtUnix() (value int) {
	if o == nil {
		return
	}
	return o.CreatedAtUnix
}

func (o *OutboundCallRecipientResponseModel) GetUpdatedAtUnix() (value int) {
	if o == nil {
		return
	}
	return o.UpdatedAtUnix
}

func (o *OutboundCallRecipientResponseModel) GetConversationId() (value *string) {
	if o == nil {
		return
	}
	return o.ConversationId
}

func (o *OutboundCallRecipientResponseModel) GetConversationInitiationClientData() (value *ConversationInitiationClientDataInternal) {
	if o == nil {
		return
	}
	return o.ConversationInitiationClientData
}

func (o *OutboundCallRecipientResponseModel) SetId(id string) {
	o.Id = id
}

func (o *OutboundCallRecipientResponseModel) SetPhoneNumber(phoneNumber *string) {
	o.PhoneNumber = phoneNumber
}

func (o *OutboundCallRecipientResponseModel) SetWhatsappUserId(whatsappUserId *string) {
	o.WhatsappUserId = whatsappUserId
}

func (o *OutboundCallRecipientResponseModel) SetStatus(status BatchCallRecipientStatus) {
	o.Status = status
}

func (o *OutboundCallRecipientResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	o.CreatedAtUnix = createdAtUnix
}

func (o *OutboundCallRecipientResponseModel) SetUpdatedAtUnix(updatedAtUnix int) {
	o.UpdatedAtUnix = updatedAtUnix
}

func (o *OutboundCallRecipientResponseModel) SetConversationId(conversationId *string) {
	o.ConversationId = conversationId
}

func (o *OutboundCallRecipientResponseModel) SetConversationInitiationClientData(conversationInitiationClientData *ConversationInitiationClientDataInternal) {
	o.ConversationInitiationClientData = conversationInitiationClientData
}

func (o *OutboundCallRecipientResponseModel) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OutboundCallRecipientResponseModel) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OutboundCallRecipientResponseModel to string"
	}
	return string(jsonData)
}

func (o *OutboundCallRecipientResponseModel) UnmarshalJSON(data []byte) error {
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

func (o *OutboundCallRecipientResponseModel) MarshalJSON() ([]byte, error) {
	type embed OutboundCallRecipientResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubmitCampaignRecipient struct {
	// The number to call, in E.164 format such as `+14155550100`.
	PhoneNumber *string `json:"phone_number,omitempty" xml:"phone_number,omitempty" maxLength:"64"`
	// More numbers to call, in order after `phone_number`. The lead moves on to the next one when calls to its current number keep failing.
	PhoneNumbers []string `json:"phone_numbers,omitempty" xml:"phone_numbers,omitempty" maxItems:"5"`
	// The contact's WhatsApp user ID. WhatsApp campaigns use it instead of the phone numbers.
	WhatsappUserId *string `json:"whatsapp_user_id,omitempty" xml:"whatsapp_user_id,omitempty" maxLength:"256"`
	// Your identifier for the recipient, returned as its lead's `recipient_id`.
	Id *string `json:"id,omitempty" xml:"id,omitempty" maxLength:"256"`
	// IANA timezone to read the calling windows and callback times in for this recipient, overriding the campaign's.
	TimezoneOverride *string `json:"timezone_override,omitempty" xml:"timezone_override,omitempty"`
	// Dynamic variables and overrides for this recipient's calls.
	ConversationInitiationClientData *ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	extraProperties                  map[string]interface{}
}

func (s *SubmitCampaignRecipient) GetPhoneNumber() (value *string) {
	if s == nil {
		return
	}
	return s.PhoneNumber
}

func (s *SubmitCampaignRecipient) GetPhoneNumbers() (value []string) {
	if s == nil {
		return
	}
	return s.PhoneNumbers
}

func (s *SubmitCampaignRecipient) GetWhatsappUserId() (value *string) {
	if s == nil {
		return
	}
	return s.WhatsappUserId
}

func (s *SubmitCampaignRecipient) GetId() (value *string) {
	if s == nil {
		return
	}
	return s.Id
}

func (s *SubmitCampaignRecipient) GetTimezoneOverride() (value *string) {
	if s == nil {
		return
	}
	return s.TimezoneOverride
}

func (s *SubmitCampaignRecipient) GetConversationInitiationClientData() (value *ConversationInitiationClientDataRequestInput) {
	if s == nil {
		return
	}
	return s.ConversationInitiationClientData
}

func (s *SubmitCampaignRecipient) SetPhoneNumber(phoneNumber *string) {
	s.PhoneNumber = phoneNumber
}

func (s *SubmitCampaignRecipient) SetPhoneNumbers(phoneNumbers []string) {
	s.PhoneNumbers = phoneNumbers
}

func (s *SubmitCampaignRecipient) SetWhatsappUserId(whatsappUserId *string) {
	s.WhatsappUserId = whatsappUserId
}

func (s *SubmitCampaignRecipient) SetId(id *string) {
	s.Id = id
}

func (s *SubmitCampaignRecipient) SetTimezoneOverride(timezoneOverride *string) {
	s.TimezoneOverride = timezoneOverride
}

func (s *SubmitCampaignRecipient) SetConversationInitiationClientData(conversationInitiationClientData *ConversationInitiationClientDataRequestInput) {
	s.ConversationInitiationClientData = conversationInitiationClientData
}

func (s *SubmitCampaignRecipient) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitCampaignRecipient) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitCampaignRecipient to string"
	}
	return string(jsonData)
}

func (s *SubmitCampaignRecipient) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitCampaignRecipient) MarshalJSON() ([]byte, error) {
	type embed SubmitCampaignRecipient
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

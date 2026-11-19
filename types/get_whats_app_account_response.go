package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetWhatsAppAccountResponse struct {
	BusinessAccountId          string               `json:"business_account_id" xml:"business_account_id" required:"true"`
	PhoneNumberId              string               `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	BusinessAccountName        string               `json:"business_account_name" xml:"business_account_name" required:"true"`
	PhoneNumberName            string               `json:"phone_number_name" xml:"phone_number_name" required:"true"`
	PhoneNumber                string               `json:"phone_number" xml:"phone_number" required:"true"`
	AccountType                *WhatsAppAccountType `json:"account_type,omitempty" xml:"account_type,omitempty"`
	AssignedAgentId            *string              `json:"assigned_agent_id,omitempty" xml:"assigned_agent_id,omitempty"`
	EnableMessaging            *bool                `json:"enable_messaging,omitempty" xml:"enable_messaging,omitempty"`
	EnableAudioMessageResponse *bool                `json:"enable_audio_message_response,omitempty" xml:"enable_audio_message_response,omitempty"`
	EnableTypingIndicator      *bool                `json:"enable_typing_indicator,omitempty" xml:"enable_typing_indicator,omitempty"`
	AssignedAgentName          *string              `json:"assigned_agent_name" xml:"assigned_agent_name" required:"true"`
	IsTokenExpired             *bool                `json:"is_token_expired,omitempty" xml:"is_token_expired,omitempty"`
	extraProperties            map[string]interface{}
}

func (g *GetWhatsAppAccountResponse) GetBusinessAccountId() (value string) {
	if g == nil {
		return
	}
	return g.BusinessAccountId
}

func (g *GetWhatsAppAccountResponse) GetPhoneNumberId() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumberId
}

func (g *GetWhatsAppAccountResponse) GetBusinessAccountName() (value string) {
	if g == nil {
		return
	}
	return g.BusinessAccountName
}

func (g *GetWhatsAppAccountResponse) GetPhoneNumberName() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumberName
}

func (g *GetWhatsAppAccountResponse) GetPhoneNumber() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumber
}

func (g *GetWhatsAppAccountResponse) GetAccountType() (value *WhatsAppAccountType) {
	if g == nil {
		return
	}
	return g.AccountType
}

func (g *GetWhatsAppAccountResponse) GetAssignedAgentId() (value *string) {
	if g == nil {
		return
	}
	return g.AssignedAgentId
}

func (g *GetWhatsAppAccountResponse) GetEnableMessaging() (value *bool) {
	if g == nil {
		return
	}
	return g.EnableMessaging
}

func (g *GetWhatsAppAccountResponse) GetEnableAudioMessageResponse() (value *bool) {
	if g == nil {
		return
	}
	return g.EnableAudioMessageResponse
}

func (g *GetWhatsAppAccountResponse) GetEnableTypingIndicator() (value *bool) {
	if g == nil {
		return
	}
	return g.EnableTypingIndicator
}

func (g *GetWhatsAppAccountResponse) GetAssignedAgentName() (value *string) {
	if g == nil {
		return
	}
	return g.AssignedAgentName
}

func (g *GetWhatsAppAccountResponse) GetIsTokenExpired() (value *bool) {
	if g == nil {
		return
	}
	return g.IsTokenExpired
}

func (g *GetWhatsAppAccountResponse) SetBusinessAccountId(businessAccountId string) {
	g.BusinessAccountId = businessAccountId
}

func (g *GetWhatsAppAccountResponse) SetPhoneNumberId(phoneNumberId string) {
	g.PhoneNumberId = phoneNumberId
}

func (g *GetWhatsAppAccountResponse) SetBusinessAccountName(businessAccountName string) {
	g.BusinessAccountName = businessAccountName
}

func (g *GetWhatsAppAccountResponse) SetPhoneNumberName(phoneNumberName string) {
	g.PhoneNumberName = phoneNumberName
}

func (g *GetWhatsAppAccountResponse) SetPhoneNumber(phoneNumber string) {
	g.PhoneNumber = phoneNumber
}

func (g *GetWhatsAppAccountResponse) SetAccountType(accountType *WhatsAppAccountType) {
	g.AccountType = accountType
}

func (g *GetWhatsAppAccountResponse) SetAssignedAgentId(assignedAgentId *string) {
	g.AssignedAgentId = assignedAgentId
}

func (g *GetWhatsAppAccountResponse) SetEnableMessaging(enableMessaging *bool) {
	g.EnableMessaging = enableMessaging
}

func (g *GetWhatsAppAccountResponse) SetEnableAudioMessageResponse(enableAudioMessageResponse *bool) {
	g.EnableAudioMessageResponse = enableAudioMessageResponse
}

func (g *GetWhatsAppAccountResponse) SetEnableTypingIndicator(enableTypingIndicator *bool) {
	g.EnableTypingIndicator = enableTypingIndicator
}

func (g *GetWhatsAppAccountResponse) SetAssignedAgentName(assignedAgentName *string) {
	g.AssignedAgentName = assignedAgentName
}

func (g *GetWhatsAppAccountResponse) SetIsTokenExpired(isTokenExpired *bool) {
	g.IsTokenExpired = isTokenExpired
}

func (g *GetWhatsAppAccountResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetWhatsAppAccountResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetWhatsAppAccountResponse to string"
	}
	return string(jsonData)
}

func (g *GetWhatsAppAccountResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetWhatsAppAccountResponse) MarshalJSON() ([]byte, error) {
	type embed GetWhatsAppAccountResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}

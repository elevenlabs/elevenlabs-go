package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SmsConversationInfo struct {
	Direction          SmsConversationInfoDirection `json:"direction" xml:"direction" required:"true"`
	PhoneNumberId      *string                      `json:"phone_number_id,omitempty" xml:"phone_number_id,omitempty"`
	SmsUserPhoneNumber string                       `json:"sms_user_phone_number" xml:"sms_user_phone_number" required:"true"`
	AgentPhoneNumber   *string                      `json:"agent_phone_number,omitempty" xml:"agent_phone_number,omitempty"`
	extraProperties    map[string]interface{}
}

func (s *SmsConversationInfo) GetDirection() (value SmsConversationInfoDirection) {
	if s == nil {
		return
	}
	return s.Direction
}

func (s *SmsConversationInfo) GetPhoneNumberId() (value *string) {
	if s == nil {
		return
	}
	return s.PhoneNumberId
}

func (s *SmsConversationInfo) GetSmsUserPhoneNumber() (value string) {
	if s == nil {
		return
	}
	return s.SmsUserPhoneNumber
}

func (s *SmsConversationInfo) GetAgentPhoneNumber() (value *string) {
	if s == nil {
		return
	}
	return s.AgentPhoneNumber
}

func (s *SmsConversationInfo) SetDirection(direction SmsConversationInfoDirection) {
	s.Direction = direction
}

func (s *SmsConversationInfo) SetPhoneNumberId(phoneNumberId *string) {
	s.PhoneNumberId = phoneNumberId
}

func (s *SmsConversationInfo) SetSmsUserPhoneNumber(smsUserPhoneNumber string) {
	s.SmsUserPhoneNumber = smsUserPhoneNumber
}

func (s *SmsConversationInfo) SetAgentPhoneNumber(agentPhoneNumber *string) {
	s.AgentPhoneNumber = agentPhoneNumber
}

func (s *SmsConversationInfo) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SmsConversationInfo) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SmsConversationInfo to string"
	}
	return string(jsonData)
}

func (s *SmsConversationInfo) UnmarshalJSON(data []byte) error {
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

func (s *SmsConversationInfo) MarshalJSON() ([]byte, error) {
	type embed SmsConversationInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipTrunkOutboundCallResponse struct {
	Success         bool    `json:"success" xml:"success" required:"true"`
	Message         string  `json:"message" xml:"message" required:"true"`
	ConversationId  *string `json:"conversation_id" xml:"conversation_id" required:"true"`
	SipCallId       *string `json:"sip_call_id" xml:"sip_call_id" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SipTrunkOutboundCallResponse) GetSuccess() (value bool) {
	if s == nil {
		return
	}
	return s.Success
}

func (s *SipTrunkOutboundCallResponse) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SipTrunkOutboundCallResponse) GetConversationId() (value *string) {
	if s == nil {
		return
	}
	return s.ConversationId
}

func (s *SipTrunkOutboundCallResponse) GetSipCallId() (value *string) {
	if s == nil {
		return
	}
	return s.SipCallId
}

func (s *SipTrunkOutboundCallResponse) SetSuccess(success bool) {
	s.Success = success
}

func (s *SipTrunkOutboundCallResponse) SetMessage(message string) {
	s.Message = message
}

func (s *SipTrunkOutboundCallResponse) SetConversationId(conversationId *string) {
	s.ConversationId = conversationId
}

func (s *SipTrunkOutboundCallResponse) SetSipCallId(sipCallId *string) {
	s.SipCallId = sipCallId
}

func (s *SipTrunkOutboundCallResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipTrunkOutboundCallResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipTrunkOutboundCallResponse to string"
	}
	return string(jsonData)
}

func (s *SipTrunkOutboundCallResponse) UnmarshalJSON(data []byte) error {
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

func (s *SipTrunkOutboundCallResponse) MarshalJSON() ([]byte, error) {
	type embed SipTrunkOutboundCallResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

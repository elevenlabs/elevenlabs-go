package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost struct {
	AgentId                          string                                              `json:"agent_id" xml:"agent_id" required:"true"`
	AgentPhoneNumberId               string                                              `json:"agent_phone_number_id" xml:"agent_phone_number_id" required:"true"`
	ToNumber                         string                                              `json:"to_number" xml:"to_number" required:"true"`
	ConversationInitiationClientData *types.ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	TelephonyCallConfig              *types.TelephonyCallConfigInput                     `json:"telephony_call_config,omitempty" xml:"telephony_call_config,omitempty"`
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) GetAgentPhoneNumberId() (value string) {
	if b == nil {
		return
	}
	return b.AgentPhoneNumberId
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) GetToNumber() (value string) {
	if b == nil {
		return
	}
	return b.ToNumber
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) GetConversationInitiationClientData() (value *types.ConversationInitiationClientDataRequestInput) {
	if b == nil {
		return
	}
	return b.ConversationInitiationClientData
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) GetTelephonyCallConfig() (value *types.TelephonyCallConfigInput) {
	if b == nil {
		return
	}
	return b.TelephonyCallConfig
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) SetAgentPhoneNumberId(agentPhoneNumberId string) {
	b.AgentPhoneNumberId = agentPhoneNumberId
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) SetToNumber(toNumber string) {
	b.ToNumber = toNumber
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) SetConversationInitiationClientData(conversationInitiationClientData *types.ConversationInitiationClientDataRequestInput) {
	b.ConversationInitiationClientData = conversationInitiationClientData
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) SetTelephonyCallConfig(telephonyCallConfig *types.TelephonyCallConfigInput) {
	b.TelephonyCallConfig = telephonyCallConfig
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost to string"
	}
	return string(jsonData)
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost) MarshalJSON() ([]byte, error) {
	type embed BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

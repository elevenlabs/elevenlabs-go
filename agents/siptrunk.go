package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost struct {
	AgentId                          string                                              `json:"agent_id" xml:"agent_id" required:"true"`
	AgentPhoneNumberId               string                                              `json:"agent_phone_number_id" xml:"agent_phone_number_id" required:"true"`
	ToNumber                         string                                              `json:"to_number" xml:"to_number" required:"true"`
	ConversationInitiationClientData *types.ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	TelephonyCallConfig              *types.TelephonyCallConfigInput                     `json:"telephony_call_config,omitempty" xml:"telephony_call_config,omitempty"`
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) GetAgentPhoneNumberId() (value string) {
	if b == nil {
		return
	}
	return b.AgentPhoneNumberId
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) GetToNumber() (value string) {
	if b == nil {
		return
	}
	return b.ToNumber
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) GetConversationInitiationClientData() (value *types.ConversationInitiationClientDataRequestInput) {
	if b == nil {
		return
	}
	return b.ConversationInitiationClientData
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) GetTelephonyCallConfig() (value *types.TelephonyCallConfigInput) {
	if b == nil {
		return
	}
	return b.TelephonyCallConfig
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) SetAgentPhoneNumberId(agentPhoneNumberId string) {
	b.AgentPhoneNumberId = agentPhoneNumberId
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) SetToNumber(toNumber string) {
	b.ToNumber = toNumber
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) SetConversationInitiationClientData(conversationInitiationClientData *types.ConversationInitiationClientDataRequestInput) {
	b.ConversationInitiationClientData = conversationInitiationClientData
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) SetTelephonyCallConfig(telephonyCallConfig *types.TelephonyCallConfigInput) {
	b.TelephonyCallConfig = telephonyCallConfig
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost to string"
	}
	return string(jsonData)
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost) MarshalJSON() ([]byte, error) {
	type embed BodyHandleAnOutboundCallViaSipTrunkV1ConvaiSipTrunkOutboundCallPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost struct {
	AgentId                          string                                              `json:"agent_id" xml:"agent_id" required:"true"`
	AgentPhoneNumberId               string                                              `json:"agent_phone_number_id" xml:"agent_phone_number_id" required:"true"`
	ToNumber                         string                                              `json:"to_number" xml:"to_number" required:"true"`
	ConversationInitiationClientData *types.ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	// Whether let Twilio record the call.
	CallRecordingEnabled *bool                           `json:"call_recording_enabled,omitempty" xml:"call_recording_enabled,omitempty"`
	TelephonyCallConfig  *types.TelephonyCallConfigInput `json:"telephony_call_config,omitempty" xml:"telephony_call_config,omitempty"`
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetAgentPhoneNumberId() (value string) {
	if b == nil {
		return
	}
	return b.AgentPhoneNumberId
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetToNumber() (value string) {
	if b == nil {
		return
	}
	return b.ToNumber
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetConversationInitiationClientData() (value *types.ConversationInitiationClientDataRequestInput) {
	if b == nil {
		return
	}
	return b.ConversationInitiationClientData
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetCallRecordingEnabled() (value *bool) {
	if b == nil {
		return
	}
	return b.CallRecordingEnabled
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetTelephonyCallConfig() (value *types.TelephonyCallConfigInput) {
	if b == nil {
		return
	}
	return b.TelephonyCallConfig
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) SetAgentPhoneNumberId(agentPhoneNumberId string) {
	b.AgentPhoneNumberId = agentPhoneNumberId
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) SetToNumber(toNumber string) {
	b.ToNumber = toNumber
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) SetConversationInitiationClientData(conversationInitiationClientData *types.ConversationInitiationClientDataRequestInput) {
	b.ConversationInitiationClientData = conversationInitiationClientData
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) SetCallRecordingEnabled(callRecordingEnabled *bool) {
	b.CallRecordingEnabled = callRecordingEnabled
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) SetTelephonyCallConfig(telephonyCallConfig *types.TelephonyCallConfigInput) {
	b.TelephonyCallConfig = telephonyCallConfig
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost to string"
	}
	return string(jsonData)
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost) MarshalJSON() ([]byte, error) {
	type embed BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost struct {
	AgentId                          string                                              `json:"agent_id" xml:"agent_id" required:"true"`
	FromNumber                       string                                              `json:"from_number" xml:"from_number" required:"true"`
	ToNumber                         string                                              `json:"to_number" xml:"to_number" required:"true"`
	Direction                        *types.TelephonyDirection                           `json:"direction,omitempty" xml:"direction,omitempty"`
	ConversationInitiationClientData *types.ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) GetFromNumber() (value string) {
	if b == nil {
		return
	}
	return b.FromNumber
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) GetToNumber() (value string) {
	if b == nil {
		return
	}
	return b.ToNumber
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) GetDirection() (value *types.TelephonyDirection) {
	if b == nil {
		return
	}
	return b.Direction
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) GetConversationInitiationClientData() (value *types.ConversationInitiationClientDataRequestInput) {
	if b == nil {
		return
	}
	return b.ConversationInitiationClientData
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) SetFromNumber(fromNumber string) {
	b.FromNumber = fromNumber
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) SetToNumber(toNumber string) {
	b.ToNumber = toNumber
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) SetDirection(direction *types.TelephonyDirection) {
	b.Direction = direction
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) SetConversationInitiationClientData(conversationInitiationClientData *types.ConversationInitiationClientDataRequestInput) {
	b.ConversationInitiationClientData = conversationInitiationClientData
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost to string"
	}
	return string(jsonData)
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost) MarshalJSON() ([]byte, error) {
	type embed BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

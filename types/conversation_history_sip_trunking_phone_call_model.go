package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistorySipTrunkingPhoneCallModel struct {
	Direction                 TelephonyDirection `json:"direction" xml:"direction" required:"true"`
	PhoneNumberId             string             `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	AgentNumber               string             `json:"agent_number" xml:"agent_number" required:"true"`
	ExternalNumber            string             `json:"external_number" xml:"external_number" required:"true"`
	CallResult                *CallResult        `json:"call_result,omitempty" xml:"call_result,omitempty"`
	type_                     string
	CallId                    *string           `json:"call_id,omitempty" xml:"call_id,omitempty"`
	CallSid                   string            `json:"call_sid" xml:"call_sid" required:"true"`
	SipHeaderDynamicVariables map[string]string `json:"sip_header_dynamic_variables,omitempty"`
	extraProperties           map[string]interface{}
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetDirection() (value TelephonyDirection) {
	if c == nil {
		return
	}
	return c.Direction
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetPhoneNumberId() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumberId
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetAgentNumber() (value string) {
	if c == nil {
		return
	}
	return c.AgentNumber
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetExternalNumber() (value string) {
	if c == nil {
		return
	}
	return c.ExternalNumber
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetCallResult() (value *CallResult) {
	if c == nil {
		return
	}
	return c.CallResult
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetCallId() (value *string) {
	if c == nil {
		return
	}
	return c.CallId
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetCallSid() (value string) {
	if c == nil {
		return
	}
	return c.CallSid
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetSipHeaderDynamicVariables() (value map[string]string) {
	if c == nil {
		return
	}
	return c.SipHeaderDynamicVariables
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetDirection(direction TelephonyDirection) {
	c.Direction = direction
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetPhoneNumberId(phoneNumberId string) {
	c.PhoneNumberId = phoneNumberId
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetAgentNumber(agentNumber string) {
	c.AgentNumber = agentNumber
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetExternalNumber(externalNumber string) {
	c.ExternalNumber = externalNumber
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetCallResult(callResult *CallResult) {
	c.CallResult = callResult
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetCallId(callId *string) {
	c.CallId = callId
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetCallSid(callSid string) {
	c.CallSid = callSid
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) SetSipHeaderDynamicVariables(sipHeaderDynamicVariables map[string]string) {
	c.SipHeaderDynamicVariables = sipHeaderDynamicVariables
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) Type() string {
	return c.type_
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistorySipTrunkingPhoneCallModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistorySipTrunkingPhoneCallModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "sip_trunking" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "sip_trunking", *literals.Type)
	}
	c.type_ = "sip_trunking"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistorySipTrunkingPhoneCallModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistorySipTrunkingPhoneCallModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "sip_trunking",
	}
	return json.Marshal(marshaler)
}

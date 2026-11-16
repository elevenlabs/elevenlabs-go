package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTwilioPhoneCallModel struct {
	Direction       TelephonyDirection `json:"direction" xml:"direction" required:"true"`
	PhoneNumberId   string             `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	AgentNumber     string             `json:"agent_number" xml:"agent_number" required:"true"`
	ExternalNumber  string             `json:"external_number" xml:"external_number" required:"true"`
	CallResult      *CallResult        `json:"call_result,omitempty" xml:"call_result,omitempty"`
	type_           string
	StreamSid       string `json:"stream_sid" xml:"stream_sid" required:"true"`
	CallSid         string `json:"call_sid" xml:"call_sid" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetDirection() (value TelephonyDirection) {
	if c == nil {
		return
	}
	return c.Direction
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetPhoneNumberId() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumberId
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetAgentNumber() (value string) {
	if c == nil {
		return
	}
	return c.AgentNumber
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetExternalNumber() (value string) {
	if c == nil {
		return
	}
	return c.ExternalNumber
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetCallResult() (value *CallResult) {
	if c == nil {
		return
	}
	return c.CallResult
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetStreamSid() (value string) {
	if c == nil {
		return
	}
	return c.StreamSid
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetCallSid() (value string) {
	if c == nil {
		return
	}
	return c.CallSid
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetDirection(direction TelephonyDirection) {
	c.Direction = direction
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetPhoneNumberId(phoneNumberId string) {
	c.PhoneNumberId = phoneNumberId
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetAgentNumber(agentNumber string) {
	c.AgentNumber = agentNumber
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetExternalNumber(externalNumber string) {
	c.ExternalNumber = externalNumber
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetCallResult(callResult *CallResult) {
	c.CallResult = callResult
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetStreamSid(streamSid string) {
	c.StreamSid = streamSid
}

func (c *ConversationHistoryTwilioPhoneCallModel) SetCallSid(callSid string) {
	c.CallSid = callSid
}

func (c *ConversationHistoryTwilioPhoneCallModel) Type() string {
	return c.type_
}

func (c *ConversationHistoryTwilioPhoneCallModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTwilioPhoneCallModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTwilioPhoneCallModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTwilioPhoneCallModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "twilio" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "twilio", *literals.Type)
	}
	c.type_ = "twilio"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTwilioPhoneCallModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTwilioPhoneCallModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "twilio",
	}
	return json.Marshal(marshaler)
}

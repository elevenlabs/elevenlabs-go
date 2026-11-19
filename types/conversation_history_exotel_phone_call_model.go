package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryExotelPhoneCallModel struct {
	Direction       TelephonyDirection `json:"direction" xml:"direction" required:"true"`
	PhoneNumberId   string             `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	AgentNumber     string             `json:"agent_number" xml:"agent_number" required:"true"`
	ExternalNumber  string             `json:"external_number" xml:"external_number" required:"true"`
	CallResult      *CallResult        `json:"call_result,omitempty" xml:"call_result,omitempty"`
	StreamSid       string             `json:"stream_sid" xml:"stream_sid" required:"true"`
	CallSid         string             `json:"call_sid" xml:"call_sid" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryExotelPhoneCallModel) GetDirection() (value TelephonyDirection) {
	if c == nil {
		return
	}
	return c.Direction
}

func (c *ConversationHistoryExotelPhoneCallModel) GetPhoneNumberId() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumberId
}

func (c *ConversationHistoryExotelPhoneCallModel) GetAgentNumber() (value string) {
	if c == nil {
		return
	}
	return c.AgentNumber
}

func (c *ConversationHistoryExotelPhoneCallModel) GetExternalNumber() (value string) {
	if c == nil {
		return
	}
	return c.ExternalNumber
}

func (c *ConversationHistoryExotelPhoneCallModel) GetCallResult() (value *CallResult) {
	if c == nil {
		return
	}
	return c.CallResult
}

func (c *ConversationHistoryExotelPhoneCallModel) GetStreamSid() (value string) {
	if c == nil {
		return
	}
	return c.StreamSid
}

func (c *ConversationHistoryExotelPhoneCallModel) GetCallSid() (value string) {
	if c == nil {
		return
	}
	return c.CallSid
}

func (c *ConversationHistoryExotelPhoneCallModel) SetDirection(direction TelephonyDirection) {
	c.Direction = direction
}

func (c *ConversationHistoryExotelPhoneCallModel) SetPhoneNumberId(phoneNumberId string) {
	c.PhoneNumberId = phoneNumberId
}

func (c *ConversationHistoryExotelPhoneCallModel) SetAgentNumber(agentNumber string) {
	c.AgentNumber = agentNumber
}

func (c *ConversationHistoryExotelPhoneCallModel) SetExternalNumber(externalNumber string) {
	c.ExternalNumber = externalNumber
}

func (c *ConversationHistoryExotelPhoneCallModel) SetCallResult(callResult *CallResult) {
	c.CallResult = callResult
}

func (c *ConversationHistoryExotelPhoneCallModel) SetStreamSid(streamSid string) {
	c.StreamSid = streamSid
}

func (c *ConversationHistoryExotelPhoneCallModel) SetCallSid(callSid string) {
	c.CallSid = callSid
}

func (c *ConversationHistoryExotelPhoneCallModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryExotelPhoneCallModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryExotelPhoneCallModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryExotelPhoneCallModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryExotelPhoneCallModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryExotelPhoneCallModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

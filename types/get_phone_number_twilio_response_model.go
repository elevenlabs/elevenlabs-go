package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPhoneNumberTwilioResponseModel struct {
	// Phone number
	PhoneNumber string `json:"phone_number" xml:"phone_number" required:"true"`
	// Label for the phone number
	Label string `json:"label" xml:"label" required:"true"`
	// This field is deprecated and will be removed in the future. Whether this phone number supports inbound calls
	SupportsInbound *bool `json:"supports_inbound,omitempty" xml:"supports_inbound,omitempty"`
	// This field is deprecated and will be removed in the future. Whether this phone number supports outbound calls
	SupportsOutbound *bool `json:"supports_outbound,omitempty" xml:"supports_outbound,omitempty"`
	// The ID of the phone number
	PhoneNumberId string `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	// The agent that is assigned to the phone number
	AssignedAgent   *PhoneNumberAgentInfo `json:"assigned_agent,omitempty" xml:"assigned_agent,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetPhoneNumberTwilioResponseModel) GetPhoneNumber() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumber
}

func (g *GetPhoneNumberTwilioResponseModel) GetLabel() (value string) {
	if g == nil {
		return
	}
	return g.Label
}

func (g *GetPhoneNumberTwilioResponseModel) GetSupportsInbound() (value *bool) {
	if g == nil {
		return
	}
	return g.SupportsInbound
}

func (g *GetPhoneNumberTwilioResponseModel) GetSupportsOutbound() (value *bool) {
	if g == nil {
		return
	}
	return g.SupportsOutbound
}

func (g *GetPhoneNumberTwilioResponseModel) GetPhoneNumberId() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumberId
}

func (g *GetPhoneNumberTwilioResponseModel) GetAssignedAgent() (value *PhoneNumberAgentInfo) {
	if g == nil {
		return
	}
	return g.AssignedAgent
}

func (g *GetPhoneNumberTwilioResponseModel) SetPhoneNumber(phoneNumber string) {
	g.PhoneNumber = phoneNumber
}

func (g *GetPhoneNumberTwilioResponseModel) SetLabel(label string) {
	g.Label = label
}

func (g *GetPhoneNumberTwilioResponseModel) SetSupportsInbound(supportsInbound *bool) {
	g.SupportsInbound = supportsInbound
}

func (g *GetPhoneNumberTwilioResponseModel) SetSupportsOutbound(supportsOutbound *bool) {
	g.SupportsOutbound = supportsOutbound
}

func (g *GetPhoneNumberTwilioResponseModel) SetPhoneNumberId(phoneNumberId string) {
	g.PhoneNumberId = phoneNumberId
}

func (g *GetPhoneNumberTwilioResponseModel) SetAssignedAgent(assignedAgent *PhoneNumberAgentInfo) {
	g.AssignedAgent = assignedAgent
}

func (g *GetPhoneNumberTwilioResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPhoneNumberTwilioResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPhoneNumberTwilioResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPhoneNumberTwilioResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "provider")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetPhoneNumberTwilioResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPhoneNumberTwilioResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}

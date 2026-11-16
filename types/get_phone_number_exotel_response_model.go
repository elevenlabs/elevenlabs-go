package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPhoneNumberExotelResponseModel struct {
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

func (g *GetPhoneNumberExotelResponseModel) GetPhoneNumber() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumber
}

func (g *GetPhoneNumberExotelResponseModel) GetLabel() (value string) {
	if g == nil {
		return
	}
	return g.Label
}

func (g *GetPhoneNumberExotelResponseModel) GetSupportsInbound() (value *bool) {
	if g == nil {
		return
	}
	return g.SupportsInbound
}

func (g *GetPhoneNumberExotelResponseModel) GetSupportsOutbound() (value *bool) {
	if g == nil {
		return
	}
	return g.SupportsOutbound
}

func (g *GetPhoneNumberExotelResponseModel) GetPhoneNumberId() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumberId
}

func (g *GetPhoneNumberExotelResponseModel) GetAssignedAgent() (value *PhoneNumberAgentInfo) {
	if g == nil {
		return
	}
	return g.AssignedAgent
}

func (g *GetPhoneNumberExotelResponseModel) SetPhoneNumber(phoneNumber string) {
	g.PhoneNumber = phoneNumber
}

func (g *GetPhoneNumberExotelResponseModel) SetLabel(label string) {
	g.Label = label
}

func (g *GetPhoneNumberExotelResponseModel) SetSupportsInbound(supportsInbound *bool) {
	g.SupportsInbound = supportsInbound
}

func (g *GetPhoneNumberExotelResponseModel) SetSupportsOutbound(supportsOutbound *bool) {
	g.SupportsOutbound = supportsOutbound
}

func (g *GetPhoneNumberExotelResponseModel) SetPhoneNumberId(phoneNumberId string) {
	g.PhoneNumberId = phoneNumberId
}

func (g *GetPhoneNumberExotelResponseModel) SetAssignedAgent(assignedAgent *PhoneNumberAgentInfo) {
	g.AssignedAgent = assignedAgent
}

func (g *GetPhoneNumberExotelResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPhoneNumberExotelResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPhoneNumberExotelResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPhoneNumberExotelResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetPhoneNumberExotelResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPhoneNumberExotelResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}

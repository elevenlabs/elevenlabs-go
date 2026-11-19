package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPhoneNumberSipTrunkResponseModel struct {
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
	AssignedAgent  *PhoneNumberAgentInfo                              `json:"assigned_agent,omitempty" xml:"assigned_agent,omitempty"`
	ProviderConfig *GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"provider_config,omitempty" xml:"provider_config,omitempty"`
	// Configuration of the Outbound SIP trunk - if configured.
	OutboundTrunk *GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"outbound_trunk,omitempty" xml:"outbound_trunk,omitempty"`
	// Configuration of the Inbound SIP trunk - if configured.
	InboundTrunk *GetPhoneNumberInboundSipTrunkConfigResponseModel `json:"inbound_trunk,omitempty" xml:"inbound_trunk,omitempty"`
	LivekitStack LivekitStackType                                  `json:"livekit_stack" xml:"livekit_stack" required:"true"`
	// Whether to store SIP messages for this phone number.
	StoreSipMessages *bool `json:"store_sip_messages,omitempty" xml:"store_sip_messages,omitempty"`
	extraProperties  map[string]interface{}
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetPhoneNumber() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumber
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetLabel() (value string) {
	if g == nil {
		return
	}
	return g.Label
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetSupportsInbound() (value *bool) {
	if g == nil {
		return
	}
	return g.SupportsInbound
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetSupportsOutbound() (value *bool) {
	if g == nil {
		return
	}
	return g.SupportsOutbound
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetPhoneNumberId() (value string) {
	if g == nil {
		return
	}
	return g.PhoneNumberId
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetAssignedAgent() (value *PhoneNumberAgentInfo) {
	if g == nil {
		return
	}
	return g.AssignedAgent
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetProviderConfig() (value *GetPhoneNumberOutboundSipTrunkConfigResponseModel) {
	if g == nil {
		return
	}
	return g.ProviderConfig
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetOutboundTrunk() (value *GetPhoneNumberOutboundSipTrunkConfigResponseModel) {
	if g == nil {
		return
	}
	return g.OutboundTrunk
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetInboundTrunk() (value *GetPhoneNumberInboundSipTrunkConfigResponseModel) {
	if g == nil {
		return
	}
	return g.InboundTrunk
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetLivekitStack() (value LivekitStackType) {
	if g == nil {
		return
	}
	return g.LivekitStack
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetStoreSipMessages() (value *bool) {
	if g == nil {
		return
	}
	return g.StoreSipMessages
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetPhoneNumber(phoneNumber string) {
	g.PhoneNumber = phoneNumber
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetLabel(label string) {
	g.Label = label
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetSupportsInbound(supportsInbound *bool) {
	g.SupportsInbound = supportsInbound
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetSupportsOutbound(supportsOutbound *bool) {
	g.SupportsOutbound = supportsOutbound
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetPhoneNumberId(phoneNumberId string) {
	g.PhoneNumberId = phoneNumberId
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetAssignedAgent(assignedAgent *PhoneNumberAgentInfo) {
	g.AssignedAgent = assignedAgent
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetProviderConfig(providerConfig *GetPhoneNumberOutboundSipTrunkConfigResponseModel) {
	g.ProviderConfig = providerConfig
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetOutboundTrunk(outboundTrunk *GetPhoneNumberOutboundSipTrunkConfigResponseModel) {
	g.OutboundTrunk = outboundTrunk
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetInboundTrunk(inboundTrunk *GetPhoneNumberInboundSipTrunkConfigResponseModel) {
	g.InboundTrunk = inboundTrunk
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetLivekitStack(livekitStack LivekitStackType) {
	g.LivekitStack = livekitStack
}

func (g *GetPhoneNumberSipTrunkResponseModel) SetStoreSipMessages(storeSipMessages *bool) {
	g.StoreSipMessages = storeSipMessages
}

func (g *GetPhoneNumberSipTrunkResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPhoneNumberSipTrunkResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPhoneNumberSipTrunkResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPhoneNumberSipTrunkResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetPhoneNumberSipTrunkResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPhoneNumberSipTrunkResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}

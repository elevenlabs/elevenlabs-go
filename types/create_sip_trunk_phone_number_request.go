package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateSipTrunkPhoneNumberRequest struct {
	// Phone number
	PhoneNumber string `json:"phone_number" xml:"phone_number" required:"true"`
	// Label for the phone number
	Label string `json:"label" xml:"label" required:"true"`
	// This field is deprecated and will be removed in the future. Whether this phone number supports inbound calls
	SupportsInbound *bool `json:"supports_inbound,omitempty" xml:"supports_inbound,omitempty"`
	// This field is deprecated and will be removed in the future. Whether this phone number supports outbound calls
	SupportsOutbound *bool   `json:"supports_outbound,omitempty" xml:"supports_outbound,omitempty"`
	Provider         *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// Agent ID to assign the phone number to
	AgentId             *string                             `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	InboundTrunkConfig  *InboundSipTrunkConfigRequestModel  `json:"inbound_trunk_config,omitempty" xml:"inbound_trunk_config,omitempty"`
	OutboundTrunkConfig *OutboundSipTrunkConfigRequestModel `json:"outbound_trunk_config,omitempty" xml:"outbound_trunk_config,omitempty"`
	extraProperties     map[string]interface{}
}

func (c *CreateSipTrunkPhoneNumberRequest) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

func (c *CreateSipTrunkPhoneNumberRequest) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CreateSipTrunkPhoneNumberRequest) GetSupportsInbound() (value *bool) {
	if c == nil {
		return
	}
	return c.SupportsInbound
}

func (c *CreateSipTrunkPhoneNumberRequest) GetSupportsOutbound() (value *bool) {
	if c == nil {
		return
	}
	return c.SupportsOutbound
}

func (c *CreateSipTrunkPhoneNumberRequest) GetProvider() (value *string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateSipTrunkPhoneNumberRequest) GetAgentId() (value *string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CreateSipTrunkPhoneNumberRequest) GetInboundTrunkConfig() (value *InboundSipTrunkConfigRequestModel) {
	if c == nil {
		return
	}
	return c.InboundTrunkConfig
}

func (c *CreateSipTrunkPhoneNumberRequest) GetOutboundTrunkConfig() (value *OutboundSipTrunkConfigRequestModel) {
	if c == nil {
		return
	}
	return c.OutboundTrunkConfig
}

func (c *CreateSipTrunkPhoneNumberRequest) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

func (c *CreateSipTrunkPhoneNumberRequest) SetLabel(label string) {
	c.Label = label
}

func (c *CreateSipTrunkPhoneNumberRequest) SetSupportsInbound(supportsInbound *bool) {
	c.SupportsInbound = supportsInbound
}

func (c *CreateSipTrunkPhoneNumberRequest) SetSupportsOutbound(supportsOutbound *bool) {
	c.SupportsOutbound = supportsOutbound
}

func (c *CreateSipTrunkPhoneNumberRequest) SetProvider(provider *string) {
	c.Provider = provider
}

func (c *CreateSipTrunkPhoneNumberRequest) SetAgentId(agentId *string) {
	c.AgentId = agentId
}

func (c *CreateSipTrunkPhoneNumberRequest) SetInboundTrunkConfig(inboundTrunkConfig *InboundSipTrunkConfigRequestModel) {
	c.InboundTrunkConfig = inboundTrunkConfig
}

func (c *CreateSipTrunkPhoneNumberRequest) SetOutboundTrunkConfig(outboundTrunkConfig *OutboundSipTrunkConfigRequestModel) {
	c.OutboundTrunkConfig = outboundTrunkConfig
}

func (c *CreateSipTrunkPhoneNumberRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateSipTrunkPhoneNumberRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateSipTrunkPhoneNumberRequest to string"
	}
	return string(jsonData)
}

func (c *CreateSipTrunkPhoneNumberRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateSipTrunkPhoneNumberRequest) MarshalJSON() ([]byte, error) {
	type embed CreateSipTrunkPhoneNumberRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

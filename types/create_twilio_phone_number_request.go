package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateTwilioPhoneNumberRequest struct {
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
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// Twilio Account SID (starts with `AC`) or API Key SID (starts with `SK`)
	Sid string `json:"sid" xml:"sid" required:"true"`
	// Secret paired with `sid`: the Account Auth Token for an Account SID, or the API Key Secret for an API Key SID
	Token string `json:"token" xml:"token" required:"true"`
	// Twilio Account Auth Token, required for API Key imports to validate inbound webhook signatures
	AccountAuthToken *string `json:"account_auth_token,omitempty" xml:"account_auth_token,omitempty"`
	// Twilio Additional Region Configuration
	RegionConfig *RegionConfigRequest `json:"region_config,omitempty" xml:"region_config,omitempty"`
	// Route inbound SMS to ElevenLabs. On by default; set to false to skip SMS configuration for numbers that don't support it.
	EnableSms       *bool `json:"enable_sms,omitempty" xml:"enable_sms,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateTwilioPhoneNumberRequest) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

func (c *CreateTwilioPhoneNumberRequest) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CreateTwilioPhoneNumberRequest) GetSupportsInbound() (value *bool) {
	if c == nil {
		return
	}
	return c.SupportsInbound
}

func (c *CreateTwilioPhoneNumberRequest) GetSupportsOutbound() (value *bool) {
	if c == nil {
		return
	}
	return c.SupportsOutbound
}

func (c *CreateTwilioPhoneNumberRequest) GetProvider() (value *string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateTwilioPhoneNumberRequest) GetAgentId() (value *string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CreateTwilioPhoneNumberRequest) GetSid() (value string) {
	if c == nil {
		return
	}
	return c.Sid
}

func (c *CreateTwilioPhoneNumberRequest) GetToken() (value string) {
	if c == nil {
		return
	}
	return c.Token
}

func (c *CreateTwilioPhoneNumberRequest) GetAccountAuthToken() (value *string) {
	if c == nil {
		return
	}
	return c.AccountAuthToken
}

func (c *CreateTwilioPhoneNumberRequest) GetRegionConfig() (value *RegionConfigRequest) {
	if c == nil {
		return
	}
	return c.RegionConfig
}

func (c *CreateTwilioPhoneNumberRequest) GetEnableSms() (value *bool) {
	if c == nil {
		return
	}
	return c.EnableSms
}

func (c *CreateTwilioPhoneNumberRequest) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

func (c *CreateTwilioPhoneNumberRequest) SetLabel(label string) {
	c.Label = label
}

func (c *CreateTwilioPhoneNumberRequest) SetSupportsInbound(supportsInbound *bool) {
	c.SupportsInbound = supportsInbound
}

func (c *CreateTwilioPhoneNumberRequest) SetSupportsOutbound(supportsOutbound *bool) {
	c.SupportsOutbound = supportsOutbound
}

func (c *CreateTwilioPhoneNumberRequest) SetProvider(provider *string) {
	c.Provider = provider
}

func (c *CreateTwilioPhoneNumberRequest) SetAgentId(agentId *string) {
	c.AgentId = agentId
}

func (c *CreateTwilioPhoneNumberRequest) SetSid(sid string) {
	c.Sid = sid
}

func (c *CreateTwilioPhoneNumberRequest) SetToken(token string) {
	c.Token = token
}

func (c *CreateTwilioPhoneNumberRequest) SetAccountAuthToken(accountAuthToken *string) {
	c.AccountAuthToken = accountAuthToken
}

func (c *CreateTwilioPhoneNumberRequest) SetRegionConfig(regionConfig *RegionConfigRequest) {
	c.RegionConfig = regionConfig
}

func (c *CreateTwilioPhoneNumberRequest) SetEnableSms(enableSms *bool) {
	c.EnableSms = enableSms
}

func (c *CreateTwilioPhoneNumberRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateTwilioPhoneNumberRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTwilioPhoneNumberRequest to string"
	}
	return string(jsonData)
}

func (c *CreateTwilioPhoneNumberRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreateTwilioPhoneNumberRequest) MarshalJSON() ([]byte, error) {
	type embed CreateTwilioPhoneNumberRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

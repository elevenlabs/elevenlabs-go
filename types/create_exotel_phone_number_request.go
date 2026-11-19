package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateExotelPhoneNumberRequest struct {
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
	// Exotel Account SID
	AccountSid string `json:"account_sid" xml:"account_sid" required:"true"`
	// Exotel API Key
	ApiKey string `json:"api_key" xml:"api_key" required:"true"`
	// Exotel API Token
	ApiToken     string             `json:"api_token" xml:"api_token" required:"true"`
	ApiSubdomain ExotelApiSubdomain `json:"api_subdomain" xml:"api_subdomain" required:"true"`
	// Exotel applet identifier used in Calls/connect
	AppId string `json:"app_id" xml:"app_id" required:"true"`
	// Optional full applet URL override. Defaults to Exotel start_voice URL derived from account SID and app ID.
	AppletUrl       *string `json:"applet_url,omitempty" xml:"applet_url,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateExotelPhoneNumberRequest) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

func (c *CreateExotelPhoneNumberRequest) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *CreateExotelPhoneNumberRequest) GetSupportsInbound() (value *bool) {
	if c == nil {
		return
	}
	return c.SupportsInbound
}

func (c *CreateExotelPhoneNumberRequest) GetSupportsOutbound() (value *bool) {
	if c == nil {
		return
	}
	return c.SupportsOutbound
}

func (c *CreateExotelPhoneNumberRequest) GetProvider() (value *string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CreateExotelPhoneNumberRequest) GetAgentId() (value *string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CreateExotelPhoneNumberRequest) GetAccountSid() (value string) {
	if c == nil {
		return
	}
	return c.AccountSid
}

func (c *CreateExotelPhoneNumberRequest) GetApiKey() (value string) {
	if c == nil {
		return
	}
	return c.ApiKey
}

func (c *CreateExotelPhoneNumberRequest) GetApiToken() (value string) {
	if c == nil {
		return
	}
	return c.ApiToken
}

func (c *CreateExotelPhoneNumberRequest) GetApiSubdomain() (value ExotelApiSubdomain) {
	if c == nil {
		return
	}
	return c.ApiSubdomain
}

func (c *CreateExotelPhoneNumberRequest) GetAppId() (value string) {
	if c == nil {
		return
	}
	return c.AppId
}

func (c *CreateExotelPhoneNumberRequest) GetAppletUrl() (value *string) {
	if c == nil {
		return
	}
	return c.AppletUrl
}

func (c *CreateExotelPhoneNumberRequest) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = phoneNumber
}

func (c *CreateExotelPhoneNumberRequest) SetLabel(label string) {
	c.Label = label
}

func (c *CreateExotelPhoneNumberRequest) SetSupportsInbound(supportsInbound *bool) {
	c.SupportsInbound = supportsInbound
}

func (c *CreateExotelPhoneNumberRequest) SetSupportsOutbound(supportsOutbound *bool) {
	c.SupportsOutbound = supportsOutbound
}

func (c *CreateExotelPhoneNumberRequest) SetProvider(provider *string) {
	c.Provider = provider
}

func (c *CreateExotelPhoneNumberRequest) SetAgentId(agentId *string) {
	c.AgentId = agentId
}

func (c *CreateExotelPhoneNumberRequest) SetAccountSid(accountSid string) {
	c.AccountSid = accountSid
}

func (c *CreateExotelPhoneNumberRequest) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
}

func (c *CreateExotelPhoneNumberRequest) SetApiToken(apiToken string) {
	c.ApiToken = apiToken
}

func (c *CreateExotelPhoneNumberRequest) SetApiSubdomain(apiSubdomain ExotelApiSubdomain) {
	c.ApiSubdomain = apiSubdomain
}

func (c *CreateExotelPhoneNumberRequest) SetAppId(appId string) {
	c.AppId = appId
}

func (c *CreateExotelPhoneNumberRequest) SetAppletUrl(appletUrl *string) {
	c.AppletUrl = appletUrl
}

func (c *CreateExotelPhoneNumberRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateExotelPhoneNumberRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateExotelPhoneNumberRequest to string"
	}
	return string(jsonData)
}

func (c *CreateExotelPhoneNumberRequest) UnmarshalJSON(data []byte) error {
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

func (c *CreateExotelPhoneNumberRequest) MarshalJSON() ([]byte, error) {
	type embed CreateExotelPhoneNumberRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

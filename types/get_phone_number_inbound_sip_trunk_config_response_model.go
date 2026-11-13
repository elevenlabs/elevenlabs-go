package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPhoneNumberInboundSipTrunkConfigResponseModel struct {
	// List of IP addresses that are allowed to use the trunk. Each item in the list can be an individual IP address or a Classless Inter-Domain Routing notation representing a CIDR block.
	AllowedAddresses []string `json:"allowed_addresses" xml:"allowed_addresses" required:"true"`
	// List of phone numbers that are allowed to use the trunk.
	AllowedNumbers  []string               `json:"allowed_numbers" xml:"allowed_numbers" required:"true"`
	MediaEncryption SipMediaEncryptionEnum `json:"media_encryption" xml:"media_encryption" required:"true"`
	// Whether authentication credentials are configured
	HasAuthCredentials bool `json:"has_auth_credentials" xml:"has_auth_credentials" required:"true"`
	// SIP trunk username (if available)
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// Domains of remote SIP servers used to validate TLS certificates.
	RemoteDomains []string `json:"remote_domains,omitempty" xml:"remote_domains,omitempty"`
	// Map of dynamic variable name to header name for attributes_to_headers
	AttributesToHeaders map[string]string `json:"attributes_to_headers,omitempty"`
	extraProperties     map[string]interface{}
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetAllowedAddresses() (value []string) {
	if g == nil {
		return
	}
	return g.AllowedAddresses
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetAllowedNumbers() (value []string) {
	if g == nil {
		return
	}
	return g.AllowedNumbers
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetMediaEncryption() (value SipMediaEncryptionEnum) {
	if g == nil {
		return
	}
	return g.MediaEncryption
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetHasAuthCredentials() (value bool) {
	if g == nil {
		return
	}
	return g.HasAuthCredentials
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetUsername() (value *string) {
	if g == nil {
		return
	}
	return g.Username
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetRemoteDomains() (value []string) {
	if g == nil {
		return
	}
	return g.RemoteDomains
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetAttributesToHeaders() (value map[string]string) {
	if g == nil {
		return
	}
	return g.AttributesToHeaders
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetAllowedAddresses(allowedAddresses []string) {
	g.AllowedAddresses = allowedAddresses
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetAllowedNumbers(allowedNumbers []string) {
	g.AllowedNumbers = allowedNumbers
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetMediaEncryption(mediaEncryption SipMediaEncryptionEnum) {
	g.MediaEncryption = mediaEncryption
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetHasAuthCredentials(hasAuthCredentials bool) {
	g.HasAuthCredentials = hasAuthCredentials
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetUsername(username *string) {
	g.Username = username
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetRemoteDomains(remoteDomains []string) {
	g.RemoteDomains = remoteDomains
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) SetAttributesToHeaders(attributesToHeaders map[string]string) {
	g.AttributesToHeaders = attributesToHeaders
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPhoneNumberInboundSipTrunkConfigResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPhoneNumberInboundSipTrunkConfigResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetPhoneNumberInboundSipTrunkConfigResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPhoneNumberInboundSipTrunkConfigResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}

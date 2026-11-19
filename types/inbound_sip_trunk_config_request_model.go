package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InboundSipTrunkConfigRequestModel struct {
	// List of IP addresses that are allowed to use the trunk. Each item in the list can be an individual IP address or a Classless Inter-Domain Routing notation representing a CIDR block.
	AllowedAddresses []string `json:"allowed_addresses,omitempty" xml:"allowed_addresses,omitempty"`
	// List of phone numbers that are allowed to use the trunk.
	AllowedNumbers  []string                `json:"allowed_numbers,omitempty" xml:"allowed_numbers,omitempty"`
	MediaEncryption *SipMediaEncryptionEnum `json:"media_encryption,omitempty" xml:"media_encryption,omitempty"`
	// Optional digest authentication credentials (username/password).
	Credentials *SipTrunkCredentialsRequestModel `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// Domains of remote SIP servers used to validate TLS certificates.
	RemoteDomains []string `json:"remote_domains,omitempty" xml:"remote_domains,omitempty"`
	// Map of dynamic variable name to header name for attributes_to_headers
	AttributesToHeaders map[string]string `json:"attributes_to_headers,omitempty"`
	extraProperties     map[string]interface{}
}

func (i *InboundSipTrunkConfigRequestModel) GetAllowedAddresses() (value []string) {
	if i == nil {
		return
	}
	return i.AllowedAddresses
}

func (i *InboundSipTrunkConfigRequestModel) GetAllowedNumbers() (value []string) {
	if i == nil {
		return
	}
	return i.AllowedNumbers
}

func (i *InboundSipTrunkConfigRequestModel) GetMediaEncryption() (value *SipMediaEncryptionEnum) {
	if i == nil {
		return
	}
	return i.MediaEncryption
}

func (i *InboundSipTrunkConfigRequestModel) GetCredentials() (value *SipTrunkCredentialsRequestModel) {
	if i == nil {
		return
	}
	return i.Credentials
}

func (i *InboundSipTrunkConfigRequestModel) GetRemoteDomains() (value []string) {
	if i == nil {
		return
	}
	return i.RemoteDomains
}

func (i *InboundSipTrunkConfigRequestModel) GetAttributesToHeaders() (value map[string]string) {
	if i == nil {
		return
	}
	return i.AttributesToHeaders
}

func (i *InboundSipTrunkConfigRequestModel) SetAllowedAddresses(allowedAddresses []string) {
	i.AllowedAddresses = allowedAddresses
}

func (i *InboundSipTrunkConfigRequestModel) SetAllowedNumbers(allowedNumbers []string) {
	i.AllowedNumbers = allowedNumbers
}

func (i *InboundSipTrunkConfigRequestModel) SetMediaEncryption(mediaEncryption *SipMediaEncryptionEnum) {
	i.MediaEncryption = mediaEncryption
}

func (i *InboundSipTrunkConfigRequestModel) SetCredentials(credentials *SipTrunkCredentialsRequestModel) {
	i.Credentials = credentials
}

func (i *InboundSipTrunkConfigRequestModel) SetRemoteDomains(remoteDomains []string) {
	i.RemoteDomains = remoteDomains
}

func (i *InboundSipTrunkConfigRequestModel) SetAttributesToHeaders(attributesToHeaders map[string]string) {
	i.AttributesToHeaders = attributesToHeaders
}

func (i *InboundSipTrunkConfigRequestModel) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InboundSipTrunkConfigRequestModel) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboundSipTrunkConfigRequestModel to string"
	}
	return string(jsonData)
}

func (i *InboundSipTrunkConfigRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InboundSipTrunkConfigRequestModel) MarshalJSON() ([]byte, error) {
	type embed InboundSipTrunkConfigRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}

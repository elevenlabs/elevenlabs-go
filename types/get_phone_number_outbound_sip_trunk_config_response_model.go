package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// SIP Trunk configuration details for a phone number
type GetPhoneNumberOutboundSipTrunkConfigResponseModel struct {
	// Hostname or IP the SIP INVITE is sent to
	Address         string                 `json:"address" xml:"address" required:"true"`
	Transport       SipTrunkTransportEnum  `json:"transport" xml:"transport" required:"true"`
	MediaEncryption SipMediaEncryptionEnum `json:"media_encryption" xml:"media_encryption" required:"true"`
	// SIP headers for INVITE request
	Headers map[string]string `json:"headers,omitempty"`
	// Map of dynamic variable name to header name for attributes_to_headers
	AttributesToHeaders map[string]string `json:"attributes_to_headers,omitempty"`
	// Whether authentication credentials are configured
	HasAuthCredentials bool `json:"has_auth_credentials" xml:"has_auth_credentials" required:"true"`
	// SIP trunk username (if available)
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// Whether a LiveKit SIP outbound trunk is configured
	HasOutboundTrunk *bool `json:"has_outbound_trunk,omitempty" xml:"has_outbound_trunk,omitempty"`
	// Media codecs that are offered in the SDP for outbound calls. If empty, all supported codecs are offered.
	EnabledCodecs   []MediaCodec `json:"enabled_codecs,omitempty" xml:"enabled_codecs,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetAddress() (value string) {
	if g == nil {
		return
	}
	return g.Address
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetTransport() (value SipTrunkTransportEnum) {
	if g == nil {
		return
	}
	return g.Transport
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetMediaEncryption() (value SipMediaEncryptionEnum) {
	if g == nil {
		return
	}
	return g.MediaEncryption
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetHeaders() (value map[string]string) {
	if g == nil {
		return
	}
	return g.Headers
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetAttributesToHeaders() (value map[string]string) {
	if g == nil {
		return
	}
	return g.AttributesToHeaders
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetHasAuthCredentials() (value bool) {
	if g == nil {
		return
	}
	return g.HasAuthCredentials
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetUsername() (value *string) {
	if g == nil {
		return
	}
	return g.Username
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetHasOutboundTrunk() (value *bool) {
	if g == nil {
		return
	}
	return g.HasOutboundTrunk
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetEnabledCodecs() (value []MediaCodec) {
	if g == nil {
		return
	}
	return g.EnabledCodecs
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetAddress(address string) {
	g.Address = address
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetTransport(transport SipTrunkTransportEnum) {
	g.Transport = transport
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetMediaEncryption(mediaEncryption SipMediaEncryptionEnum) {
	g.MediaEncryption = mediaEncryption
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetHeaders(headers map[string]string) {
	g.Headers = headers
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetAttributesToHeaders(attributesToHeaders map[string]string) {
	g.AttributesToHeaders = attributesToHeaders
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetHasAuthCredentials(hasAuthCredentials bool) {
	g.HasAuthCredentials = hasAuthCredentials
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetUsername(username *string) {
	g.Username = username
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetHasOutboundTrunk(hasOutboundTrunk *bool) {
	g.HasOutboundTrunk = hasOutboundTrunk
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) SetEnabledCodecs(enabledCodecs []MediaCodec) {
	g.EnabledCodecs = enabledCodecs
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPhoneNumberOutboundSipTrunkConfigResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPhoneNumberOutboundSipTrunkConfigResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetPhoneNumberOutboundSipTrunkConfigResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPhoneNumberOutboundSipTrunkConfigResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}

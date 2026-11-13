package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OutboundSipTrunkConfigRequestModel struct {
	// Hostname or IP the SIP INVITE is sent to.
	Address         string                  `json:"address" xml:"address" required:"true"`
	Transport       *SipTrunkTransportEnum  `json:"transport,omitempty" xml:"transport,omitempty"`
	MediaEncryption *SipMediaEncryptionEnum `json:"media_encryption,omitempty" xml:"media_encryption,omitempty"`
	// SIP X-* headers for INVITE request. These headers are sent as-is and may help identify this call.
	Headers map[string]string `json:"headers,omitempty"`
	// Map of dynamic variable name to header name for attributes_to_headers
	AttributesToHeaders map[string]string `json:"attributes_to_headers,omitempty"`
	// Optional digest authentication credentials (username/password). If not provided, ACL authentication is assumed.
	Credentials *SipTrunkCredentialsRequestModel `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// Media codecs that should be offered in the SDP for outbound calls. If empty, all supported codecs are offered.
	EnabledCodecs   []MediaCodec `json:"enabled_codecs,omitempty" xml:"enabled_codecs,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OutboundSipTrunkConfigRequestModel) GetAddress() (value string) {
	if o == nil {
		return
	}
	return o.Address
}

func (o *OutboundSipTrunkConfigRequestModel) GetTransport() (value *SipTrunkTransportEnum) {
	if o == nil {
		return
	}
	return o.Transport
}

func (o *OutboundSipTrunkConfigRequestModel) GetMediaEncryption() (value *SipMediaEncryptionEnum) {
	if o == nil {
		return
	}
	return o.MediaEncryption
}

func (o *OutboundSipTrunkConfigRequestModel) GetHeaders() (value map[string]string) {
	if o == nil {
		return
	}
	return o.Headers
}

func (o *OutboundSipTrunkConfigRequestModel) GetAttributesToHeaders() (value map[string]string) {
	if o == nil {
		return
	}
	return o.AttributesToHeaders
}

func (o *OutboundSipTrunkConfigRequestModel) GetCredentials() (value *SipTrunkCredentialsRequestModel) {
	if o == nil {
		return
	}
	return o.Credentials
}

func (o *OutboundSipTrunkConfigRequestModel) GetEnabledCodecs() (value []MediaCodec) {
	if o == nil {
		return
	}
	return o.EnabledCodecs
}

func (o *OutboundSipTrunkConfigRequestModel) SetAddress(address string) {
	o.Address = address
}

func (o *OutboundSipTrunkConfigRequestModel) SetTransport(transport *SipTrunkTransportEnum) {
	o.Transport = transport
}

func (o *OutboundSipTrunkConfigRequestModel) SetMediaEncryption(mediaEncryption *SipMediaEncryptionEnum) {
	o.MediaEncryption = mediaEncryption
}

func (o *OutboundSipTrunkConfigRequestModel) SetHeaders(headers map[string]string) {
	o.Headers = headers
}

func (o *OutboundSipTrunkConfigRequestModel) SetAttributesToHeaders(attributesToHeaders map[string]string) {
	o.AttributesToHeaders = attributesToHeaders
}

func (o *OutboundSipTrunkConfigRequestModel) SetCredentials(credentials *SipTrunkCredentialsRequestModel) {
	o.Credentials = credentials
}

func (o *OutboundSipTrunkConfigRequestModel) SetEnabledCodecs(enabledCodecs []MediaCodec) {
	o.EnabledCodecs = enabledCodecs
}

func (o *OutboundSipTrunkConfigRequestModel) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OutboundSipTrunkConfigRequestModel) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OutboundSipTrunkConfigRequestModel to string"
	}
	return string(jsonData)
}

func (o *OutboundSipTrunkConfigRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OutboundSipTrunkConfigRequestModel) MarshalJSON() ([]byte, error) {
	type embed OutboundSipTrunkConfigRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

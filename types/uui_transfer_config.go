package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// User-to-User Information envelope for SIP REFER transfers (RFC 7433).
//
// Outbound payloads are hex-encoded (the only encoding RFC 7433 defines). The
// protocol discriminator axis lets per-platform formats (Talkdesk, Genesys, ...)
// be expressed by configuration rather than scattered transfer flags. Further
// axes (ASCII encoding, header name, purpose/content parameters) can be added
// here without touching the transfer model.
type UuiTransferConfig struct {
	// UUI payload to send on SIP REFER transfers. Supports inline dynamic variables and is hex-encoded at transfer time.
	Data string `json:"data" xml:"data" required:"true" maxLength:"8192"`
	// Optional one-octet protocol discriminator (two hex digits, e.g. '00'). Required by platforms such as Genesys Cloud, which otherwise strip the first octet of the payload. Leave unset for platforms like Talkdesk that expect a bare hex payload.
	ProtocolDiscriminator *string `json:"protocol_discriminator,omitempty" xml:"protocol_discriminator,omitempty" pattern:"^[0-9A-Fa-f]{2}$"`
	// How to attach protocol_discriminator. 'prefix' prepends the octet to the hex payload (User-to-User=XX<hex>;encoding=hex). 'pd_parameter' sends it as a separate parameter (User-to-User=<hex>;pd=XX;encoding=hex). Ignored when protocol_discriminator is unset.
	ProtocolDiscriminatorMode *UuiTransferConfigProtocolDiscriminatorMode `json:"protocol_discriminator_mode,omitempty" xml:"protocol_discriminator_mode,omitempty"`
	extraProperties           map[string]interface{}
}

func (u *UuiTransferConfig) GetData() (value string) {
	if u == nil {
		return
	}
	return u.Data
}

func (u *UuiTransferConfig) GetProtocolDiscriminator() (value *string) {
	if u == nil {
		return
	}
	return u.ProtocolDiscriminator
}

func (u *UuiTransferConfig) GetProtocolDiscriminatorMode() (value *UuiTransferConfigProtocolDiscriminatorMode) {
	if u == nil {
		return
	}
	return u.ProtocolDiscriminatorMode
}

func (u *UuiTransferConfig) SetData(data string) {
	u.Data = data
}

func (u *UuiTransferConfig) SetProtocolDiscriminator(protocolDiscriminator *string) {
	u.ProtocolDiscriminator = protocolDiscriminator
}

func (u *UuiTransferConfig) SetProtocolDiscriminatorMode(protocolDiscriminatorMode *UuiTransferConfigProtocolDiscriminatorMode) {
	u.ProtocolDiscriminatorMode = protocolDiscriminatorMode
}

func (u *UuiTransferConfig) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UuiTransferConfig) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UuiTransferConfig to string"
	}
	return string(jsonData)
}

func (u *UuiTransferConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UuiTransferConfig) MarshalJSON() ([]byte, error) {
	type embed UuiTransferConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

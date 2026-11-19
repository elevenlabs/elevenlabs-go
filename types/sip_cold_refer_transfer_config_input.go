package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipColdReferTransferConfigInput struct {
	// Custom SIP headers to include when transferring the call. Each header can be either a static value or a dynamic variable reference.
	CustomSipHeaders    []*SipColdReferTransferConfigInputCustomSipHeadersItem `json:"custom_sip_headers,omitempty" xml:"custom_sip_headers,omitempty" maxItems:"20"`
	TransferDestination *SipColdReferTransferConfigInputTransferDestination    `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	// User-to-User Information (RFC 7433) attached to the REFER. Carries call context such as CRM identifiers or escalation reason across the transfer boundary.
	Uui *UuiTransferConfig `json:"uui,omitempty" xml:"uui,omitempty"`
	// Play a ringing tone on the original call leg while the REFER completes. Disable if the receiving SBC/contact center should not hear it; the caller then hears silence until the transfer completes.
	PlayDialtone    *bool `json:"play_dialtone,omitempty" xml:"play_dialtone,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SipColdReferTransferConfigInput) GetCustomSipHeaders() (value []*SipColdReferTransferConfigInputCustomSipHeadersItem) {
	if s == nil {
		return
	}
	return s.CustomSipHeaders
}

func (s *SipColdReferTransferConfigInput) GetTransferDestination() (value *SipColdReferTransferConfigInputTransferDestination) {
	if s == nil {
		return
	}
	return s.TransferDestination
}

func (s *SipColdReferTransferConfigInput) GetUui() (value *UuiTransferConfig) {
	if s == nil {
		return
	}
	return s.Uui
}

func (s *SipColdReferTransferConfigInput) GetPlayDialtone() (value *bool) {
	if s == nil {
		return
	}
	return s.PlayDialtone
}

func (s *SipColdReferTransferConfigInput) SetCustomSipHeaders(customSipHeaders []*SipColdReferTransferConfigInputCustomSipHeadersItem) {
	s.CustomSipHeaders = customSipHeaders
}

func (s *SipColdReferTransferConfigInput) SetTransferDestination(transferDestination *SipColdReferTransferConfigInputTransferDestination) {
	s.TransferDestination = transferDestination
}

func (s *SipColdReferTransferConfigInput) SetUui(uui *UuiTransferConfig) {
	s.Uui = uui
}

func (s *SipColdReferTransferConfigInput) SetPlayDialtone(playDialtone *bool) {
	s.PlayDialtone = playDialtone
}

func (s *SipColdReferTransferConfigInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipColdReferTransferConfigInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipColdReferTransferConfigInput to string"
	}
	return string(jsonData)
}

func (s *SipColdReferTransferConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipColdReferTransferConfigInput) MarshalJSON() ([]byte, error) {
	type embed SipColdReferTransferConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipColdReferTransferConfigOutput struct {
	// Custom SIP headers to include when transferring the call. Each header can be either a static value or a dynamic variable reference.
	CustomSipHeaders    []*SipColdReferTransferConfigOutputCustomSipHeadersItem `json:"custom_sip_headers" xml:"custom_sip_headers" required:"true" maxItems:"20"`
	TransferDestination *SipColdReferTransferConfigOutputTransferDestination    `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	type_               string
	// User-to-User Information (RFC 7433) attached to the REFER. Carries call context such as CRM identifiers or escalation reason across the transfer boundary.
	Uui *UuiTransferConfig `json:"uui" xml:"uui" required:"true"`
	// Play a ringing tone on the original call leg while the REFER completes. Disable if the receiving SBC/contact center should not hear it; the caller then hears silence until the transfer completes.
	PlayDialtone    bool `json:"play_dialtone" xml:"play_dialtone" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SipColdReferTransferConfigOutput) GetCustomSipHeaders() (value []*SipColdReferTransferConfigOutputCustomSipHeadersItem) {
	if s == nil {
		return
	}
	return s.CustomSipHeaders
}

func (s *SipColdReferTransferConfigOutput) GetTransferDestination() (value *SipColdReferTransferConfigOutputTransferDestination) {
	if s == nil {
		return
	}
	return s.TransferDestination
}

func (s *SipColdReferTransferConfigOutput) GetUui() (value *UuiTransferConfig) {
	if s == nil {
		return
	}
	return s.Uui
}

func (s *SipColdReferTransferConfigOutput) GetPlayDialtone() (value bool) {
	if s == nil {
		return
	}
	return s.PlayDialtone
}

func (s *SipColdReferTransferConfigOutput) SetCustomSipHeaders(customSipHeaders []*SipColdReferTransferConfigOutputCustomSipHeadersItem) {
	s.CustomSipHeaders = customSipHeaders
}

func (s *SipColdReferTransferConfigOutput) SetTransferDestination(transferDestination *SipColdReferTransferConfigOutputTransferDestination) {
	s.TransferDestination = transferDestination
}

func (s *SipColdReferTransferConfigOutput) SetUui(uui *UuiTransferConfig) {
	s.Uui = uui
}

func (s *SipColdReferTransferConfigOutput) SetPlayDialtone(playDialtone bool) {
	s.PlayDialtone = playDialtone
}

func (s *SipColdReferTransferConfigOutput) Type() string {
	return s.type_
}

func (s *SipColdReferTransferConfigOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipColdReferTransferConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipColdReferTransferConfigOutput to string"
	}
	return string(jsonData)
}

func (s *SipColdReferTransferConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "sip_cold_refer_transfer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "sip_cold_refer_transfer", *literals.Type)
	}
	s.type_ = "sip_cold_refer_transfer"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipColdReferTransferConfigOutput) MarshalJSON() ([]byte, error) {
	type embed SipColdReferTransferConfigOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "sip_cold_refer_transfer",
	}
	return json.Marshal(marshaler)
}

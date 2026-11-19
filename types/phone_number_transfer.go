package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneNumberTransfer struct {
	// Custom SIP headers to include when transferring the call. Each header can be either a static value or a dynamic variable reference.
	CustomSipHeaders    []*PhoneNumberTransferCustomSipHeadersItem `json:"custom_sip_headers,omitempty" xml:"custom_sip_headers,omitempty" maxItems:"20"`
	TransferDestination *PhoneNumberTransferTransferDestination    `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	TransferType        *TransferTypeEnum                          `json:"transfer_type,omitempty" xml:"transfer_type,omitempty"`
	// When True, a ringing tone is played on the original call leg while a SIP REFER transfer completes. The tone is carried over RTP to the SIP peer executing the REFER, so disable this if the receiving system (e.g. an SBC or contact center) should not hear it. When disabled the caller hears silence until the transfer completes. SIP REFER transfers only.
	SipReferPlayDialtone *bool `json:"sip_refer_play_dialtone,omitempty" xml:"sip_refer_play_dialtone,omitempty"`
	// User-to-User Information (RFC 7433) to attach to SIP REFER transfers. Carries call context such as CRM identifiers or escalation reason across the transfer boundary.
	Uui *UuiTransferConfig `json:"uui,omitempty" xml:"uui,omitempty"`
	// DTMF digits to send after call connects (e.g., 'ww1234' for extension). Can be either a static value or a dynamic variable reference. Use 'w' for 0.5s pause. Only supported for Twilio transfers.
	PostDialDigits  *PhoneNumberTransferPostDialDigits `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty"`
	PhoneNumber     *string                            `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	Condition       string                             `json:"condition" xml:"condition" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PhoneNumberTransfer) GetCustomSipHeaders() (value []*PhoneNumberTransferCustomSipHeadersItem) {
	if p == nil {
		return
	}
	return p.CustomSipHeaders
}

func (p *PhoneNumberTransfer) GetTransferDestination() (value *PhoneNumberTransferTransferDestination) {
	if p == nil {
		return
	}
	return p.TransferDestination
}

func (p *PhoneNumberTransfer) GetTransferType() (value *TransferTypeEnum) {
	if p == nil {
		return
	}
	return p.TransferType
}

func (p *PhoneNumberTransfer) GetSipReferPlayDialtone() (value *bool) {
	if p == nil {
		return
	}
	return p.SipReferPlayDialtone
}

func (p *PhoneNumberTransfer) GetUui() (value *UuiTransferConfig) {
	if p == nil {
		return
	}
	return p.Uui
}

func (p *PhoneNumberTransfer) GetPostDialDigits() (value *PhoneNumberTransferPostDialDigits) {
	if p == nil {
		return
	}
	return p.PostDialDigits
}

func (p *PhoneNumberTransfer) GetPhoneNumber() (value *string) {
	if p == nil {
		return
	}
	return p.PhoneNumber
}

func (p *PhoneNumberTransfer) GetCondition() (value string) {
	if p == nil {
		return
	}
	return p.Condition
}

func (p *PhoneNumberTransfer) SetCustomSipHeaders(customSipHeaders []*PhoneNumberTransferCustomSipHeadersItem) {
	p.CustomSipHeaders = customSipHeaders
}

func (p *PhoneNumberTransfer) SetTransferDestination(transferDestination *PhoneNumberTransferTransferDestination) {
	p.TransferDestination = transferDestination
}

func (p *PhoneNumberTransfer) SetTransferType(transferType *TransferTypeEnum) {
	p.TransferType = transferType
}

func (p *PhoneNumberTransfer) SetSipReferPlayDialtone(sipReferPlayDialtone *bool) {
	p.SipReferPlayDialtone = sipReferPlayDialtone
}

func (p *PhoneNumberTransfer) SetUui(uui *UuiTransferConfig) {
	p.Uui = uui
}

func (p *PhoneNumberTransfer) SetPostDialDigits(postDialDigits *PhoneNumberTransferPostDialDigits) {
	p.PostDialDigits = postDialDigits
}

func (p *PhoneNumberTransfer) SetPhoneNumber(phoneNumber *string) {
	p.PhoneNumber = phoneNumber
}

func (p *PhoneNumberTransfer) SetCondition(condition string) {
	p.Condition = condition
}

func (p *PhoneNumberTransfer) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneNumberTransfer) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneNumberTransfer to string"
	}
	return string(jsonData)
}

func (p *PhoneNumberTransfer) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneNumberTransfer) MarshalJSON() ([]byte, error) {
	type embed PhoneNumberTransfer
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}

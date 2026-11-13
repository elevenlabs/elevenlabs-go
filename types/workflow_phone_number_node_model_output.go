package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowPhoneNumberNodeModelOutput struct {
	// Custom SIP headers to include when transferring the call. Each header can be either a static value or a dynamic variable reference.
	CustomSipHeaders    []*WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem `json:"custom_sip_headers" xml:"custom_sip_headers" required:"true" maxItems:"20"`
	TransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination    `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	TransferType        TransferTypeEnum                                          `json:"transfer_type" xml:"transfer_type" required:"true"`
	// When True, a ringing tone is played on the original call leg while a SIP REFER transfer completes. The tone is carried over RTP to the SIP peer executing the REFER, so disable this if the receiving system (e.g. an SBC or contact center) should not hear it. When disabled the caller hears silence until the transfer completes. SIP REFER transfers only.
	SipReferPlayDialtone bool `json:"sip_refer_play_dialtone" xml:"sip_refer_play_dialtone" required:"true"`
	// User-to-User Information (RFC 7433) to attach to SIP REFER transfers. Carries call context such as CRM identifiers or escalation reason across the transfer boundary.
	Uui *UuiTransferConfig `json:"uui" xml:"uui" required:"true"`
	// DTMF digits to send after call connects (e.g., 'ww1234' for extension). Can be either a static value or a dynamic variable reference. Use 'w' for 0.5s pause. Only supported for Twilio transfers.
	PostDialDigits *WorkflowPhoneNumberNodeModelOutputPostDialDigits `json:"post_dial_digits" xml:"post_dial_digits" required:"true"`
	type_          string
	Position       *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder       []string `json:"edge_order" xml:"edge_order" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetCustomSipHeaders() (value []*WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) {
	if w == nil {
		return
	}
	return w.CustomSipHeaders
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetTransferDestination() (value *WorkflowPhoneNumberNodeModelOutputTransferDestination) {
	if w == nil {
		return
	}
	return w.TransferDestination
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetTransferType() (value TransferTypeEnum) {
	if w == nil {
		return
	}
	return w.TransferType
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetSipReferPlayDialtone() (value bool) {
	if w == nil {
		return
	}
	return w.SipReferPlayDialtone
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetUui() (value *UuiTransferConfig) {
	if w == nil {
		return
	}
	return w.Uui
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetPostDialDigits() (value *WorkflowPhoneNumberNodeModelOutputPostDialDigits) {
	if w == nil {
		return
	}
	return w.PostDialDigits
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetCustomSipHeaders(customSipHeaders []*WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) {
	w.CustomSipHeaders = customSipHeaders
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetTransferDestination(transferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) {
	w.TransferDestination = transferDestination
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetTransferType(transferType TransferTypeEnum) {
	w.TransferType = transferType
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetSipReferPlayDialtone(sipReferPlayDialtone bool) {
	w.SipReferPlayDialtone = sipReferPlayDialtone
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetUui(uui *UuiTransferConfig) {
	w.Uui = uui
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetPostDialDigits(postDialDigits *WorkflowPhoneNumberNodeModelOutputPostDialDigits) {
	w.PostDialDigits = postDialDigits
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowPhoneNumberNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowPhoneNumberNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowPhoneNumberNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowPhoneNumberNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowPhoneNumberNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowPhoneNumberNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "phone_number" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "phone_number", *literals.Type)
	}
	w.type_ = "phone_number"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowPhoneNumberNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowPhoneNumberNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "phone_number",
	}
	return json.Marshal(marshaler)
}

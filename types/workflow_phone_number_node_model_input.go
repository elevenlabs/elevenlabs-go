package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowPhoneNumberNodeModelInput struct {
	// Custom SIP headers to include when transferring the call. Each header can be either a static value or a dynamic variable reference.
	CustomSipHeaders    []*WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem `json:"custom_sip_headers,omitempty" xml:"custom_sip_headers,omitempty" maxItems:"20"`
	TransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination    `json:"transfer_destination" xml:"transfer_destination" required:"true"`
	TransferType        *TransferTypeEnum                                        `json:"transfer_type,omitempty" xml:"transfer_type,omitempty"`
	// When True, a ringing tone is played on the original call leg while a SIP REFER transfer completes. The tone is carried over RTP to the SIP peer executing the REFER, so disable this if the receiving system (e.g. an SBC or contact center) should not hear it. When disabled the caller hears silence until the transfer completes. SIP REFER transfers only.
	SipReferPlayDialtone *bool `json:"sip_refer_play_dialtone,omitempty" xml:"sip_refer_play_dialtone,omitempty"`
	// User-to-User Information (RFC 7433) to attach to SIP REFER transfers. Carries call context such as CRM identifiers or escalation reason across the transfer boundary.
	Uui *UuiTransferConfig `json:"uui,omitempty" xml:"uui,omitempty"`
	// DTMF digits to send after call connects (e.g., 'ww1234' for extension). Can be either a static value or a dynamic variable reference. Use 'w' for 0.5s pause. Only supported for Twilio transfers.
	PostDialDigits *WorkflowPhoneNumberNodeModelInputPostDialDigits `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty"`
	Position       *PositionInput                                   `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder []string `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	// Per-provider transfer configuration. Mutually exclusive with a non-empty `transfer_destination`: when present, the configuration matching the call's telephony provider is executed and the legacy transfer fields are ignored.
	ProviderConfigs []*WorkflowPhoneNumberNodeModelInputProviderConfigsItem `json:"provider_configs,omitempty" xml:"provider_configs,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowPhoneNumberNodeModelInput) GetCustomSipHeaders() (value []*WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) {
	if w == nil {
		return
	}
	return w.CustomSipHeaders
}

func (w *WorkflowPhoneNumberNodeModelInput) GetTransferDestination() (value *WorkflowPhoneNumberNodeModelInputTransferDestination) {
	if w == nil {
		return
	}
	return w.TransferDestination
}

func (w *WorkflowPhoneNumberNodeModelInput) GetTransferType() (value *TransferTypeEnum) {
	if w == nil {
		return
	}
	return w.TransferType
}

func (w *WorkflowPhoneNumberNodeModelInput) GetSipReferPlayDialtone() (value *bool) {
	if w == nil {
		return
	}
	return w.SipReferPlayDialtone
}

func (w *WorkflowPhoneNumberNodeModelInput) GetUui() (value *UuiTransferConfig) {
	if w == nil {
		return
	}
	return w.Uui
}

func (w *WorkflowPhoneNumberNodeModelInput) GetPostDialDigits() (value *WorkflowPhoneNumberNodeModelInputPostDialDigits) {
	if w == nil {
		return
	}
	return w.PostDialDigits
}

func (w *WorkflowPhoneNumberNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowPhoneNumberNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowPhoneNumberNodeModelInput) GetProviderConfigs() (value []*WorkflowPhoneNumberNodeModelInputProviderConfigsItem) {
	if w == nil {
		return
	}
	return w.ProviderConfigs
}

func (w *WorkflowPhoneNumberNodeModelInput) SetCustomSipHeaders(customSipHeaders []*WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) {
	w.CustomSipHeaders = customSipHeaders
}

func (w *WorkflowPhoneNumberNodeModelInput) SetTransferDestination(transferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) {
	w.TransferDestination = transferDestination
}

func (w *WorkflowPhoneNumberNodeModelInput) SetTransferType(transferType *TransferTypeEnum) {
	w.TransferType = transferType
}

func (w *WorkflowPhoneNumberNodeModelInput) SetSipReferPlayDialtone(sipReferPlayDialtone *bool) {
	w.SipReferPlayDialtone = sipReferPlayDialtone
}

func (w *WorkflowPhoneNumberNodeModelInput) SetUui(uui *UuiTransferConfig) {
	w.Uui = uui
}

func (w *WorkflowPhoneNumberNodeModelInput) SetPostDialDigits(postDialDigits *WorkflowPhoneNumberNodeModelInputPostDialDigits) {
	w.PostDialDigits = postDialDigits
}

func (w *WorkflowPhoneNumberNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowPhoneNumberNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowPhoneNumberNodeModelInput) SetProviderConfigs(providerConfigs []*WorkflowPhoneNumberNodeModelInputProviderConfigsItem) {
	w.ProviderConfigs = providerConfigs
}

func (w *WorkflowPhoneNumberNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowPhoneNumberNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowPhoneNumberNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowPhoneNumberNodeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowPhoneNumberNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowPhoneNumberNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

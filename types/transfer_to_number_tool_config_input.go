package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferToNumberToolConfigInput struct {
	Transfers []*PhoneNumberTransfer `json:"transfers" xml:"transfers" required:"true"`
	// Per-provider transfer configuration. Mutually exclusive with `transfers`: a tool sets exactly one of the two. When present, the configuration matching the call's telephony provider is offered to the agent and executed.
	TransferRules []*TransferRuleInput `json:"transfer_rules,omitempty" xml:"transfer_rules,omitempty"`
	// Whether to play a message to the client while they wait for transfer. Defaults to true for backward compatibility.
	EnableClientMessage *bool `json:"enable_client_message,omitempty" xml:"enable_client_message,omitempty"`
	extraProperties     map[string]interface{}
}

func (t *TransferToNumberToolConfigInput) GetTransfers() (value []*PhoneNumberTransfer) {
	if t == nil {
		return
	}
	return t.Transfers
}

func (t *TransferToNumberToolConfigInput) GetTransferRules() (value []*TransferRuleInput) {
	if t == nil {
		return
	}
	return t.TransferRules
}

func (t *TransferToNumberToolConfigInput) GetEnableClientMessage() (value *bool) {
	if t == nil {
		return
	}
	return t.EnableClientMessage
}

func (t *TransferToNumberToolConfigInput) SetTransfers(transfers []*PhoneNumberTransfer) {
	t.Transfers = transfers
}

func (t *TransferToNumberToolConfigInput) SetTransferRules(transferRules []*TransferRuleInput) {
	t.TransferRules = transferRules
}

func (t *TransferToNumberToolConfigInput) SetEnableClientMessage(enableClientMessage *bool) {
	t.EnableClientMessage = enableClientMessage
}

func (t *TransferToNumberToolConfigInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberToolConfigInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberToolConfigInput to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberToolConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler TransferToNumberToolConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TransferToNumberToolConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "system_tool_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToNumberToolConfigInput) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberToolConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

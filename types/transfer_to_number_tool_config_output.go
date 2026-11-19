package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferToNumberToolConfigOutput struct {
	Transfers []*PhoneNumberTransfer `json:"transfers" xml:"transfers" required:"true"`
	// Per-provider transfer configuration. Mutually exclusive with `transfers`: a tool sets exactly one of the two. When present, the configuration matching the call's telephony provider is offered to the agent and executed.
	TransferRules []*TransferRuleOutput `json:"transfer_rules,omitempty" xml:"transfer_rules,omitempty"`
	// Whether to play a message to the client while they wait for transfer. Defaults to true for backward compatibility.
	EnableClientMessage *bool `json:"enable_client_message,omitempty" xml:"enable_client_message,omitempty"`
	extraProperties     map[string]interface{}
}

func (t *TransferToNumberToolConfigOutput) GetTransfers() (value []*PhoneNumberTransfer) {
	if t == nil {
		return
	}
	return t.Transfers
}

func (t *TransferToNumberToolConfigOutput) GetTransferRules() (value []*TransferRuleOutput) {
	if t == nil {
		return
	}
	return t.TransferRules
}

func (t *TransferToNumberToolConfigOutput) GetEnableClientMessage() (value *bool) {
	if t == nil {
		return
	}
	return t.EnableClientMessage
}

func (t *TransferToNumberToolConfigOutput) SetTransfers(transfers []*PhoneNumberTransfer) {
	t.Transfers = transfers
}

func (t *TransferToNumberToolConfigOutput) SetTransferRules(transferRules []*TransferRuleOutput) {
	t.TransferRules = transferRules
}

func (t *TransferToNumberToolConfigOutput) SetEnableClientMessage(enableClientMessage *bool) {
	t.EnableClientMessage = enableClientMessage
}

func (t *TransferToNumberToolConfigOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberToolConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberToolConfigOutput to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberToolConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler TransferToNumberToolConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TransferToNumberToolConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "system_tool_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToNumberToolConfigOutput) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberToolConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

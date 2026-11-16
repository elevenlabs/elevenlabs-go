package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferToAgentToolConfigInput struct {
	Transfers       []*AgentTransferInput `json:"transfers" xml:"transfers" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferToAgentToolConfigInput) GetTransfers() (value []*AgentTransferInput) {
	if t == nil {
		return
	}
	return t.Transfers
}

func (t *TransferToAgentToolConfigInput) SetTransfers(transfers []*AgentTransferInput) {
	t.Transfers = transfers
}

func (t *TransferToAgentToolConfigInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToAgentToolConfigInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToAgentToolConfigInput to string"
	}
	return string(jsonData)
}

func (t *TransferToAgentToolConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler TransferToAgentToolConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TransferToAgentToolConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "system_tool_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToAgentToolConfigInput) MarshalJSON() ([]byte, error) {
	type embed TransferToAgentToolConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

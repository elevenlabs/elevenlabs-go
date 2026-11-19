package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferToAgentToolConfigOutput struct {
	Transfers       []*AgentTransferOutput `json:"transfers" xml:"transfers" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferToAgentToolConfigOutput) GetTransfers() (value []*AgentTransferOutput) {
	if t == nil {
		return
	}
	return t.Transfers
}

func (t *TransferToAgentToolConfigOutput) SetTransfers(transfers []*AgentTransferOutput) {
	t.Transfers = transfers
}

func (t *TransferToAgentToolConfigOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToAgentToolConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToAgentToolConfigOutput to string"
	}
	return string(jsonData)
}

func (t *TransferToAgentToolConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler TransferToAgentToolConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TransferToAgentToolConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "system_tool_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToAgentToolConfigOutput) MarshalJSON() ([]byte, error) {
	type embed TransferToAgentToolConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

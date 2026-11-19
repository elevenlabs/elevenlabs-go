package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToNumberResultExotelSuccessModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	TransferNumber  string  `json:"transfer_number" xml:"transfer_number" required:"true"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	AgentMessage    *string `json:"agent_message,omitempty" xml:"agent_message,omitempty"`
	Note            *string `json:"note,omitempty" xml:"note,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TransferToNumberResultExotelSuccessModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToNumberResultExotelSuccessModel) GetTransferNumber() (value string) {
	if t == nil {
		return
	}
	return t.TransferNumber
}

func (t *TransferToNumberResultExotelSuccessModel) GetReason() (value *string) {
	if t == nil {
		return
	}
	return t.Reason
}

func (t *TransferToNumberResultExotelSuccessModel) GetAgentMessage() (value *string) {
	if t == nil {
		return
	}
	return t.AgentMessage
}

func (t *TransferToNumberResultExotelSuccessModel) GetNote() (value *string) {
	if t == nil {
		return
	}
	return t.Note
}

func (t *TransferToNumberResultExotelSuccessModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToNumberResultExotelSuccessModel) SetTransferNumber(transferNumber string) {
	t.TransferNumber = transferNumber
}

func (t *TransferToNumberResultExotelSuccessModel) SetReason(reason *string) {
	t.Reason = reason
}

func (t *TransferToNumberResultExotelSuccessModel) SetAgentMessage(agentMessage *string) {
	t.AgentMessage = agentMessage
}

func (t *TransferToNumberResultExotelSuccessModel) SetNote(note *string) {
	t.Note = note
}

func (t *TransferToNumberResultExotelSuccessModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberResultExotelSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberResultExotelSuccessModel to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberResultExotelSuccessModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "result_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToNumberResultExotelSuccessModel) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberResultExotelSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

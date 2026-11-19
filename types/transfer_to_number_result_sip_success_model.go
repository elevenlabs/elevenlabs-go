package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToNumberResultSipSuccessModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	TransferNumber  string  `json:"transfer_number" xml:"transfer_number" required:"true"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Note            *string `json:"note,omitempty" xml:"note,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TransferToNumberResultSipSuccessModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToNumberResultSipSuccessModel) GetTransferNumber() (value string) {
	if t == nil {
		return
	}
	return t.TransferNumber
}

func (t *TransferToNumberResultSipSuccessModel) GetReason() (value *string) {
	if t == nil {
		return
	}
	return t.Reason
}

func (t *TransferToNumberResultSipSuccessModel) GetNote() (value *string) {
	if t == nil {
		return
	}
	return t.Note
}

func (t *TransferToNumberResultSipSuccessModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToNumberResultSipSuccessModel) SetTransferNumber(transferNumber string) {
	t.TransferNumber = transferNumber
}

func (t *TransferToNumberResultSipSuccessModel) SetReason(reason *string) {
	t.Reason = reason
}

func (t *TransferToNumberResultSipSuccessModel) SetNote(note *string) {
	t.Note = note
}

func (t *TransferToNumberResultSipSuccessModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberResultSipSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberResultSipSuccessModel to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberResultSipSuccessModel) UnmarshalJSON(data []byte) error {
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

func (t *TransferToNumberResultSipSuccessModel) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberResultSipSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

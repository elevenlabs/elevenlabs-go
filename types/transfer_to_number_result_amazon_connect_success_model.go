package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToNumberResultAmazonConnectSuccessModel struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Why the agent escalated, as sent to Amazon Connect with the Escalate outcome.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// Message spoken to the caller before the session was handed back to Amazon Connect.
	ClientMessage *string `json:"client_message,omitempty" xml:"client_message,omitempty"`
	// Set when the agent echoed an invalid option and the only configured one was used.
	Note            *string `json:"note,omitempty" xml:"note,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) GetReason() (value *string) {
	if t == nil {
		return
	}
	return t.Reason
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) GetClientMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ClientMessage
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) GetNote() (value *string) {
	if t == nil {
		return
	}
	return t.Note
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) SetReason(reason *string) {
	t.Reason = reason
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) SetClientMessage(clientMessage *string) {
	t.ClientMessage = clientMessage
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) SetNote(note *string) {
	t.Note = note
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberResultAmazonConnectSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberResultAmazonConnectSuccessModel to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberResultAmazonConnectSuccessModel) UnmarshalJSON(data []byte) error {
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

func (t *TransferToNumberResultAmazonConnectSuccessModel) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberResultAmazonConnectSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

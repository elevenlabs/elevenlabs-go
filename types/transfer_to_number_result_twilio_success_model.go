package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToNumberResultTwilioSuccessModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	TransferNumber  string  `json:"transfer_number" xml:"transfer_number" required:"true"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	ClientMessage   *string `json:"client_message,omitempty" xml:"client_message,omitempty"`
	AgentMessage    string  `json:"agent_message" xml:"agent_message" required:"true"`
	ConferenceName  string  `json:"conference_name" xml:"conference_name" required:"true"`
	PostDialDigits  *string `json:"post_dial_digits,omitempty" xml:"post_dial_digits,omitempty"`
	Note            *string `json:"note,omitempty" xml:"note,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TransferToNumberResultTwilioSuccessModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToNumberResultTwilioSuccessModel) GetTransferNumber() (value string) {
	if t == nil {
		return
	}
	return t.TransferNumber
}

func (t *TransferToNumberResultTwilioSuccessModel) GetReason() (value *string) {
	if t == nil {
		return
	}
	return t.Reason
}

func (t *TransferToNumberResultTwilioSuccessModel) GetClientMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ClientMessage
}

func (t *TransferToNumberResultTwilioSuccessModel) GetAgentMessage() (value string) {
	if t == nil {
		return
	}
	return t.AgentMessage
}

func (t *TransferToNumberResultTwilioSuccessModel) GetConferenceName() (value string) {
	if t == nil {
		return
	}
	return t.ConferenceName
}

func (t *TransferToNumberResultTwilioSuccessModel) GetPostDialDigits() (value *string) {
	if t == nil {
		return
	}
	return t.PostDialDigits
}

func (t *TransferToNumberResultTwilioSuccessModel) GetNote() (value *string) {
	if t == nil {
		return
	}
	return t.Note
}

func (t *TransferToNumberResultTwilioSuccessModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToNumberResultTwilioSuccessModel) SetTransferNumber(transferNumber string) {
	t.TransferNumber = transferNumber
}

func (t *TransferToNumberResultTwilioSuccessModel) SetReason(reason *string) {
	t.Reason = reason
}

func (t *TransferToNumberResultTwilioSuccessModel) SetClientMessage(clientMessage *string) {
	t.ClientMessage = clientMessage
}

func (t *TransferToNumberResultTwilioSuccessModel) SetAgentMessage(agentMessage string) {
	t.AgentMessage = agentMessage
}

func (t *TransferToNumberResultTwilioSuccessModel) SetConferenceName(conferenceName string) {
	t.ConferenceName = conferenceName
}

func (t *TransferToNumberResultTwilioSuccessModel) SetPostDialDigits(postDialDigits *string) {
	t.PostDialDigits = postDialDigits
}

func (t *TransferToNumberResultTwilioSuccessModel) SetNote(note *string) {
	t.Note = note
}

func (t *TransferToNumberResultTwilioSuccessModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToNumberResultTwilioSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToNumberResultTwilioSuccessModel to string"
	}
	return string(jsonData)
}

func (t *TransferToNumberResultTwilioSuccessModel) UnmarshalJSON(data []byte) error {
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

func (t *TransferToNumberResultTwilioSuccessModel) MarshalJSON() ([]byte, error) {
	type embed TransferToNumberResultTwilioSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

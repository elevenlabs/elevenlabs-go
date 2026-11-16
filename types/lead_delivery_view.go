package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LeadDeliveryView struct {
	Status BatchCallRecipientStatus `json:"status" xml:"status" required:"true"`
	// The furthest the message got, or `failed` if sending it failed.
	Stage           LeadDeliveryViewStage `json:"stage" xml:"stage" required:"true"`
	SentAtUnix      *int                  `json:"sent_at_unix" xml:"sent_at_unix" required:"true"`
	DeliveredAtUnix *int                  `json:"delivered_at_unix" xml:"delivered_at_unix" required:"true"`
	ReadAtUnix      *int                  `json:"read_at_unix" xml:"read_at_unix" required:"true"`
	RepliedAtUnix   *int                  `json:"replied_at_unix" xml:"replied_at_unix" required:"true"`
	FailedAtUnix    *int                  `json:"failed_at_unix" xml:"failed_at_unix" required:"true"`
	// Why the message failed to reach the contact.
	ErrorMessage    *string `json:"error_message" xml:"error_message" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LeadDeliveryView) GetStatus() (value BatchCallRecipientStatus) {
	if l == nil {
		return
	}
	return l.Status
}

func (l *LeadDeliveryView) GetStage() (value LeadDeliveryViewStage) {
	if l == nil {
		return
	}
	return l.Stage
}

func (l *LeadDeliveryView) GetSentAtUnix() (value *int) {
	if l == nil {
		return
	}
	return l.SentAtUnix
}

func (l *LeadDeliveryView) GetDeliveredAtUnix() (value *int) {
	if l == nil {
		return
	}
	return l.DeliveredAtUnix
}

func (l *LeadDeliveryView) GetReadAtUnix() (value *int) {
	if l == nil {
		return
	}
	return l.ReadAtUnix
}

func (l *LeadDeliveryView) GetRepliedAtUnix() (value *int) {
	if l == nil {
		return
	}
	return l.RepliedAtUnix
}

func (l *LeadDeliveryView) GetFailedAtUnix() (value *int) {
	if l == nil {
		return
	}
	return l.FailedAtUnix
}

func (l *LeadDeliveryView) GetErrorMessage() (value *string) {
	if l == nil {
		return
	}
	return l.ErrorMessage
}

func (l *LeadDeliveryView) SetStatus(status BatchCallRecipientStatus) {
	l.Status = status
}

func (l *LeadDeliveryView) SetStage(stage LeadDeliveryViewStage) {
	l.Stage = stage
}

func (l *LeadDeliveryView) SetSentAtUnix(sentAtUnix *int) {
	l.SentAtUnix = sentAtUnix
}

func (l *LeadDeliveryView) SetDeliveredAtUnix(deliveredAtUnix *int) {
	l.DeliveredAtUnix = deliveredAtUnix
}

func (l *LeadDeliveryView) SetReadAtUnix(readAtUnix *int) {
	l.ReadAtUnix = readAtUnix
}

func (l *LeadDeliveryView) SetRepliedAtUnix(repliedAtUnix *int) {
	l.RepliedAtUnix = repliedAtUnix
}

func (l *LeadDeliveryView) SetFailedAtUnix(failedAtUnix *int) {
	l.FailedAtUnix = failedAtUnix
}

func (l *LeadDeliveryView) SetErrorMessage(errorMessage *string) {
	l.ErrorMessage = errorMessage
}

func (l *LeadDeliveryView) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LeadDeliveryView) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LeadDeliveryView to string"
	}
	return string(jsonData)
}

func (l *LeadDeliveryView) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LeadDeliveryView) MarshalJSON() ([]byte, error) {
	type embed LeadDeliveryView
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

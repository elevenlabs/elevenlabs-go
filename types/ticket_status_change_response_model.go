package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TicketStatusChangeResponseModel struct {
	Status            AgentConversationTicketStatus `json:"status" xml:"status" required:"true"`
	ChangedByUserId   *string                       `json:"changed_by_user_id" xml:"changed_by_user_id" required:"true"`
	ChangedAtUnixSecs int                           `json:"changed_at_unix_secs" xml:"changed_at_unix_secs" required:"true"`
	extraProperties   map[string]interface{}
}

func (t *TicketStatusChangeResponseModel) GetStatus() (value AgentConversationTicketStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TicketStatusChangeResponseModel) GetChangedByUserId() (value *string) {
	if t == nil {
		return
	}
	return t.ChangedByUserId
}

func (t *TicketStatusChangeResponseModel) GetChangedAtUnixSecs() (value int) {
	if t == nil {
		return
	}
	return t.ChangedAtUnixSecs
}

func (t *TicketStatusChangeResponseModel) SetStatus(status AgentConversationTicketStatus) {
	t.Status = status
}

func (t *TicketStatusChangeResponseModel) SetChangedByUserId(changedByUserId *string) {
	t.ChangedByUserId = changedByUserId
}

func (t *TicketStatusChangeResponseModel) SetChangedAtUnixSecs(changedAtUnixSecs int) {
	t.ChangedAtUnixSecs = changedAtUnixSecs
}

func (t *TicketStatusChangeResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TicketStatusChangeResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TicketStatusChangeResponseModel to string"
	}
	return string(jsonData)
}

func (t *TicketStatusChangeResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TicketStatusChangeResponseModel) MarshalJSON() ([]byte, error) {
	type embed TicketStatusChangeResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

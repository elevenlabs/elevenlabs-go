package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TicketPriorityChangeResponseModel struct {
	Priority          *AgentConversationTicketPriority `json:"priority" xml:"priority" required:"true"`
	ChangedByUserId   string                           `json:"changed_by_user_id" xml:"changed_by_user_id" required:"true"`
	ChangedAtUnixSecs int                              `json:"changed_at_unix_secs" xml:"changed_at_unix_secs" required:"true"`
	extraProperties   map[string]interface{}
}

func (t *TicketPriorityChangeResponseModel) GetPriority() (value *AgentConversationTicketPriority) {
	if t == nil {
		return
	}
	return t.Priority
}

func (t *TicketPriorityChangeResponseModel) GetChangedByUserId() (value string) {
	if t == nil {
		return
	}
	return t.ChangedByUserId
}

func (t *TicketPriorityChangeResponseModel) GetChangedAtUnixSecs() (value int) {
	if t == nil {
		return
	}
	return t.ChangedAtUnixSecs
}

func (t *TicketPriorityChangeResponseModel) SetPriority(priority *AgentConversationTicketPriority) {
	t.Priority = priority
}

func (t *TicketPriorityChangeResponseModel) SetChangedByUserId(changedByUserId string) {
	t.ChangedByUserId = changedByUserId
}

func (t *TicketPriorityChangeResponseModel) SetChangedAtUnixSecs(changedAtUnixSecs int) {
	t.ChangedAtUnixSecs = changedAtUnixSecs
}

func (t *TicketPriorityChangeResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TicketPriorityChangeResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TicketPriorityChangeResponseModel to string"
	}
	return string(jsonData)
}

func (t *TicketPriorityChangeResponseModel) UnmarshalJSON(data []byte) error {
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

func (t *TicketPriorityChangeResponseModel) MarshalJSON() ([]byte, error) {
	type embed TicketPriorityChangeResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

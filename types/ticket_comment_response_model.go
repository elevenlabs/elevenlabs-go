package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TicketCommentResponseModel struct {
	Comment           string  `json:"comment" xml:"comment" required:"true"`
	CreatedAtUnixSecs int     `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	OwnerUserId       *string `json:"owner_user_id" xml:"owner_user_id" required:"true"`
	extraProperties   map[string]interface{}
}

func (t *TicketCommentResponseModel) GetComment() (value string) {
	if t == nil {
		return
	}
	return t.Comment
}

func (t *TicketCommentResponseModel) GetCreatedAtUnixSecs() (value int) {
	if t == nil {
		return
	}
	return t.CreatedAtUnixSecs
}

func (t *TicketCommentResponseModel) GetOwnerUserId() (value *string) {
	if t == nil {
		return
	}
	return t.OwnerUserId
}

func (t *TicketCommentResponseModel) SetComment(comment string) {
	t.Comment = comment
}

func (t *TicketCommentResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	t.CreatedAtUnixSecs = createdAtUnixSecs
}

func (t *TicketCommentResponseModel) SetOwnerUserId(ownerUserId *string) {
	t.OwnerUserId = ownerUserId
}

func (t *TicketCommentResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TicketCommentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TicketCommentResponseModel to string"
	}
	return string(jsonData)
}

func (t *TicketCommentResponseModel) UnmarshalJSON(data []byte) error {
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

func (t *TicketCommentResponseModel) MarshalJSON() ([]byte, error) {
	type embed TicketCommentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

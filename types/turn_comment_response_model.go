package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TurnCommentResponseModel struct {
	TurnIndex         int     `json:"turn_index" xml:"turn_index" required:"true"`
	Comment           string  `json:"comment" xml:"comment" required:"true"`
	CreatedAtUnixSecs int     `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	OwnerUserId       *string `json:"owner_user_id" xml:"owner_user_id" required:"true"`
	extraProperties   map[string]interface{}
}

func (t *TurnCommentResponseModel) GetTurnIndex() (value int) {
	if t == nil {
		return
	}
	return t.TurnIndex
}

func (t *TurnCommentResponseModel) GetComment() (value string) {
	if t == nil {
		return
	}
	return t.Comment
}

func (t *TurnCommentResponseModel) GetCreatedAtUnixSecs() (value int) {
	if t == nil {
		return
	}
	return t.CreatedAtUnixSecs
}

func (t *TurnCommentResponseModel) GetOwnerUserId() (value *string) {
	if t == nil {
		return
	}
	return t.OwnerUserId
}

func (t *TurnCommentResponseModel) SetTurnIndex(turnIndex int) {
	t.TurnIndex = turnIndex
}

func (t *TurnCommentResponseModel) SetComment(comment string) {
	t.Comment = comment
}

func (t *TurnCommentResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	t.CreatedAtUnixSecs = createdAtUnixSecs
}

func (t *TurnCommentResponseModel) SetOwnerUserId(ownerUserId *string) {
	t.OwnerUserId = ownerUserId
}

func (t *TurnCommentResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TurnCommentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TurnCommentResponseModel to string"
	}
	return string(jsonData)
}

func (t *TurnCommentResponseModel) UnmarshalJSON(data []byte) error {
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

func (t *TurnCommentResponseModel) MarshalJSON() ([]byte, error) {
	type embed TurnCommentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

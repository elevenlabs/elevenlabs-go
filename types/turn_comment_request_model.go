package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TurnCommentRequestModel struct {
	// Zero-based index of the transcript turn this comment refers to.
	TurnIndex int `json:"turn_index" xml:"turn_index" required:"true" min:"0"`
	// What went wrong at this turn.
	Comment         string `json:"comment" xml:"comment" required:"true" maxLength:"10000" minLength:"1"`
	extraProperties map[string]interface{}
}

func (t *TurnCommentRequestModel) GetTurnIndex() (value int) {
	if t == nil {
		return
	}
	return t.TurnIndex
}

func (t *TurnCommentRequestModel) GetComment() (value string) {
	if t == nil {
		return
	}
	return t.Comment
}

func (t *TurnCommentRequestModel) SetTurnIndex(turnIndex int) {
	t.TurnIndex = turnIndex
}

func (t *TurnCommentRequestModel) SetComment(comment string) {
	t.Comment = comment
}

func (t *TurnCommentRequestModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TurnCommentRequestModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TurnCommentRequestModel to string"
	}
	return string(jsonData)
}

func (t *TurnCommentRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TurnCommentRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TurnCommentRequestModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TurnCommentRequestModel) MarshalJSON() ([]byte, error) {
	type embed TurnCommentRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

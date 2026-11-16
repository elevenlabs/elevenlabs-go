package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserTranscript struct {
	type_           string
	EventId         int                                 `json:"event_id" xml:"event_id" required:"true"`
	UserTranscript  []*UserTranscriptUserTranscriptItem `json:"user_transcript" xml:"user_transcript" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UserTranscript) GetEventId() (value int) {
	if u == nil {
		return
	}
	return u.EventId
}

func (u *UserTranscript) GetUserTranscript() (value []*UserTranscriptUserTranscriptItem) {
	if u == nil {
		return
	}
	return u.UserTranscript
}

func (u *UserTranscript) SetEventId(eventId int) {
	u.EventId = eventId
}

func (u *UserTranscript) SetUserTranscript(userTranscript []*UserTranscriptUserTranscriptItem) {
	u.UserTranscript = userTranscript
}

func (u *UserTranscript) Type() string {
	return u.type_
}

func (u *UserTranscript) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserTranscript) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserTranscript to string"
	}
	return string(jsonData)
}

func (u *UserTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UserTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UserTranscript(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "user_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", u, "user_transcript", *literals.Type)
	}
	u.type_ = "user_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserTranscript) MarshalJSON() ([]byte, error) {
	type embed UserTranscript
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*u),
		Type:  "user_transcript",
	}
	return json.Marshal(marshaler)
}

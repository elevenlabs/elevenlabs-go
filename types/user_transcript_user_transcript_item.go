package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserTranscriptUserTranscriptItem struct {
	Role            string `json:"role" xml:"role" required:"true"`
	Content         string `json:"content" xml:"content" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UserTranscriptUserTranscriptItem) GetRole() (value string) {
	if u == nil {
		return
	}
	return u.Role
}

func (u *UserTranscriptUserTranscriptItem) GetContent() (value string) {
	if u == nil {
		return
	}
	return u.Content
}

func (u *UserTranscriptUserTranscriptItem) SetRole(role string) {
	u.Role = role
}

func (u *UserTranscriptUserTranscriptItem) SetContent(content string) {
	u.Content = content
}

func (u *UserTranscriptUserTranscriptItem) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserTranscriptUserTranscriptItem) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserTranscriptUserTranscriptItem to string"
	}
	return string(jsonData)
}

func (u *UserTranscriptUserTranscriptItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UserTranscriptUserTranscriptItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UserTranscriptUserTranscriptItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserTranscriptUserTranscriptItem) MarshalJSON() ([]byte, error) {
	type embed UserTranscriptUserTranscriptItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

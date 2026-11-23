package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UserMessage struct {
	Id              *string                  `json:"id,omitempty" xml:"id,omitempty"`
	Type            *string                  `json:"type,omitempty" xml:"type,omitempty"`
	Text            *string                  `json:"text,omitempty" xml:"text,omitempty"`
	UserIdentifier  *string                  `json:"user_identifier,omitempty" xml:"user_identifier,omitempty"`
	SourceMedium    *UserMessageSourceMedium `json:"source_medium,omitempty" xml:"source_medium,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UserMessage) GetId() (value *string) {
	if u == nil {
		return
	}
	return u.Id
}

func (u *UserMessage) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UserMessage) GetText() (value *string) {
	if u == nil {
		return
	}
	return u.Text
}

func (u *UserMessage) GetUserIdentifier() (value *string) {
	if u == nil {
		return
	}
	return u.UserIdentifier
}

func (u *UserMessage) GetSourceMedium() (value *UserMessageSourceMedium) {
	if u == nil {
		return
	}
	return u.SourceMedium
}

func (u *UserMessage) SetId(id *string) {
	u.Id = id
}

func (u *UserMessage) SetType(type_ *string) {
	u.Type = type_
}

func (u *UserMessage) SetText(text *string) {
	u.Text = text
}

func (u *UserMessage) SetUserIdentifier(userIdentifier *string) {
	u.UserIdentifier = userIdentifier
}

func (u *UserMessage) SetSourceMedium(sourceMedium *UserMessageSourceMedium) {
	u.SourceMedium = sourceMedium
}

func (u *UserMessage) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserMessage) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserMessage to string"
	}
	return string(jsonData)
}

func (u *UserMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler UserMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserMessage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserMessage) MarshalJSON() ([]byte, error) {
	type embed UserMessage
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

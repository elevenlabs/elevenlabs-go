package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserMessageInputItem struct {
	type_ string
	// The parts of the message: at most one text part and up to 5 file parts.
	Content         []*UserMessageInputItemContentItem `json:"content" xml:"content" required:"true" minItems:"1" maxItems:"6"`
	extraProperties map[string]interface{}
}

func (u *UserMessageInputItem) GetContent() (value []*UserMessageInputItemContentItem) {
	if u == nil {
		return
	}
	return u.Content
}

func (u *UserMessageInputItem) SetContent(content []*UserMessageInputItemContentItem) {
	u.Content = content
}

func (u *UserMessageInputItem) Type() string {
	return u.type_
}

func (u *UserMessageInputItem) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserMessageInputItem) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserMessageInputItem to string"
	}
	return string(jsonData)
}

func (u *UserMessageInputItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UserMessageInputItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UserMessageInputItem(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "user_message" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", u, "user_message", *literals.Type)
	}
	u.type_ = "user_message"
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserMessageInputItem) MarshalJSON() ([]byte, error) {
	type embed UserMessageInputItem
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*u),
		Type:  "user_message",
	}
	return json.Marshal(marshaler)
}

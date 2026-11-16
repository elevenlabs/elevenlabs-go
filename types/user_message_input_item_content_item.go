package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UserMessageInputItemContentItem struct {
	Type string
	File *FileContentPart
	Text *TextContentPart
}

func (u *UserMessageInputItemContentItem) GetFile() *FileContentPart {
	if u == nil {
		return nil
	}
	if u.File == nil {
		return nil
	}
	return u.File
}

func (u *UserMessageInputItemContentItem) GetText() *TextContentPart {
	if u == nil {
		return nil
	}
	if u.Text == nil {
		return nil
	}
	return u.Text
}

func (u *UserMessageInputItemContentItem) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UserMessageInputItemContentItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", u)
	}
	switch unmarshaler.Type {
	case "file":
		value := new(FileContentPart)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.File = value
	case "text":
		value := new(TextContentPart)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Text = value
	}
	return nil
}

func (u UserMessageInputItemContentItem) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.File != nil {
		return internal.MarshalJSONWithDiscriminant(u.File, "type", "file")
	}
	if u.Text != nil {
		return internal.MarshalJSONWithDiscriminant(u.Text, "type", "text")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UserMessageInputItemContentItemVisitor interface {
	VisitFile(*FileContentPart) error
	VisitText(*TextContentPart) error
}

func (u *UserMessageInputItemContentItem) Accept(visitor UserMessageInputItemContentItemVisitor) error {
	if u.File != nil {
		return visitor.VisitFile(u.File)
	}
	if u.Text != nil {
		return visitor.VisitText(u.Text)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UserMessageInputItemContentItem) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.File != nil {
		fields = append(fields, "file")
	}
	if u.Text != nil {
		fields = append(fields, "text")
	}
	if len(fields) == 0 {
		if u.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Type)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Type != "" {
		field := fields[0]
		if u.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Type,
				u,
			)
		}
	}
	return nil
}

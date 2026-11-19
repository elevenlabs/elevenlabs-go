package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TextOutputItem struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	role            string
	Id              *string `json:"id,omitempty" xml:"id,omitempty"`
	Text            string  `json:"text" xml:"text" required:"true"`
	CreatedAtUnixMs *int    `json:"created_at_unix_ms,omitempty" xml:"created_at_unix_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TextOutputItem) GetType() (value *string) {
	if t == nil {
		return
	}
	return t.Type
}

func (t *TextOutputItem) GetId() (value *string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TextOutputItem) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TextOutputItem) GetCreatedAtUnixMs() (value *int) {
	if t == nil {
		return
	}
	return t.CreatedAtUnixMs
}

func (t *TextOutputItem) SetType(type_ *string) {
	t.Type = type_
}

func (t *TextOutputItem) SetId(id *string) {
	t.Id = id
}

func (t *TextOutputItem) SetText(text string) {
	t.Text = text
}

func (t *TextOutputItem) SetCreatedAtUnixMs(createdAtUnixMs *int) {
	t.CreatedAtUnixMs = createdAtUnixMs
}

func (t *TextOutputItem) Role() string {
	return t.role
}

func (t *TextOutputItem) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextOutputItem) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextOutputItem to string"
	}
	return string(jsonData)
}

func (t *TextOutputItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Role *string `json:"role"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Role != nil && *literals.Role != "agent" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "agent", *literals.Role)
	}
	t.role = "agent"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "role")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextOutputItem) MarshalJSON() ([]byte, error) {
	type embed TextOutputItem
	var marshaler = struct {
		embed
		Role string `json:"role"`
	}{
		embed: embed(*t),
		Role:  "agent",
	}
	return json.Marshal(marshaler)
}

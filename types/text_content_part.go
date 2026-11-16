package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TextContentPart struct {
	type_           string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TextContentPart) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TextContentPart) SetText(text string) {
	t.Text = text
}

func (t *TextContentPart) Type() string {
	return t.type_
}

func (t *TextContentPart) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextContentPart) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextContentPart to string"
	}
	return string(jsonData)
}

func (t *TextContentPart) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TextContentPart
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TextContentPart(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "text" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "text", *literals.Type)
	}
	t.type_ = "text"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextContentPart) MarshalJSON() ([]byte, error) {
	type embed TextContentPart
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "text",
	}
	return json.Marshal(marshaler)
}

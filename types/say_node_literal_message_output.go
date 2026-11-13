package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SayNodeLiteralMessageOutput struct {
	type_ string
	// Exact text spoken by the agent. Supports `{{variable_name}}` dynamic variable placeholders.
	Text string `json:"text" xml:"text" required:"true" minLength:"1"`
	// Translations of `text`, keyed by language code. When the conversation language has an entry, its value is spoken instead of `text`.
	TextTranslations map[string]*TranslatedString `json:"text_translations" required:"true"`
	extraProperties  map[string]interface{}
}

func (s *SayNodeLiteralMessageOutput) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SayNodeLiteralMessageOutput) GetTextTranslations() (value map[string]*TranslatedString) {
	if s == nil {
		return
	}
	return s.TextTranslations
}

func (s *SayNodeLiteralMessageOutput) SetText(text string) {
	s.Text = text
}

func (s *SayNodeLiteralMessageOutput) SetTextTranslations(textTranslations map[string]*TranslatedString) {
	s.TextTranslations = textTranslations
}

func (s *SayNodeLiteralMessageOutput) Type() string {
	return s.type_
}

func (s *SayNodeLiteralMessageOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SayNodeLiteralMessageOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SayNodeLiteralMessageOutput to string"
	}
	return string(jsonData)
}

func (s *SayNodeLiteralMessageOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SayNodeLiteralMessageOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SayNodeLiteralMessageOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "literal" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "literal", *literals.Type)
	}
	s.type_ = "literal"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SayNodeLiteralMessageOutput) MarshalJSON() ([]byte, error) {
	type embed SayNodeLiteralMessageOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "literal",
	}
	return json.Marshal(marshaler)
}

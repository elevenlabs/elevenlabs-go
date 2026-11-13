package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SayNodeLiteralMessageInput struct {
	// Exact text spoken by the agent. Supports `{{variable_name}}` dynamic variable placeholders.
	Text string `json:"text" xml:"text" required:"true" minLength:"1"`
	// Translations of `text`, keyed by language code. When the conversation language has an entry, its value is spoken instead of `text`.
	TextTranslations map[string]*TranslatedString `json:"text_translations,omitempty"`
	extraProperties  map[string]interface{}
}

func (s *SayNodeLiteralMessageInput) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SayNodeLiteralMessageInput) GetTextTranslations() (value map[string]*TranslatedString) {
	if s == nil {
		return
	}
	return s.TextTranslations
}

func (s *SayNodeLiteralMessageInput) SetText(text string) {
	s.Text = text
}

func (s *SayNodeLiteralMessageInput) SetTextTranslations(textTranslations map[string]*TranslatedString) {
	s.TextTranslations = textTranslations
}

func (s *SayNodeLiteralMessageInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SayNodeLiteralMessageInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SayNodeLiteralMessageInput to string"
	}
	return string(jsonData)
}

func (s *SayNodeLiteralMessageInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SayNodeLiteralMessageInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SayNodeLiteralMessageInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SayNodeLiteralMessageInput) MarshalJSON() ([]byte, error) {
	type embed SayNodeLiteralMessageInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Names the approved Meta template to send. The server loads the template
// itself, so the client never supplies its components.
type SubmitWhatsAppTemplateRef struct {
	Name string `json:"name" xml:"name" required:"true" maxLength:"512" minLength:"1"`
	// The template's language code, such as `en_US`.
	Language        string `json:"language" xml:"language" required:"true" maxLength:"64" minLength:"1"`
	extraProperties map[string]interface{}
}

func (s *SubmitWhatsAppTemplateRef) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SubmitWhatsAppTemplateRef) GetLanguage() (value string) {
	if s == nil {
		return
	}
	return s.Language
}

func (s *SubmitWhatsAppTemplateRef) SetName(name string) {
	s.Name = name
}

func (s *SubmitWhatsAppTemplateRef) SetLanguage(language string) {
	s.Language = language
}

func (s *SubmitWhatsAppTemplateRef) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitWhatsAppTemplateRef) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitWhatsAppTemplateRef to string"
	}
	return string(jsonData)
}

func (s *SubmitWhatsAppTemplateRef) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubmitWhatsAppTemplateRef
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubmitWhatsAppTemplateRef(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitWhatsAppTemplateRef) MarshalJSON() ([]byte, error) {
	type embed SubmitWhatsAppTemplateRef
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

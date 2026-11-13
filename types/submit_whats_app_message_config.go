package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubmitWhatsAppMessageConfig struct {
	// Names the approved Meta template to send. The server loads the template
	// itself, so the client never supplies its components.
	Template        *SubmitWhatsAppTemplateRef `json:"template" xml:"template" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubmitWhatsAppMessageConfig) GetTemplate() (value *SubmitWhatsAppTemplateRef) {
	if s == nil {
		return
	}
	return s.Template
}

func (s *SubmitWhatsAppMessageConfig) SetTemplate(template *SubmitWhatsAppTemplateRef) {
	s.Template = template
}

func (s *SubmitWhatsAppMessageConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitWhatsAppMessageConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitWhatsAppMessageConfig to string"
	}
	return string(jsonData)
}

func (s *SubmitWhatsAppMessageConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubmitWhatsAppMessageConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubmitWhatsAppMessageConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "mode")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitWhatsAppMessageConfig) MarshalJSON() ([]byte, error) {
	type embed SubmitWhatsAppMessageConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

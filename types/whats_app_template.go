package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplate struct {
	Name string `json:"name" xml:"name" required:"true"`
	// https://developers.facebook.com/documentation/business-messaging/whatsapp/templates/supported-languages
	Language        string                           `json:"language" xml:"language" required:"true"`
	ParameterFormat *WhatsAppTemplateParameterFormat `json:"parameter_format,omitempty" xml:"parameter_format,omitempty"`
	Components      []*WhatsAppTemplateComponent     `json:"components" xml:"components" required:"true"`
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/
	Category        *string `json:"category,omitempty" xml:"category,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplate) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WhatsAppTemplate) GetLanguage() (value string) {
	if w == nil {
		return
	}
	return w.Language
}

func (w *WhatsAppTemplate) GetParameterFormat() (value *WhatsAppTemplateParameterFormat) {
	if w == nil {
		return
	}
	return w.ParameterFormat
}

func (w *WhatsAppTemplate) GetComponents() (value []*WhatsAppTemplateComponent) {
	if w == nil {
		return
	}
	return w.Components
}

func (w *WhatsAppTemplate) GetCategory() (value *string) {
	if w == nil {
		return
	}
	return w.Category
}

func (w *WhatsAppTemplate) SetName(name string) {
	w.Name = name
}

func (w *WhatsAppTemplate) SetLanguage(language string) {
	w.Language = language
}

func (w *WhatsAppTemplate) SetParameterFormat(parameterFormat *WhatsAppTemplateParameterFormat) {
	w.ParameterFormat = parameterFormat
}

func (w *WhatsAppTemplate) SetComponents(components []*WhatsAppTemplateComponent) {
	w.Components = components
}

func (w *WhatsAppTemplate) SetCategory(category *string) {
	w.Category = category
}

func (w *WhatsAppTemplate) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplate) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplate to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplate) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplate) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplate
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

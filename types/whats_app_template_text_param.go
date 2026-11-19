package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateTextParam struct {
	ParameterName   *string `json:"parameter_name,omitempty" xml:"parameter_name,omitempty"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	Text            string  `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateTextParam) GetParameterName() (value *string) {
	if w == nil {
		return
	}
	return w.ParameterName
}

func (w *WhatsAppTemplateTextParam) GetType() (value *string) {
	if w == nil {
		return
	}
	return w.Type
}

func (w *WhatsAppTemplateTextParam) GetText() (value string) {
	if w == nil {
		return
	}
	return w.Text
}

func (w *WhatsAppTemplateTextParam) SetParameterName(parameterName *string) {
	w.ParameterName = parameterName
}

func (w *WhatsAppTemplateTextParam) SetType(type_ *string) {
	w.Type = type_
}

func (w *WhatsAppTemplateTextParam) SetText(text string) {
	w.Text = text
}

func (w *WhatsAppTemplateTextParam) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateTextParam) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateTextParam to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateTextParam) UnmarshalJSON(data []byte) error {
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

func (w *WhatsAppTemplateTextParam) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateTextParam
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

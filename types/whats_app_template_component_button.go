package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateComponentButton struct {
	Type            string  `json:"type" xml:"type" required:"true"`
	Text            *string `json:"text,omitempty" xml:"text,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateComponentButton) GetType() (value string) {
	if w == nil {
		return
	}
	return w.Type
}

func (w *WhatsAppTemplateComponentButton) GetText() (value *string) {
	if w == nil {
		return
	}
	return w.Text
}

func (w *WhatsAppTemplateComponentButton) GetUrl() (value *string) {
	if w == nil {
		return
	}
	return w.Url
}

func (w *WhatsAppTemplateComponentButton) SetType(type_ string) {
	w.Type = type_
}

func (w *WhatsAppTemplateComponentButton) SetText(text *string) {
	w.Text = text
}

func (w *WhatsAppTemplateComponentButton) SetUrl(url *string) {
	w.Url = url
}

func (w *WhatsAppTemplateComponentButton) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateComponentButton) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateComponentButton to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateComponentButton) UnmarshalJSON(data []byte) error {
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

func (w *WhatsAppTemplateComponentButton) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateComponentButton
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

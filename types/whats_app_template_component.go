package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateComponent struct {
	Type            string                             `json:"type" xml:"type" required:"true"`
	Format          *string                            `json:"format,omitempty" xml:"format,omitempty"`
	Text            *string                            `json:"text,omitempty" xml:"text,omitempty"`
	Buttons         []*WhatsAppTemplateComponentButton `json:"buttons,omitempty" xml:"buttons,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateComponent) GetType() (value string) {
	if w == nil {
		return
	}
	return w.Type
}

func (w *WhatsAppTemplateComponent) GetFormat() (value *string) {
	if w == nil {
		return
	}
	return w.Format
}

func (w *WhatsAppTemplateComponent) GetText() (value *string) {
	if w == nil {
		return
	}
	return w.Text
}

func (w *WhatsAppTemplateComponent) GetButtons() (value []*WhatsAppTemplateComponentButton) {
	if w == nil {
		return
	}
	return w.Buttons
}

func (w *WhatsAppTemplateComponent) SetType(type_ string) {
	w.Type = type_
}

func (w *WhatsAppTemplateComponent) SetFormat(format *string) {
	w.Format = format
}

func (w *WhatsAppTemplateComponent) SetText(text *string) {
	w.Text = text
}

func (w *WhatsAppTemplateComponent) SetButtons(buttons []*WhatsAppTemplateComponentButton) {
	w.Buttons = buttons
}

func (w *WhatsAppTemplateComponent) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateComponent) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateComponent to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateComponent) UnmarshalJSON(data []byte) error {
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

func (w *WhatsAppTemplateComponent) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateComponent
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

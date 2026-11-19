package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WhatsAppTemplateBodyComponentParams struct {
	Parameters      []*WhatsAppTemplateTextParam `json:"parameters" xml:"parameters" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateBodyComponentParams) GetParameters() (value []*WhatsAppTemplateTextParam) {
	if w == nil {
		return
	}
	return w.Parameters
}

func (w *WhatsAppTemplateBodyComponentParams) SetParameters(parameters []*WhatsAppTemplateTextParam) {
	w.Parameters = parameters
}

func (w *WhatsAppTemplateBodyComponentParams) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateBodyComponentParams) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateBodyComponentParams to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateBodyComponentParams) UnmarshalJSON(data []byte) error {
	type unmarshaler WhatsAppTemplateBodyComponentParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WhatsAppTemplateBodyComponentParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateBodyComponentParams) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateBodyComponentParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

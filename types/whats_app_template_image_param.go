package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateImageParam struct {
	Image           *WhatsAppTemplateImageParamDetails `json:"image" xml:"image" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateImageParam) GetImage() (value *WhatsAppTemplateImageParamDetails) {
	if w == nil {
		return
	}
	return w.Image
}

func (w *WhatsAppTemplateImageParam) SetImage(image *WhatsAppTemplateImageParamDetails) {
	w.Image = image
}

func (w *WhatsAppTemplateImageParam) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateImageParam) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateImageParam to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateImageParam) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppTemplateImageParam
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppTemplateImageParam(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateImageParam) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateImageParam
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

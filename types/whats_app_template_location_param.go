package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateLocationParam struct {
	Location        *WhatsAppTemplateLocationParamDetails `json:"location" xml:"location" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateLocationParam) GetLocation() (value *WhatsAppTemplateLocationParamDetails) {
	if w == nil {
		return
	}
	return w.Location
}

func (w *WhatsAppTemplateLocationParam) SetLocation(location *WhatsAppTemplateLocationParamDetails) {
	w.Location = location
}

func (w *WhatsAppTemplateLocationParam) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateLocationParam) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateLocationParam to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateLocationParam) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppTemplateLocationParam
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppTemplateLocationParam(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateLocationParam) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateLocationParam
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

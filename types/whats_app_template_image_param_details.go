package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateImageParamDetails struct {
	Link            string `json:"link" xml:"link" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateImageParamDetails) GetLink() (value string) {
	if w == nil {
		return
	}
	return w.Link
}

func (w *WhatsAppTemplateImageParamDetails) SetLink(link string) {
	w.Link = link
}

func (w *WhatsAppTemplateImageParamDetails) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateImageParamDetails) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateImageParamDetails to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateImageParamDetails) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppTemplateImageParamDetails
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppTemplateImageParamDetails(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateImageParamDetails) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateImageParamDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

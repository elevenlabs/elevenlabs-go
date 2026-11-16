package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateDocumentParam struct {
	Document        *WhatsAppTemplateDocumentParamDetails `json:"document" xml:"document" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateDocumentParam) GetDocument() (value *WhatsAppTemplateDocumentParamDetails) {
	if w == nil {
		return
	}
	return w.Document
}

func (w *WhatsAppTemplateDocumentParam) SetDocument(document *WhatsAppTemplateDocumentParamDetails) {
	w.Document = document
}

func (w *WhatsAppTemplateDocumentParam) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateDocumentParam) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateDocumentParam to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateDocumentParam) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppTemplateDocumentParam
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppTemplateDocumentParam(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateDocumentParam) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateDocumentParam
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

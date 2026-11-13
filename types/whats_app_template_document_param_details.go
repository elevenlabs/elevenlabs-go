package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateDocumentParamDetails struct {
	Link            string  `json:"link" xml:"link" required:"true"`
	Filename        *string `json:"filename,omitempty" xml:"filename,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateDocumentParamDetails) GetLink() (value string) {
	if w == nil {
		return
	}
	return w.Link
}

func (w *WhatsAppTemplateDocumentParamDetails) GetFilename() (value *string) {
	if w == nil {
		return
	}
	return w.Filename
}

func (w *WhatsAppTemplateDocumentParamDetails) SetLink(link string) {
	w.Link = link
}

func (w *WhatsAppTemplateDocumentParamDetails) SetFilename(filename *string) {
	w.Filename = filename
}

func (w *WhatsAppTemplateDocumentParamDetails) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateDocumentParamDetails) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateDocumentParamDetails to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateDocumentParamDetails) UnmarshalJSON(data []byte) error {
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

func (w *WhatsAppTemplateDocumentParamDetails) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateDocumentParamDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

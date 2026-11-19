package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppTemplateButtonComponentParams struct {
	Parameters      []*WhatsAppTemplateTextParam `json:"parameters" xml:"parameters" required:"true"`
	Index           int                          `json:"index" xml:"index" required:"true"`
	SubType         string                       `json:"sub_type" xml:"sub_type" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateButtonComponentParams) GetParameters() (value []*WhatsAppTemplateTextParam) {
	if w == nil {
		return
	}
	return w.Parameters
}

func (w *WhatsAppTemplateButtonComponentParams) GetIndex() (value int) {
	if w == nil {
		return
	}
	return w.Index
}

func (w *WhatsAppTemplateButtonComponentParams) GetSubType() (value string) {
	if w == nil {
		return
	}
	return w.SubType
}

func (w *WhatsAppTemplateButtonComponentParams) SetParameters(parameters []*WhatsAppTemplateTextParam) {
	w.Parameters = parameters
}

func (w *WhatsAppTemplateButtonComponentParams) SetIndex(index int) {
	w.Index = index
}

func (w *WhatsAppTemplateButtonComponentParams) SetSubType(subType string) {
	w.SubType = subType
}

func (w *WhatsAppTemplateButtonComponentParams) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateButtonComponentParams) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateButtonComponentParams to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateButtonComponentParams) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppTemplateButtonComponentParams
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppTemplateButtonComponentParams(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateButtonComponentParams) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateButtonComponentParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

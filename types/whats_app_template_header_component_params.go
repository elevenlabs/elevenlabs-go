package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WhatsAppTemplateHeaderComponentParams struct {
	Parameters      []*WhatsAppTemplateHeaderComponentParamsParametersItem `json:"parameters" xml:"parameters" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppTemplateHeaderComponentParams) GetParameters() (value []*WhatsAppTemplateHeaderComponentParamsParametersItem) {
	if w == nil {
		return
	}
	return w.Parameters
}

func (w *WhatsAppTemplateHeaderComponentParams) SetParameters(parameters []*WhatsAppTemplateHeaderComponentParamsParametersItem) {
	w.Parameters = parameters
}

func (w *WhatsAppTemplateHeaderComponentParams) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppTemplateHeaderComponentParams) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppTemplateHeaderComponentParams to string"
	}
	return string(jsonData)
}

func (w *WhatsAppTemplateHeaderComponentParams) UnmarshalJSON(data []byte) error {
	type unmarshaler WhatsAppTemplateHeaderComponentParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WhatsAppTemplateHeaderComponentParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppTemplateHeaderComponentParams) MarshalJSON() ([]byte, error) {
	type embed WhatsAppTemplateHeaderComponentParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

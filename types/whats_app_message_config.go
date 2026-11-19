package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppMessageConfig struct {
	Template        *WhatsAppTemplate `json:"template" xml:"template" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WhatsAppMessageConfig) GetTemplate() (value *WhatsAppTemplate) {
	if w == nil {
		return
	}
	return w.Template
}

func (w *WhatsAppMessageConfig) SetTemplate(template *WhatsAppTemplate) {
	w.Template = template
}

func (w *WhatsAppMessageConfig) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppMessageConfig) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppMessageConfig to string"
	}
	return string(jsonData)
}

func (w *WhatsAppMessageConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppMessageConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppMessageConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "mode")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppMessageConfig) MarshalJSON() ([]byte, error) {
	type embed WhatsAppMessageConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppVoiceConfig struct {
	// The approved template that asks the contact for permission to call.
	WhatsappCallPermissionRequestTemplateName string `json:"whatsapp_call_permission_request_template_name" xml:"whatsapp_call_permission_request_template_name" required:"true"`
	// The language code of that template.
	WhatsappCallPermissionRequestTemplateLanguageCode string `json:"whatsapp_call_permission_request_template_language_code" xml:"whatsapp_call_permission_request_template_language_code" required:"true"`
	extraProperties                                   map[string]interface{}
}

func (w *WhatsAppVoiceConfig) GetWhatsappCallPermissionRequestTemplateName() (value string) {
	if w == nil {
		return
	}
	return w.WhatsappCallPermissionRequestTemplateName
}

func (w *WhatsAppVoiceConfig) GetWhatsappCallPermissionRequestTemplateLanguageCode() (value string) {
	if w == nil {
		return
	}
	return w.WhatsappCallPermissionRequestTemplateLanguageCode
}

func (w *WhatsAppVoiceConfig) SetWhatsappCallPermissionRequestTemplateName(whatsappCallPermissionRequestTemplateName string) {
	w.WhatsappCallPermissionRequestTemplateName = whatsappCallPermissionRequestTemplateName
}

func (w *WhatsAppVoiceConfig) SetWhatsappCallPermissionRequestTemplateLanguageCode(whatsappCallPermissionRequestTemplateLanguageCode string) {
	w.WhatsappCallPermissionRequestTemplateLanguageCode = whatsappCallPermissionRequestTemplateLanguageCode
}

func (w *WhatsAppVoiceConfig) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppVoiceConfig) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppVoiceConfig to string"
	}
	return string(jsonData)
}

func (w *WhatsAppVoiceConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppVoiceConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppVoiceConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "mode")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppVoiceConfig) MarshalJSON() ([]byte, error) {
	type embed WhatsAppVoiceConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

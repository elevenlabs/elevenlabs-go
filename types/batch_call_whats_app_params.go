package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BatchCallWhatsAppParams struct {
	WhatsappPhoneNumberId                             *string `json:"whatsapp_phone_number_id,omitempty" xml:"whatsapp_phone_number_id,omitempty"`
	WhatsappCallPermissionRequestTemplateName         string  `json:"whatsapp_call_permission_request_template_name" xml:"whatsapp_call_permission_request_template_name" required:"true"`
	WhatsappCallPermissionRequestTemplateLanguageCode string  `json:"whatsapp_call_permission_request_template_language_code" xml:"whatsapp_call_permission_request_template_language_code" required:"true"`
	extraProperties                                   map[string]interface{}
}

func (b *BatchCallWhatsAppParams) GetWhatsappPhoneNumberId() (value *string) {
	if b == nil {
		return
	}
	return b.WhatsappPhoneNumberId
}

func (b *BatchCallWhatsAppParams) GetWhatsappCallPermissionRequestTemplateName() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappCallPermissionRequestTemplateName
}

func (b *BatchCallWhatsAppParams) GetWhatsappCallPermissionRequestTemplateLanguageCode() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappCallPermissionRequestTemplateLanguageCode
}

func (b *BatchCallWhatsAppParams) SetWhatsappPhoneNumberId(whatsappPhoneNumberId *string) {
	b.WhatsappPhoneNumberId = whatsappPhoneNumberId
}

func (b *BatchCallWhatsAppParams) SetWhatsappCallPermissionRequestTemplateName(whatsappCallPermissionRequestTemplateName string) {
	b.WhatsappCallPermissionRequestTemplateName = whatsappCallPermissionRequestTemplateName
}

func (b *BatchCallWhatsAppParams) SetWhatsappCallPermissionRequestTemplateLanguageCode(whatsappCallPermissionRequestTemplateLanguageCode string) {
	b.WhatsappCallPermissionRequestTemplateLanguageCode = whatsappCallPermissionRequestTemplateLanguageCode
}

func (b *BatchCallWhatsAppParams) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BatchCallWhatsAppParams) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BatchCallWhatsAppParams to string"
	}
	return string(jsonData)
}

func (b *BatchCallWhatsAppParams) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BatchCallWhatsAppParams) MarshalJSON() ([]byte, error) {
	type embed BatchCallWhatsAppParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

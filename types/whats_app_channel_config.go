package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppChannelConfig struct {
	// The WhatsApp phone number to call or message from.
	WhatsappPhoneNumberId string                           `json:"whatsapp_phone_number_id" xml:"whatsapp_phone_number_id" required:"true" minLength:"1"`
	ModeConfig            *WhatsAppChannelConfigModeConfig `json:"mode_config" xml:"mode_config" required:"true"`
	extraProperties       map[string]interface{}
}

func (w *WhatsAppChannelConfig) GetWhatsappPhoneNumberId() (value string) {
	if w == nil {
		return
	}
	return w.WhatsappPhoneNumberId
}

func (w *WhatsAppChannelConfig) GetModeConfig() (value *WhatsAppChannelConfigModeConfig) {
	if w == nil {
		return
	}
	return w.ModeConfig
}

func (w *WhatsAppChannelConfig) SetWhatsappPhoneNumberId(whatsappPhoneNumberId string) {
	w.WhatsappPhoneNumberId = whatsappPhoneNumberId
}

func (w *WhatsAppChannelConfig) SetModeConfig(modeConfig *WhatsAppChannelConfigModeConfig) {
	w.ModeConfig = modeConfig
}

func (w *WhatsAppChannelConfig) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppChannelConfig) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppChannelConfig to string"
	}
	return string(jsonData)
}

func (w *WhatsAppChannelConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WhatsAppChannelConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WhatsAppChannelConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "channel")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppChannelConfig) MarshalJSON() ([]byte, error) {
	type embed WhatsAppChannelConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}

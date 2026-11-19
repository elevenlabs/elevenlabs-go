package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubmitWhatsAppChannelConfig struct {
	// The WhatsApp phone number to call or message from.
	WhatsappPhoneNumberId string                                 `json:"whatsapp_phone_number_id" xml:"whatsapp_phone_number_id" required:"true" minLength:"1"`
	ModeConfig            *SubmitWhatsAppChannelConfigModeConfig `json:"mode_config" xml:"mode_config" required:"true"`
	extraProperties       map[string]interface{}
}

func (s *SubmitWhatsAppChannelConfig) GetWhatsappPhoneNumberId() (value string) {
	if s == nil {
		return
	}
	return s.WhatsappPhoneNumberId
}

func (s *SubmitWhatsAppChannelConfig) GetModeConfig() (value *SubmitWhatsAppChannelConfigModeConfig) {
	if s == nil {
		return
	}
	return s.ModeConfig
}

func (s *SubmitWhatsAppChannelConfig) SetWhatsappPhoneNumberId(whatsappPhoneNumberId string) {
	s.WhatsappPhoneNumberId = whatsappPhoneNumberId
}

func (s *SubmitWhatsAppChannelConfig) SetModeConfig(modeConfig *SubmitWhatsAppChannelConfigModeConfig) {
	s.ModeConfig = modeConfig
}

func (s *SubmitWhatsAppChannelConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitWhatsAppChannelConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitWhatsAppChannelConfig to string"
	}
	return string(jsonData)
}

func (s *SubmitWhatsAppChannelConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubmitWhatsAppChannelConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubmitWhatsAppChannelConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "channel")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitWhatsAppChannelConfig) MarshalJSON() ([]byte, error) {
	type embed SubmitWhatsAppChannelConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}

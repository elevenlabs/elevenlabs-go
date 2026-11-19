package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SubmitWhatsAppChannelConfigModeConfig struct {
	Mode    string
	Message *SubmitWhatsAppMessageConfig
	Voice   *WhatsAppVoiceConfig
}

func (s *SubmitWhatsAppChannelConfigModeConfig) GetMessage() *SubmitWhatsAppMessageConfig {
	if s == nil {
		return nil
	}
	if s.Message == nil {
		return nil
	}
	return s.Message
}

func (s *SubmitWhatsAppChannelConfigModeConfig) GetVoice() *WhatsAppVoiceConfig {
	if s == nil {
		return nil
	}
	if s.Voice == nil {
		return nil
	}
	return s.Voice
}

func (s *SubmitWhatsAppChannelConfigModeConfig) GetMode() string {
	if s == nil {
		return ""
	}
	return s.Mode
}

func (s *SubmitWhatsAppChannelConfigModeConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Mode string `json:"mode"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Mode = unmarshaler.Mode
	if unmarshaler.Mode == "" {
		return fmt.Errorf("%T did not include discriminant mode", s)
	}
	switch unmarshaler.Mode {
	case "message":
		value := new(SubmitWhatsAppMessageConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Message = value
	case "voice":
		value := new(WhatsAppVoiceConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Voice = value
	}
	return nil
}

func (s SubmitWhatsAppChannelConfigModeConfig) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Message != nil {
		return internal.MarshalJSONWithDiscriminant(s.Message, "mode", "message")
	}
	if s.Voice != nil {
		return internal.MarshalJSONWithDiscriminant(s.Voice, "mode", "voice")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SubmitWhatsAppChannelConfigModeConfigVisitor interface {
	VisitMessage(*SubmitWhatsAppMessageConfig) error
	VisitVoice(*WhatsAppVoiceConfig) error
}

func (s *SubmitWhatsAppChannelConfigModeConfig) Accept(visitor SubmitWhatsAppChannelConfigModeConfigVisitor) error {
	if s.Message != nil {
		return visitor.VisitMessage(s.Message)
	}
	if s.Voice != nil {
		return visitor.VisitVoice(s.Voice)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SubmitWhatsAppChannelConfigModeConfig) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Message != nil {
		fields = append(fields, "message")
	}
	if s.Voice != nil {
		fields = append(fields, "voice")
	}
	if len(fields) == 0 {
		if s.Mode != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Mode)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Mode != "" {
		field := fields[0]
		if s.Mode != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Mode,
				s,
			)
		}
	}
	return nil
}

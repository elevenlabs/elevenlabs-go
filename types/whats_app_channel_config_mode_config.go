package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WhatsAppChannelConfigModeConfig struct {
	Mode    string
	Message *WhatsAppMessageConfig
	Voice   *WhatsAppVoiceConfig
}

func (w *WhatsAppChannelConfigModeConfig) GetMessage() *WhatsAppMessageConfig {
	if w == nil {
		return nil
	}
	if w.Message == nil {
		return nil
	}
	return w.Message
}

func (w *WhatsAppChannelConfigModeConfig) GetVoice() *WhatsAppVoiceConfig {
	if w == nil {
		return nil
	}
	if w.Voice == nil {
		return nil
	}
	return w.Voice
}

func (w *WhatsAppChannelConfigModeConfig) GetMode() string {
	if w == nil {
		return ""
	}
	return w.Mode
}

func (w *WhatsAppChannelConfigModeConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Mode string `json:"mode"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Mode = unmarshaler.Mode
	if unmarshaler.Mode == "" {
		return fmt.Errorf("%T did not include discriminant mode", w)
	}
	switch unmarshaler.Mode {
	case "message":
		value := new(WhatsAppMessageConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Message = value
	case "voice":
		value := new(WhatsAppVoiceConfig)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Voice = value
	}
	return nil
}

func (w WhatsAppChannelConfigModeConfig) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Message != nil {
		return internal.MarshalJSONWithDiscriminant(w.Message, "mode", "message")
	}
	if w.Voice != nil {
		return internal.MarshalJSONWithDiscriminant(w.Voice, "mode", "voice")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WhatsAppChannelConfigModeConfigVisitor interface {
	VisitMessage(*WhatsAppMessageConfig) error
	VisitVoice(*WhatsAppVoiceConfig) error
}

func (w *WhatsAppChannelConfigModeConfig) Accept(visitor WhatsAppChannelConfigModeConfigVisitor) error {
	if w.Message != nil {
		return visitor.VisitMessage(w.Message)
	}
	if w.Voice != nil {
		return visitor.VisitVoice(w.Voice)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WhatsAppChannelConfigModeConfig) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Message != nil {
		fields = append(fields, "message")
	}
	if w.Voice != nil {
		fields = append(fields, "voice")
	}
	if len(fields) == 0 {
		if w.Mode != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Mode)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Mode != "" {
		field := fields[0]
		if w.Mode != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Mode,
				w,
			)
		}
	}
	return nil
}

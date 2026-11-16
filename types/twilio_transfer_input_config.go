package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TwilioTransferInputConfig struct {
	Type                         string
	TwilioColdBlindTransfer      *TwilioColdBlindTransferConfigInput
	TwilioWarmConferenceTransfer *TwilioWarmConferenceTransferConfigInput
}

func (t *TwilioTransferInputConfig) GetTwilioColdBlindTransfer() *TwilioColdBlindTransferConfigInput {
	if t == nil {
		return nil
	}
	if t.TwilioColdBlindTransfer == nil {
		return nil
	}
	return t.TwilioColdBlindTransfer
}

func (t *TwilioTransferInputConfig) GetTwilioWarmConferenceTransfer() *TwilioWarmConferenceTransferConfigInput {
	if t == nil {
		return nil
	}
	if t.TwilioWarmConferenceTransfer == nil {
		return nil
	}
	return t.TwilioWarmConferenceTransfer
}

func (t *TwilioTransferInputConfig) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TwilioTransferInputConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", t)
	}
	switch unmarshaler.Type {
	case "twilio_cold_blind_transfer":
		value := new(TwilioColdBlindTransferConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.TwilioColdBlindTransfer = value
	case "twilio_warm_conference_transfer":
		value := new(TwilioWarmConferenceTransferConfigInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.TwilioWarmConferenceTransfer = value
	}
	return nil
}

func (t TwilioTransferInputConfig) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.TwilioColdBlindTransfer != nil {
		return internal.MarshalJSONWithDiscriminant(t.TwilioColdBlindTransfer, "type", "twilio_cold_blind_transfer")
	}
	if t.TwilioWarmConferenceTransfer != nil {
		return internal.MarshalJSONWithDiscriminant(t.TwilioWarmConferenceTransfer, "type", "twilio_warm_conference_transfer")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TwilioTransferInputConfigVisitor interface {
	VisitTwilioColdBlindTransfer(*TwilioColdBlindTransferConfigInput) error
	VisitTwilioWarmConferenceTransfer(*TwilioWarmConferenceTransferConfigInput) error
}

func (t *TwilioTransferInputConfig) Accept(visitor TwilioTransferInputConfigVisitor) error {
	if t.TwilioColdBlindTransfer != nil {
		return visitor.VisitTwilioColdBlindTransfer(t.TwilioColdBlindTransfer)
	}
	if t.TwilioWarmConferenceTransfer != nil {
		return visitor.VisitTwilioWarmConferenceTransfer(t.TwilioWarmConferenceTransfer)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TwilioTransferInputConfig) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.TwilioColdBlindTransfer != nil {
		fields = append(fields, "twilio_cold_blind_transfer")
	}
	if t.TwilioWarmConferenceTransfer != nil {
		fields = append(fields, "twilio_warm_conference_transfer")
	}
	if len(fields) == 0 {
		if t.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.Type)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.Type != "" {
		field := fields[0]
		if t.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.Type,
				t,
			)
		}
	}
	return nil
}

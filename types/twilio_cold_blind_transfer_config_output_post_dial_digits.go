package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TwilioColdBlindTransferConfigOutputPostDialDigits struct {
	Type    string
	Dynamic *PostDialDigitsDynamicVariable
	Static  *PostDialDigitsStatic
}

func (t *TwilioColdBlindTransferConfigOutputPostDialDigits) GetDynamic() *PostDialDigitsDynamicVariable {
	if t == nil {
		return nil
	}
	if t.Dynamic == nil {
		return nil
	}
	return t.Dynamic
}

func (t *TwilioColdBlindTransferConfigOutputPostDialDigits) GetStatic() *PostDialDigitsStatic {
	if t == nil {
		return nil
	}
	if t.Static == nil {
		return nil
	}
	return t.Static
}

func (t *TwilioColdBlindTransferConfigOutputPostDialDigits) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TwilioColdBlindTransferConfigOutputPostDialDigits) UnmarshalJSON(data []byte) error {
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
	case "dynamic":
		value := new(PostDialDigitsDynamicVariable)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Dynamic = value
	case "static":
		value := new(PostDialDigitsStatic)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Static = value
	}
	return nil
}

func (t TwilioColdBlindTransferConfigOutputPostDialDigits) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Dynamic != nil {
		return internal.MarshalJSONWithDiscriminant(t.Dynamic, "type", "dynamic")
	}
	if t.Static != nil {
		return internal.MarshalJSONWithDiscriminant(t.Static, "type", "static")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TwilioColdBlindTransferConfigOutputPostDialDigitsVisitor interface {
	VisitDynamic(*PostDialDigitsDynamicVariable) error
	VisitStatic(*PostDialDigitsStatic) error
}

func (t *TwilioColdBlindTransferConfigOutputPostDialDigits) Accept(visitor TwilioColdBlindTransferConfigOutputPostDialDigitsVisitor) error {
	if t.Dynamic != nil {
		return visitor.VisitDynamic(t.Dynamic)
	}
	if t.Static != nil {
		return visitor.VisitStatic(t.Static)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TwilioColdBlindTransferConfigOutputPostDialDigits) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Dynamic != nil {
		fields = append(fields, "dynamic")
	}
	if t.Static != nil {
		fields = append(fields, "static")
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

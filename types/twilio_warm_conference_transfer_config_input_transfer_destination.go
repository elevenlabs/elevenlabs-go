package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TwilioWarmConferenceTransferConfigInputTransferDestination struct {
	Type                 string
	Phone                *PhoneDestinationInput
	PhoneDynamicVariable *PhoneDynamicVariableDestinationInput
}

func (t *TwilioWarmConferenceTransferConfigInputTransferDestination) GetPhone() *PhoneDestinationInput {
	if t == nil {
		return nil
	}
	if t.Phone == nil {
		return nil
	}
	return t.Phone
}

func (t *TwilioWarmConferenceTransferConfigInputTransferDestination) GetPhoneDynamicVariable() *PhoneDynamicVariableDestinationInput {
	if t == nil {
		return nil
	}
	if t.PhoneDynamicVariable == nil {
		return nil
	}
	return t.PhoneDynamicVariable
}

func (t *TwilioWarmConferenceTransferConfigInputTransferDestination) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TwilioWarmConferenceTransferConfigInputTransferDestination) UnmarshalJSON(data []byte) error {
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
	case "phone":
		value := new(PhoneDestinationInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneDynamicVariableDestinationInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.PhoneDynamicVariable = value
	}
	return nil
}

func (t TwilioWarmConferenceTransferConfigInputTransferDestination) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(t.Phone, "type", "phone")
	}
	if t.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(t.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TwilioWarmConferenceTransferConfigInputTransferDestinationVisitor interface {
	VisitPhone(*PhoneDestinationInput) error
	VisitPhoneDynamicVariable(*PhoneDynamicVariableDestinationInput) error
}

func (t *TwilioWarmConferenceTransferConfigInputTransferDestination) Accept(visitor TwilioWarmConferenceTransferConfigInputTransferDestinationVisitor) error {
	if t.Phone != nil {
		return visitor.VisitPhone(t.Phone)
	}
	if t.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(t.PhoneDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TwilioWarmConferenceTransferConfigInputTransferDestination) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Phone != nil {
		fields = append(fields, "phone")
	}
	if t.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
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

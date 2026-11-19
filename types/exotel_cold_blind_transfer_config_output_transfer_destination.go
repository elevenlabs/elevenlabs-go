package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ExotelColdBlindTransferConfigOutputTransferDestination struct {
	Type                 string
	Phone                *PhoneDestinationOutput
	PhoneDynamicVariable *PhoneDynamicVariableDestinationOutput
}

func (e *ExotelColdBlindTransferConfigOutputTransferDestination) GetPhone() *PhoneDestinationOutput {
	if e == nil {
		return nil
	}
	if e.Phone == nil {
		return nil
	}
	return e.Phone
}

func (e *ExotelColdBlindTransferConfigOutputTransferDestination) GetPhoneDynamicVariable() *PhoneDynamicVariableDestinationOutput {
	if e == nil {
		return nil
	}
	if e.PhoneDynamicVariable == nil {
		return nil
	}
	return e.PhoneDynamicVariable
}

func (e *ExotelColdBlindTransferConfigOutputTransferDestination) GetType() string {
	if e == nil {
		return ""
	}
	return e.Type
}

func (e *ExotelColdBlindTransferConfigOutputTransferDestination) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", e)
	}
	switch unmarshaler.Type {
	case "phone":
		value := new(PhoneDestinationOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneDynamicVariableDestinationOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.PhoneDynamicVariable = value
	}
	return nil
}

func (e ExotelColdBlindTransferConfigOutputTransferDestination) MarshalJSON() ([]byte, error) {
	if err := e.validate(); err != nil {
		return nil, err
	}
	if e.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(e.Phone, "type", "phone")
	}
	if e.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(e.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", e)
}

type ExotelColdBlindTransferConfigOutputTransferDestinationVisitor interface {
	VisitPhone(*PhoneDestinationOutput) error
	VisitPhoneDynamicVariable(*PhoneDynamicVariableDestinationOutput) error
}

func (e *ExotelColdBlindTransferConfigOutputTransferDestination) Accept(visitor ExotelColdBlindTransferConfigOutputTransferDestinationVisitor) error {
	if e.Phone != nil {
		return visitor.VisitPhone(e.Phone)
	}
	if e.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(e.PhoneDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", e)
}

func (e *ExotelColdBlindTransferConfigOutputTransferDestination) validate() error {
	if e == nil {
		return fmt.Errorf("type %T is nil", e)
	}
	var fields []string
	if e.Phone != nil {
		fields = append(fields, "phone")
	}
	if e.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
	}
	if len(fields) == 0 {
		if e.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", e, e.Type)
		}
		return fmt.Errorf("type %T is empty", e)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", e, fields)
	}
	if e.Type != "" {
		field := fields[0]
		if e.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				e,
				e.Type,
				e,
			)
		}
	}
	return nil
}

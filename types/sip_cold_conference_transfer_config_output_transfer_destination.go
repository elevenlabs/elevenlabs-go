package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SipColdConferenceTransferConfigOutputTransferDestination struct {
	Type                 string
	Phone                *PhoneDestinationOutput
	PhoneDynamicVariable *PhoneDynamicVariableDestinationOutput
}

func (s *SipColdConferenceTransferConfigOutputTransferDestination) GetPhone() *PhoneDestinationOutput {
	if s == nil {
		return nil
	}
	if s.Phone == nil {
		return nil
	}
	return s.Phone
}

func (s *SipColdConferenceTransferConfigOutputTransferDestination) GetPhoneDynamicVariable() *PhoneDynamicVariableDestinationOutput {
	if s == nil {
		return nil
	}
	if s.PhoneDynamicVariable == nil {
		return nil
	}
	return s.PhoneDynamicVariable
}

func (s *SipColdConferenceTransferConfigOutputTransferDestination) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *SipColdConferenceTransferConfigOutputTransferDestination) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", s)
	}
	switch unmarshaler.Type {
	case "phone":
		value := new(PhoneDestinationOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneDynamicVariableDestinationOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.PhoneDynamicVariable = value
	}
	return nil
}

func (s SipColdConferenceTransferConfigOutputTransferDestination) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(s.Phone, "type", "phone")
	}
	if s.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(s.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SipColdConferenceTransferConfigOutputTransferDestinationVisitor interface {
	VisitPhone(*PhoneDestinationOutput) error
	VisitPhoneDynamicVariable(*PhoneDynamicVariableDestinationOutput) error
}

func (s *SipColdConferenceTransferConfigOutputTransferDestination) Accept(visitor SipColdConferenceTransferConfigOutputTransferDestinationVisitor) error {
	if s.Phone != nil {
		return visitor.VisitPhone(s.Phone)
	}
	if s.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(s.PhoneDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SipColdConferenceTransferConfigOutputTransferDestination) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Phone != nil {
		fields = append(fields, "phone")
	}
	if s.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
	}
	if len(fields) == 0 {
		if s.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Type)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Type != "" {
		field := fields[0]
		if s.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Type,
				s,
			)
		}
	}
	return nil
}

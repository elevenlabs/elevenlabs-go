package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SipColdReferTransferConfigInputCustomSipHeadersItem struct {
	Type    string
	Dynamic *CustomSipHeaderWithDynamicVariable
	Static  *CustomSipHeader
}

func (s *SipColdReferTransferConfigInputCustomSipHeadersItem) GetDynamic() *CustomSipHeaderWithDynamicVariable {
	if s == nil {
		return nil
	}
	if s.Dynamic == nil {
		return nil
	}
	return s.Dynamic
}

func (s *SipColdReferTransferConfigInputCustomSipHeadersItem) GetStatic() *CustomSipHeader {
	if s == nil {
		return nil
	}
	if s.Static == nil {
		return nil
	}
	return s.Static
}

func (s *SipColdReferTransferConfigInputCustomSipHeadersItem) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *SipColdReferTransferConfigInputCustomSipHeadersItem) UnmarshalJSON(data []byte) error {
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
	case "dynamic":
		value := new(CustomSipHeaderWithDynamicVariable)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Dynamic = value
	case "static":
		value := new(CustomSipHeader)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Static = value
	}
	return nil
}

func (s SipColdReferTransferConfigInputCustomSipHeadersItem) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Dynamic != nil {
		return internal.MarshalJSONWithDiscriminant(s.Dynamic, "type", "dynamic")
	}
	if s.Static != nil {
		return internal.MarshalJSONWithDiscriminant(s.Static, "type", "static")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SipColdReferTransferConfigInputCustomSipHeadersItemVisitor interface {
	VisitDynamic(*CustomSipHeaderWithDynamicVariable) error
	VisitStatic(*CustomSipHeader) error
}

func (s *SipColdReferTransferConfigInputCustomSipHeadersItem) Accept(visitor SipColdReferTransferConfigInputCustomSipHeadersItemVisitor) error {
	if s.Dynamic != nil {
		return visitor.VisitDynamic(s.Dynamic)
	}
	if s.Static != nil {
		return visitor.VisitStatic(s.Static)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SipColdReferTransferConfigInputCustomSipHeadersItem) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Dynamic != nil {
		fields = append(fields, "dynamic")
	}
	if s.Static != nil {
		fields = append(fields, "static")
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

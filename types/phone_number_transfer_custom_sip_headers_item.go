package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PhoneNumberTransferCustomSipHeadersItem struct {
	Type    string
	Dynamic *CustomSipHeaderWithDynamicVariable
	Static  *CustomSipHeader
}

func (p *PhoneNumberTransferCustomSipHeadersItem) GetDynamic() *CustomSipHeaderWithDynamicVariable {
	if p == nil {
		return nil
	}
	if p.Dynamic == nil {
		return nil
	}
	return p.Dynamic
}

func (p *PhoneNumberTransferCustomSipHeadersItem) GetStatic() *CustomSipHeader {
	if p == nil {
		return nil
	}
	if p.Static == nil {
		return nil
	}
	return p.Static
}

func (p *PhoneNumberTransferCustomSipHeadersItem) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *PhoneNumberTransferCustomSipHeadersItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", p)
	}
	switch unmarshaler.Type {
	case "dynamic":
		value := new(CustomSipHeaderWithDynamicVariable)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Dynamic = value
	case "static":
		value := new(CustomSipHeader)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Static = value
	}
	return nil
}

func (p PhoneNumberTransferCustomSipHeadersItem) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.Dynamic != nil {
		return internal.MarshalJSONWithDiscriminant(p.Dynamic, "type", "dynamic")
	}
	if p.Static != nil {
		return internal.MarshalJSONWithDiscriminant(p.Static, "type", "static")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PhoneNumberTransferCustomSipHeadersItemVisitor interface {
	VisitDynamic(*CustomSipHeaderWithDynamicVariable) error
	VisitStatic(*CustomSipHeader) error
}

func (p *PhoneNumberTransferCustomSipHeadersItem) Accept(visitor PhoneNumberTransferCustomSipHeadersItemVisitor) error {
	if p.Dynamic != nil {
		return visitor.VisitDynamic(p.Dynamic)
	}
	if p.Static != nil {
		return visitor.VisitStatic(p.Static)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *PhoneNumberTransferCustomSipHeadersItem) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.Dynamic != nil {
		fields = append(fields, "dynamic")
	}
	if p.Static != nil {
		fields = append(fields, "static")
	}
	if len(fields) == 0 {
		if p.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Type)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Type != "" {
		field := fields[0]
		if p.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Type,
				p,
			)
		}
	}
	return nil
}

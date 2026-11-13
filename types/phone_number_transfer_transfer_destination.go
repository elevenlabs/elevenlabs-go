package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PhoneNumberTransferTransferDestination struct {
	Type                  string
	Phone                 *PhoneNumberTransferDestination
	PhoneDynamicVariable  *PhoneNumberDynamicVariableTransferDestination
	SipUri                *SipUriTransferDestination
	SipUriDynamicVariable *SipUriDynamicVariableTransferDestination
}

func (p *PhoneNumberTransferTransferDestination) GetPhone() *PhoneNumberTransferDestination {
	if p == nil {
		return nil
	}
	if p.Phone == nil {
		return nil
	}
	return p.Phone
}

func (p *PhoneNumberTransferTransferDestination) GetPhoneDynamicVariable() *PhoneNumberDynamicVariableTransferDestination {
	if p == nil {
		return nil
	}
	if p.PhoneDynamicVariable == nil {
		return nil
	}
	return p.PhoneDynamicVariable
}

func (p *PhoneNumberTransferTransferDestination) GetSipUri() *SipUriTransferDestination {
	if p == nil {
		return nil
	}
	if p.SipUri == nil {
		return nil
	}
	return p.SipUri
}

func (p *PhoneNumberTransferTransferDestination) GetSipUriDynamicVariable() *SipUriDynamicVariableTransferDestination {
	if p == nil {
		return nil
	}
	if p.SipUriDynamicVariable == nil {
		return nil
	}
	return p.SipUriDynamicVariable
}

func (p *PhoneNumberTransferTransferDestination) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *PhoneNumberTransferTransferDestination) UnmarshalJSON(data []byte) error {
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
	case "phone":
		value := new(PhoneNumberTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneNumberDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.PhoneDynamicVariable = value
	case "sip_uri":
		value := new(SipUriTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.SipUri = value
	case "sip_uri_dynamic_variable":
		value := new(SipUriDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.SipUriDynamicVariable = value
	}
	return nil
}

func (p PhoneNumberTransferTransferDestination) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(p.Phone, "type", "phone")
	}
	if p.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(p.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	if p.SipUri != nil {
		return internal.MarshalJSONWithDiscriminant(p.SipUri, "type", "sip_uri")
	}
	if p.SipUriDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(p.SipUriDynamicVariable, "type", "sip_uri_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PhoneNumberTransferTransferDestinationVisitor interface {
	VisitPhone(*PhoneNumberTransferDestination) error
	VisitPhoneDynamicVariable(*PhoneNumberDynamicVariableTransferDestination) error
	VisitSipUri(*SipUriTransferDestination) error
	VisitSipUriDynamicVariable(*SipUriDynamicVariableTransferDestination) error
}

func (p *PhoneNumberTransferTransferDestination) Accept(visitor PhoneNumberTransferTransferDestinationVisitor) error {
	if p.Phone != nil {
		return visitor.VisitPhone(p.Phone)
	}
	if p.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(p.PhoneDynamicVariable)
	}
	if p.SipUri != nil {
		return visitor.VisitSipUri(p.SipUri)
	}
	if p.SipUriDynamicVariable != nil {
		return visitor.VisitSipUriDynamicVariable(p.SipUriDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *PhoneNumberTransferTransferDestination) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.Phone != nil {
		fields = append(fields, "phone")
	}
	if p.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
	}
	if p.SipUri != nil {
		fields = append(fields, "sip_uri")
	}
	if p.SipUriDynamicVariable != nil {
		fields = append(fields, "sip_uri_dynamic_variable")
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

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowPhoneNumberNodeModelInputTransferDestination struct {
	Type                  string
	Phone                 *PhoneNumberTransferDestination
	PhoneDynamicVariable  *PhoneNumberDynamicVariableTransferDestination
	SipUri                *SipUriTransferDestination
	SipUriDynamicVariable *SipUriDynamicVariableTransferDestination
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) GetPhone() *PhoneNumberTransferDestination {
	if w == nil {
		return nil
	}
	if w.Phone == nil {
		return nil
	}
	return w.Phone
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) GetPhoneDynamicVariable() *PhoneNumberDynamicVariableTransferDestination {
	if w == nil {
		return nil
	}
	if w.PhoneDynamicVariable == nil {
		return nil
	}
	return w.PhoneDynamicVariable
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) GetSipUri() *SipUriTransferDestination {
	if w == nil {
		return nil
	}
	if w.SipUri == nil {
		return nil
	}
	return w.SipUri
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) GetSipUriDynamicVariable() *SipUriDynamicVariableTransferDestination {
	if w == nil {
		return nil
	}
	if w.SipUriDynamicVariable == nil {
		return nil
	}
	return w.SipUriDynamicVariable
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", w)
	}
	switch unmarshaler.Type {
	case "phone":
		value := new(PhoneNumberTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneNumberDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.PhoneDynamicVariable = value
	case "sip_uri":
		value := new(SipUriTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.SipUri = value
	case "sip_uri_dynamic_variable":
		value := new(SipUriDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.SipUriDynamicVariable = value
	}
	return nil
}

func (w WorkflowPhoneNumberNodeModelInputTransferDestination) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(w.Phone, "type", "phone")
	}
	if w.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(w.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	if w.SipUri != nil {
		return internal.MarshalJSONWithDiscriminant(w.SipUri, "type", "sip_uri")
	}
	if w.SipUriDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(w.SipUriDynamicVariable, "type", "sip_uri_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowPhoneNumberNodeModelInputTransferDestinationVisitor interface {
	VisitPhone(*PhoneNumberTransferDestination) error
	VisitPhoneDynamicVariable(*PhoneNumberDynamicVariableTransferDestination) error
	VisitSipUri(*SipUriTransferDestination) error
	VisitSipUriDynamicVariable(*SipUriDynamicVariableTransferDestination) error
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) Accept(visitor WorkflowPhoneNumberNodeModelInputTransferDestinationVisitor) error {
	if w.Phone != nil {
		return visitor.VisitPhone(w.Phone)
	}
	if w.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(w.PhoneDynamicVariable)
	}
	if w.SipUri != nil {
		return visitor.VisitSipUri(w.SipUri)
	}
	if w.SipUriDynamicVariable != nil {
		return visitor.VisitSipUriDynamicVariable(w.SipUriDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowPhoneNumberNodeModelInputTransferDestination) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Phone != nil {
		fields = append(fields, "phone")
	}
	if w.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
	}
	if w.SipUri != nil {
		fields = append(fields, "sip_uri")
	}
	if w.SipUriDynamicVariable != nil {
		fields = append(fields, "sip_uri_dynamic_variable")
	}
	if len(fields) == 0 {
		if w.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Type)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Type != "" {
		field := fields[0]
		if w.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Type,
				w,
			)
		}
	}
	return nil
}

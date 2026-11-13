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

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) GetPhone() *PhoneNumberTransferDestination {
	if workflowPhoneNumberNodeModelInputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.Phone == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelInputTransferDestination.Phone
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) GetPhoneDynamicVariable() *PhoneNumberDynamicVariableTransferDestination {
	if workflowPhoneNumberNodeModelInputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) GetSipUri() *SipUriTransferDestination {
	if workflowPhoneNumberNodeModelInputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUri == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelInputTransferDestination.SipUri
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) GetSipUriDynamicVariable() *SipUriDynamicVariableTransferDestination {
	if workflowPhoneNumberNodeModelInputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) GetType() string {
	if workflowPhoneNumberNodeModelInputTransferDestination == nil {
		return ""
	}
	return workflowPhoneNumberNodeModelInputTransferDestination.Type
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	workflowPhoneNumberNodeModelInputTransferDestination.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", workflowPhoneNumberNodeModelInputTransferDestination)
	}
	switch unmarshaler.Type {
	case "phone":
		value := new(PhoneNumberTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelInputTransferDestination.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneNumberDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable = value
	case "sip_uri":
		value := new(SipUriTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelInputTransferDestination.SipUri = value
	case "sip_uri_dynamic_variable":
		value := new(SipUriDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable = value
	}
	return nil
}

func (workflowPhoneNumberNodeModelInputTransferDestination WorkflowPhoneNumberNodeModelInputTransferDestination) MarshalJSON() ([]byte, error) {
	if err := workflowPhoneNumberNodeModelInputTransferDestination.validate(); err != nil {
		return nil, err
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelInputTransferDestination.Phone, "type", "phone")
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUri != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelInputTransferDestination.SipUri, "type", "sip_uri")
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable, "type", "sip_uri_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", workflowPhoneNumberNodeModelInputTransferDestination)
}

type WorkflowPhoneNumberNodeModelInputTransferDestinationVisitor interface {
	VisitPhone(*PhoneNumberTransferDestination) error
	VisitPhoneDynamicVariable(*PhoneNumberDynamicVariableTransferDestination) error
	VisitSipUri(*SipUriTransferDestination) error
	VisitSipUriDynamicVariable(*SipUriDynamicVariableTransferDestination) error
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) Accept(visitor WorkflowPhoneNumberNodeModelInputTransferDestinationVisitor) error {
	if workflowPhoneNumberNodeModelInputTransferDestination.Phone != nil {
		return visitor.VisitPhone(workflowPhoneNumberNodeModelInputTransferDestination.Phone)
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable)
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUri != nil {
		return visitor.VisitSipUri(workflowPhoneNumberNodeModelInputTransferDestination.SipUri)
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable != nil {
		return visitor.VisitSipUriDynamicVariable(workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", workflowPhoneNumberNodeModelInputTransferDestination)
}

func (workflowPhoneNumberNodeModelInputTransferDestination *WorkflowPhoneNumberNodeModelInputTransferDestination) validate() error {
	if workflowPhoneNumberNodeModelInputTransferDestination == nil {
		return fmt.Errorf("type %T is nil", workflowPhoneNumberNodeModelInputTransferDestination)
	}
	var fields []string
	if workflowPhoneNumberNodeModelInputTransferDestination.Phone != nil {
		fields = append(fields, "phone")
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUri != nil {
		fields = append(fields, "sip_uri")
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.SipUriDynamicVariable != nil {
		fields = append(fields, "sip_uri_dynamic_variable")
	}
	if len(fields) == 0 {
		if workflowPhoneNumberNodeModelInputTransferDestination.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", workflowPhoneNumberNodeModelInputTransferDestination, workflowPhoneNumberNodeModelInputTransferDestination.Type)
		}
		return fmt.Errorf("type %T is empty", workflowPhoneNumberNodeModelInputTransferDestination)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", workflowPhoneNumberNodeModelInputTransferDestination, fields)
	}
	if workflowPhoneNumberNodeModelInputTransferDestination.Type != "" {
		field := fields[0]
		if workflowPhoneNumberNodeModelInputTransferDestination.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				workflowPhoneNumberNodeModelInputTransferDestination,
				workflowPhoneNumberNodeModelInputTransferDestination.Type,
				workflowPhoneNumberNodeModelInputTransferDestination,
			)
		}
	}
	return nil
}

package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowPhoneNumberNodeModelOutputTransferDestination struct {
	Type                  string
	Phone                 *PhoneNumberTransferDestination
	PhoneDynamicVariable  *PhoneNumberDynamicVariableTransferDestination
	SipUri                *SipUriTransferDestination
	SipUriDynamicVariable *SipUriDynamicVariableTransferDestination
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) GetPhone() *PhoneNumberTransferDestination {
	if workflowPhoneNumberNodeModelOutputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.Phone == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelOutputTransferDestination.Phone
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) GetPhoneDynamicVariable() *PhoneNumberDynamicVariableTransferDestination {
	if workflowPhoneNumberNodeModelOutputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) GetSipUri() *SipUriTransferDestination {
	if workflowPhoneNumberNodeModelOutputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUri == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelOutputTransferDestination.SipUri
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) GetSipUriDynamicVariable() *SipUriDynamicVariableTransferDestination {
	if workflowPhoneNumberNodeModelOutputTransferDestination == nil {
		return nil
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable == nil {
		return nil
	}
	return workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) GetType() string {
	if workflowPhoneNumberNodeModelOutputTransferDestination == nil {
		return ""
	}
	return workflowPhoneNumberNodeModelOutputTransferDestination.Type
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	workflowPhoneNumberNodeModelOutputTransferDestination.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", workflowPhoneNumberNodeModelOutputTransferDestination)
	}
	switch unmarshaler.Type {
	case "phone":
		value := new(PhoneNumberTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelOutputTransferDestination.Phone = value
	case "phone_dynamic_variable":
		value := new(PhoneNumberDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable = value
	case "sip_uri":
		value := new(SipUriTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelOutputTransferDestination.SipUri = value
	case "sip_uri_dynamic_variable":
		value := new(SipUriDynamicVariableTransferDestination)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable = value
	}
	return nil
}

func (workflowPhoneNumberNodeModelOutputTransferDestination WorkflowPhoneNumberNodeModelOutputTransferDestination) MarshalJSON() ([]byte, error) {
	if err := workflowPhoneNumberNodeModelOutputTransferDestination.validate(); err != nil {
		return nil, err
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.Phone != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelOutputTransferDestination.Phone, "type", "phone")
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable, "type", "phone_dynamic_variable")
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUri != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelOutputTransferDestination.SipUri, "type", "sip_uri")
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable, "type", "sip_uri_dynamic_variable")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", workflowPhoneNumberNodeModelOutputTransferDestination)
}

type WorkflowPhoneNumberNodeModelOutputTransferDestinationVisitor interface {
	VisitPhone(*PhoneNumberTransferDestination) error
	VisitPhoneDynamicVariable(*PhoneNumberDynamicVariableTransferDestination) error
	VisitSipUri(*SipUriTransferDestination) error
	VisitSipUriDynamicVariable(*SipUriDynamicVariableTransferDestination) error
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) Accept(visitor WorkflowPhoneNumberNodeModelOutputTransferDestinationVisitor) error {
	if workflowPhoneNumberNodeModelOutputTransferDestination.Phone != nil {
		return visitor.VisitPhone(workflowPhoneNumberNodeModelOutputTransferDestination.Phone)
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable != nil {
		return visitor.VisitPhoneDynamicVariable(workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable)
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUri != nil {
		return visitor.VisitSipUri(workflowPhoneNumberNodeModelOutputTransferDestination.SipUri)
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable != nil {
		return visitor.VisitSipUriDynamicVariable(workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", workflowPhoneNumberNodeModelOutputTransferDestination)
}

func (workflowPhoneNumberNodeModelOutputTransferDestination *WorkflowPhoneNumberNodeModelOutputTransferDestination) validate() error {
	if workflowPhoneNumberNodeModelOutputTransferDestination == nil {
		return fmt.Errorf("type %T is nil", workflowPhoneNumberNodeModelOutputTransferDestination)
	}
	var fields []string
	if workflowPhoneNumberNodeModelOutputTransferDestination.Phone != nil {
		fields = append(fields, "phone")
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.PhoneDynamicVariable != nil {
		fields = append(fields, "phone_dynamic_variable")
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUri != nil {
		fields = append(fields, "sip_uri")
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.SipUriDynamicVariable != nil {
		fields = append(fields, "sip_uri_dynamic_variable")
	}
	if len(fields) == 0 {
		if workflowPhoneNumberNodeModelOutputTransferDestination.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", workflowPhoneNumberNodeModelOutputTransferDestination, workflowPhoneNumberNodeModelOutputTransferDestination.Type)
		}
		return fmt.Errorf("type %T is empty", workflowPhoneNumberNodeModelOutputTransferDestination)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", workflowPhoneNumberNodeModelOutputTransferDestination, fields)
	}
	if workflowPhoneNumberNodeModelOutputTransferDestination.Type != "" {
		field := fields[0]
		if workflowPhoneNumberNodeModelOutputTransferDestination.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				workflowPhoneNumberNodeModelOutputTransferDestination,
				workflowPhoneNumberNodeModelOutputTransferDestination.Type,
				workflowPhoneNumberNodeModelOutputTransferDestination,
			)
		}
	}
	return nil
}

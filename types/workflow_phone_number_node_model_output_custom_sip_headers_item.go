package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem struct {
	Type    string
	Dynamic *CustomSipHeaderWithDynamicVariable
	Static  *CustomSipHeader
}

func (w *WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) GetDynamic() *CustomSipHeaderWithDynamicVariable {
	if w == nil {
		return nil
	}
	if w.Dynamic == nil {
		return nil
	}
	return w.Dynamic
}

func (w *WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) GetStatic() *CustomSipHeader {
	if w == nil {
		return nil
	}
	if w.Static == nil {
		return nil
	}
	return w.Static
}

func (w *WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) UnmarshalJSON(data []byte) error {
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
	case "dynamic":
		value := new(CustomSipHeaderWithDynamicVariable)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Dynamic = value
	case "static":
		value := new(CustomSipHeader)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Static = value
	}
	return nil
}

func (w WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Dynamic != nil {
		return internal.MarshalJSONWithDiscriminant(w.Dynamic, "type", "dynamic")
	}
	if w.Static != nil {
		return internal.MarshalJSONWithDiscriminant(w.Static, "type", "static")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItemVisitor interface {
	VisitDynamic(*CustomSipHeaderWithDynamicVariable) error
	VisitStatic(*CustomSipHeader) error
}

func (w *WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) Accept(visitor WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItemVisitor) error {
	if w.Dynamic != nil {
		return visitor.VisitDynamic(w.Dynamic)
	}
	if w.Static != nil {
		return visitor.VisitStatic(w.Static)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowPhoneNumberNodeModelOutputCustomSipHeadersItem) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Dynamic != nil {
		fields = append(fields, "dynamic")
	}
	if w.Static != nil {
		fields = append(fields, "static")
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

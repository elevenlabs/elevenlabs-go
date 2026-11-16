package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem struct {
	Type    string
	Dynamic *CustomSipHeaderWithDynamicVariable
	Static  *CustomSipHeader
}

func (w *WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) GetDynamic() *CustomSipHeaderWithDynamicVariable {
	if w == nil {
		return nil
	}
	if w.Dynamic == nil {
		return nil
	}
	return w.Dynamic
}

func (w *WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) GetStatic() *CustomSipHeader {
	if w == nil {
		return nil
	}
	if w.Static == nil {
		return nil
	}
	return w.Static
}

func (w *WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) UnmarshalJSON(data []byte) error {
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

func (w WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) MarshalJSON() ([]byte, error) {
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

type WorkflowPhoneNumberNodeModelInputCustomSipHeadersItemVisitor interface {
	VisitDynamic(*CustomSipHeaderWithDynamicVariable) error
	VisitStatic(*CustomSipHeader) error
}

func (w *WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) Accept(visitor WorkflowPhoneNumberNodeModelInputCustomSipHeadersItemVisitor) error {
	if w.Dynamic != nil {
		return visitor.VisitDynamic(w.Dynamic)
	}
	if w.Static != nil {
		return visitor.VisitStatic(w.Static)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem) validate() error {
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

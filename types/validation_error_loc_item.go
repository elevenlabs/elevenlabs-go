package types

import (
	"encoding/json"
	"fmt"
)

type ValidationErrorLocItem struct {
	String  string
	Integer int
	typ     string
}

func (v *ValidationErrorLocItem) GetString() string {
	if v == nil {
		return ""
	}
	return v.String
}

func (v *ValidationErrorLocItem) GetInteger() int {
	if v == nil {
		return 0
	}
	return v.Integer
}

func (v *ValidationErrorLocItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		v.typ = "String"
		v.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		v.typ = "Integer"
		v.Integer = valueInteger
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, v)
}

func (v ValidationErrorLocItem) MarshalJSON() ([]byte, error) {
	if v.typ == "String" || v.String != "" {
		return json.Marshal(v.String)
	}
	if v.typ == "Integer" || v.Integer != 0 {
		return json.Marshal(v.Integer)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type ValidationErrorLocItemVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
}

func (v *ValidationErrorLocItem) Accept(visitor ValidationErrorLocItemVisitor) error {
	if v.typ == "String" || v.String != "" {
		return visitor.VisitString(v.String)
	}
	if v.typ == "Integer" || v.Integer != 0 {
		return visitor.VisitInteger(v.Integer)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}

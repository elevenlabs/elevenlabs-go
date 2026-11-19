package types

import (
	"encoding/json"
	"fmt"
)

// One element of an array-valued TemplateRunInput. Same members as TemplateRunInput; nested arrays are untyped (stopgap for a recursive-union generator bug).
type TemplateRunInputElement struct {
	String                 string
	Boolean                bool
	Integer                int
	Double                 float64
	TemplateInputReference *TemplateInputReference
	UnknownList            []any
	typ                    string
}

func (t *TemplateRunInputElement) GetString() string {
	if t == nil {
		return ""
	}
	return t.String
}

func (t *TemplateRunInputElement) GetBoolean() bool {
	if t == nil {
		return false
	}
	return t.Boolean
}

func (t *TemplateRunInputElement) GetInteger() int {
	if t == nil {
		return 0
	}
	return t.Integer
}

func (t *TemplateRunInputElement) GetDouble() float64 {
	if t == nil {
		return 0
	}
	return t.Double
}

func (t *TemplateRunInputElement) GetTemplateInputReference() *TemplateInputReference {
	if t == nil {
		return nil
	}
	if t.TemplateInputReference == nil {
		return nil
	}
	return t.TemplateInputReference
}

func (t *TemplateRunInputElement) GetUnknownList() []any {
	if t == nil {
		return nil
	}
	return t.UnknownList
}

func (t *TemplateRunInputElement) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		t.typ = "String"
		t.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		t.typ = "Boolean"
		t.Boolean = valueBoolean
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		t.typ = "Integer"
		t.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		t.typ = "Double"
		t.Double = valueDouble
		return nil
	}
	var valueTemplateInputReference *TemplateInputReference
	if err := json.Unmarshal(data, &valueTemplateInputReference); err == nil {
		t.typ = "TemplateInputReference"
		t.TemplateInputReference = valueTemplateInputReference
		return nil
	}
	var valueUnknownList []any
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		t.typ = "UnknownList"
		t.UnknownList = valueUnknownList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TemplateRunInputElement) MarshalJSON() ([]byte, error) {
	if t.typ == "String" || t.String != "" {
		return json.Marshal(t.String)
	}
	if t.typ == "Boolean" || t.Boolean != false {
		return json.Marshal(t.Boolean)
	}
	if t.typ == "Integer" || t.Integer != 0 {
		return json.Marshal(t.Integer)
	}
	if t.typ == "Double" || t.Double != 0 {
		return json.Marshal(t.Double)
	}
	if t.typ == "TemplateInputReference" || t.TemplateInputReference != nil {
		return json.Marshal(t.TemplateInputReference)
	}
	if t.typ == "UnknownList" || t.UnknownList != nil {
		return json.Marshal(t.UnknownList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TemplateRunInputElementVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitTemplateInputReference(*TemplateInputReference) error
	VisitUnknownList([]any) error
}

func (t *TemplateRunInputElement) Accept(visitor TemplateRunInputElementVisitor) error {
	if t.typ == "String" || t.String != "" {
		return visitor.VisitString(t.String)
	}
	if t.typ == "Boolean" || t.Boolean != false {
		return visitor.VisitBoolean(t.Boolean)
	}
	if t.typ == "Integer" || t.Integer != 0 {
		return visitor.VisitInteger(t.Integer)
	}
	if t.typ == "Double" || t.Double != 0 {
		return visitor.VisitDouble(t.Double)
	}
	if t.typ == "TemplateInputReference" || t.TemplateInputReference != nil {
		return visitor.VisitTemplateInputReference(t.TemplateInputReference)
	}
	if t.typ == "UnknownList" || t.UnknownList != nil {
		return visitor.VisitUnknownList(t.UnknownList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

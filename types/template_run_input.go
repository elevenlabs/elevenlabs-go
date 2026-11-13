package types

import (
	"encoding/json"
	"fmt"
)

// One value bound to a template input port, in the form its `content_schema`
// calls for. A `string` port takes the text itself; `number`, `integer` and
// `boolean` ports take a JSON value of that type. Anything stored elsewhere is
// a reference discriminated on `type`: a prior `generation`, an uploaded
// `asset`, a `voice`, or media passed `inline_base64`. An `array` port takes a
// JSON array with one value per element, admissible for the schema's `items`.
// `object` ports cannot be bound through this API yet.
//
// No constraints beyond the wire shape live here: what a port accepts (an
// `enum`, a length) is stated by its `content_schema`, and the run path
// validates against that same schema.
type TemplateRunInput struct {
	String                      string
	Boolean                     bool
	Integer                     int
	Double                      float64
	TemplateInputReference      *TemplateInputReference
	TemplateRunInputElementList []*TemplateRunInputElement
	typ                         string
}

func (t *TemplateRunInput) GetString() string {
	if t == nil {
		return ""
	}
	return t.String
}

func (t *TemplateRunInput) GetBoolean() bool {
	if t == nil {
		return false
	}
	return t.Boolean
}

func (t *TemplateRunInput) GetInteger() int {
	if t == nil {
		return 0
	}
	return t.Integer
}

func (t *TemplateRunInput) GetDouble() float64 {
	if t == nil {
		return 0
	}
	return t.Double
}

func (t *TemplateRunInput) GetTemplateInputReference() *TemplateInputReference {
	if t == nil {
		return nil
	}
	if t.TemplateInputReference == nil {
		return nil
	}
	return t.TemplateInputReference
}

func (t *TemplateRunInput) GetTemplateRunInputElementList() []*TemplateRunInputElement {
	if t == nil {
		return nil
	}
	return t.TemplateRunInputElementList
}

func (t *TemplateRunInput) UnmarshalJSON(data []byte) error {
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
	var valueTemplateRunInputElementList []*TemplateRunInputElement
	if err := json.Unmarshal(data, &valueTemplateRunInputElementList); err == nil {
		t.typ = "TemplateRunInputElementList"
		t.TemplateRunInputElementList = valueTemplateRunInputElementList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TemplateRunInput) MarshalJSON() ([]byte, error) {
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
	if t.typ == "TemplateRunInputElementList" || t.TemplateRunInputElementList != nil {
		return json.Marshal(t.TemplateRunInputElementList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TemplateRunInputVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitTemplateInputReference(*TemplateInputReference) error
	VisitTemplateRunInputElementList([]*TemplateRunInputElement) error
}

func (t *TemplateRunInput) Accept(visitor TemplateRunInputVisitor) error {
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
	if t.typ == "TemplateRunInputElementList" || t.TemplateRunInputElementList != nil {
		return visitor.VisitTemplateRunInputElementList(t.TemplateRunInputElementList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

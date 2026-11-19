package types

import (
	"encoding/json"
	"fmt"
)

// A constant value to use for this property. Mutually exclusive with description, dynamic_variable, is_system_provided, and is_omitted.
type LiteralJsonSchemaPropertyConstantValue struct {
	String  string
	Integer int
	Double  float64
	Boolean bool
	typ     string
}

func (l *LiteralJsonSchemaPropertyConstantValue) GetString() string {
	if l == nil {
		return ""
	}
	return l.String
}

func (l *LiteralJsonSchemaPropertyConstantValue) GetInteger() int {
	if l == nil {
		return 0
	}
	return l.Integer
}

func (l *LiteralJsonSchemaPropertyConstantValue) GetDouble() float64 {
	if l == nil {
		return 0
	}
	return l.Double
}

func (l *LiteralJsonSchemaPropertyConstantValue) GetBoolean() bool {
	if l == nil {
		return false
	}
	return l.Boolean
}

func (l *LiteralJsonSchemaPropertyConstantValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		l.typ = "String"
		l.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		l.typ = "Integer"
		l.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		l.typ = "Double"
		l.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		l.typ = "Boolean"
		l.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LiteralJsonSchemaPropertyConstantValue) MarshalJSON() ([]byte, error) {
	if l.typ == "String" || l.String != "" {
		return json.Marshal(l.String)
	}
	if l.typ == "Integer" || l.Integer != 0 {
		return json.Marshal(l.Integer)
	}
	if l.typ == "Double" || l.Double != 0 {
		return json.Marshal(l.Double)
	}
	if l.typ == "Boolean" || l.Boolean != false {
		return json.Marshal(l.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LiteralJsonSchemaPropertyConstantValueVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
}

func (l *LiteralJsonSchemaPropertyConstantValue) Accept(visitor LiteralJsonSchemaPropertyConstantValueVisitor) error {
	if l.typ == "String" || l.String != "" {
		return visitor.VisitString(l.String)
	}
	if l.typ == "Integer" || l.Integer != 0 {
		return visitor.VisitInteger(l.Integer)
	}
	if l.typ == "Double" || l.Double != 0 {
		return visitor.VisitDouble(l.Double)
	}
	if l.typ == "Boolean" || l.Boolean != false {
		return visitor.VisitBoolean(l.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}

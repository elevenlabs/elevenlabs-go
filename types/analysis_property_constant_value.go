package types

import (
	"encoding/json"
	"fmt"
)

// A constant value to use for this property. Mutually exclusive with description, dynamic_variable, is_system_provided, and is_omitted.
type AnalysisPropertyConstantValue struct {
	String  string
	Integer int
	Double  float64
	Boolean bool
	typ     string
}

func (a *AnalysisPropertyConstantValue) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *AnalysisPropertyConstantValue) GetInteger() int {
	if a == nil {
		return 0
	}
	return a.Integer
}

func (a *AnalysisPropertyConstantValue) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *AnalysisPropertyConstantValue) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *AnalysisPropertyConstantValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		a.typ = "Integer"
		a.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AnalysisPropertyConstantValue) MarshalJSON() ([]byte, error) {
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "Integer" || a.Integer != 0 {
		return json.Marshal(a.Integer)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AnalysisPropertyConstantValueVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
}

func (a *AnalysisPropertyConstantValue) Accept(visitor AnalysisPropertyConstantValueVisitor) error {
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "Integer" || a.Integer != 0 {
		return visitor.VisitInteger(a.Integer)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

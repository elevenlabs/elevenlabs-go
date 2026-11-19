package types

import (
	"encoding/json"
	"fmt"
)

// A constant value to use for this property. Mutually exclusive with description, dynamic_variable, is_system_provided, and is_omitted.
type DataCollectionValueSchemaConstantValue struct {
	String  string
	Integer int
	Double  float64
	Boolean bool
	typ     string
}

func (d *DataCollectionValueSchemaConstantValue) GetString() string {
	if d == nil {
		return ""
	}
	return d.String
}

func (d *DataCollectionValueSchemaConstantValue) GetInteger() int {
	if d == nil {
		return 0
	}
	return d.Integer
}

func (d *DataCollectionValueSchemaConstantValue) GetDouble() float64 {
	if d == nil {
		return 0
	}
	return d.Double
}

func (d *DataCollectionValueSchemaConstantValue) GetBoolean() bool {
	if d == nil {
		return false
	}
	return d.Boolean
}

func (d *DataCollectionValueSchemaConstantValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		d.typ = "String"
		d.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		d.typ = "Integer"
		d.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		d.typ = "Double"
		d.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		d.typ = "Boolean"
		d.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DataCollectionValueSchemaConstantValue) MarshalJSON() ([]byte, error) {
	if d.typ == "String" || d.String != "" {
		return json.Marshal(d.String)
	}
	if d.typ == "Integer" || d.Integer != 0 {
		return json.Marshal(d.Integer)
	}
	if d.typ == "Double" || d.Double != 0 {
		return json.Marshal(d.Double)
	}
	if d.typ == "Boolean" || d.Boolean != false {
		return json.Marshal(d.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", d)
}

type DataCollectionValueSchemaConstantValueVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
}

func (d *DataCollectionValueSchemaConstantValue) Accept(visitor DataCollectionValueSchemaConstantValueVisitor) error {
	if d.typ == "String" || d.String != "" {
		return visitor.VisitString(d.String)
	}
	if d.typ == "Integer" || d.Integer != 0 {
		return visitor.VisitInteger(d.Integer)
	}
	if d.typ == "Double" || d.Double != 0 {
		return visitor.VisitDouble(d.Double)
	}
	if d.typ == "Boolean" || d.Boolean != false {
		return visitor.VisitBoolean(d.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", d)
}

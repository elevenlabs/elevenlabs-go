package types

import (
	"encoding/json"
	"fmt"
)

// The constant value to use
type ConstantSchemaOverrideConstantValue struct {
	String           string
	Integer          int
	Double           float64
	Boolean          bool
	UnknownList      []any
	StringUnknownMap map[string]any
	typ              string
}

func (c *ConstantSchemaOverrideConstantValue) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *ConstantSchemaOverrideConstantValue) GetInteger() int {
	if c == nil {
		return 0
	}
	return c.Integer
}

func (c *ConstantSchemaOverrideConstantValue) GetDouble() float64 {
	if c == nil {
		return 0
	}
	return c.Double
}

func (c *ConstantSchemaOverrideConstantValue) GetBoolean() bool {
	if c == nil {
		return false
	}
	return c.Boolean
}

func (c *ConstantSchemaOverrideConstantValue) GetUnknownList() []any {
	if c == nil {
		return nil
	}
	return c.UnknownList
}

func (c *ConstantSchemaOverrideConstantValue) GetStringUnknownMap() map[string]any {
	if c == nil {
		return nil
	}
	return c.StringUnknownMap
}

func (c *ConstantSchemaOverrideConstantValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		c.typ = "Integer"
		c.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typ = "Double"
		c.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		c.typ = "Boolean"
		c.Boolean = valueBoolean
		return nil
	}
	var valueUnknownList []any
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		c.typ = "UnknownList"
		c.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]any
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		c.typ = "StringUnknownMap"
		c.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConstantSchemaOverrideConstantValue) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "Integer" || c.Integer != 0 {
		return json.Marshal(c.Integer)
	}
	if c.typ == "Double" || c.Double != 0 {
		return json.Marshal(c.Double)
	}
	if c.typ == "Boolean" || c.Boolean != false {
		return json.Marshal(c.Boolean)
	}
	if c.typ == "UnknownList" || c.UnknownList != nil {
		return json.Marshal(c.UnknownList)
	}
	if c.typ == "StringUnknownMap" || c.StringUnknownMap != nil {
		return json.Marshal(c.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConstantSchemaOverrideConstantValueVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
	VisitUnknownList([]any) error
	VisitStringUnknownMap(map[string]any) error
}

func (c *ConstantSchemaOverrideConstantValue) Accept(visitor ConstantSchemaOverrideConstantValueVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "Integer" || c.Integer != 0 {
		return visitor.VisitInteger(c.Integer)
	}
	if c.typ == "Double" || c.Double != 0 {
		return visitor.VisitDouble(c.Double)
	}
	if c.typ == "Boolean" || c.Boolean != false {
		return visitor.VisitBoolean(c.Boolean)
	}
	if c.typ == "UnknownList" || c.UnknownList != nil {
		return visitor.VisitUnknownList(c.UnknownList)
	}
	if c.typ == "StringUnknownMap" || c.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(c.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type ColumnFilterValuesItem struct {
	String   string
	Integer  int
	Double   float64
	DateTime time.Time
	Boolean  bool
	typ      string
}

func (c *ColumnFilterValuesItem) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *ColumnFilterValuesItem) GetInteger() int {
	if c == nil {
		return 0
	}
	return c.Integer
}

func (c *ColumnFilterValuesItem) GetDouble() float64 {
	if c == nil {
		return 0
	}
	return c.Double
}

func (c *ColumnFilterValuesItem) GetDateTime() time.Time {
	if c == nil {
		return (time.Time{})
	}
	return c.DateTime
}

func (c *ColumnFilterValuesItem) GetBoolean() bool {
	if c == nil {
		return false
	}
	return c.Boolean
}

func (c *ColumnFilterValuesItem) UnmarshalJSON(data []byte) error {
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
	var valueDateTime time.Time
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		c.typ = "DateTime"
		c.DateTime = valueDateTime
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		c.typ = "Boolean"
		c.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ColumnFilterValuesItem) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "Integer" || c.Integer != 0 {
		return json.Marshal(c.Integer)
	}
	if c.typ == "Double" || c.Double != 0 {
		return json.Marshal(c.Double)
	}
	if c.typ == "DateTime" || c.DateTime != (time.Time{}) {
		return json.Marshal(c.DateTime)
	}
	if c.typ == "Boolean" || c.Boolean != false {
		return json.Marshal(c.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ColumnFilterValuesItemVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitDateTime(time.Time) error
	VisitBoolean(bool) error
}

func (c *ColumnFilterValuesItem) Accept(visitor ColumnFilterValuesItemVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "Integer" || c.Integer != 0 {
		return visitor.VisitInteger(c.Integer)
	}
	if c.typ == "Double" || c.Double != 0 {
		return visitor.VisitDouble(c.Double)
	}
	if c.typ == "DateTime" || c.DateTime != (time.Time{}) {
		return visitor.VisitDateTime(c.DateTime)
	}
	if c.typ == "Boolean" || c.Boolean != false {
		return visitor.VisitBoolean(c.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

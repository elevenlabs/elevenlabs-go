package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type MetronomeUsageColumnFilterValuesItem struct {
	String   string
	Integer  int
	Double   float64
	DateTime time.Time
	Boolean  bool
	typ      string
}

func (m *MetronomeUsageColumnFilterValuesItem) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *MetronomeUsageColumnFilterValuesItem) GetInteger() int {
	if m == nil {
		return 0
	}
	return m.Integer
}

func (m *MetronomeUsageColumnFilterValuesItem) GetDouble() float64 {
	if m == nil {
		return 0
	}
	return m.Double
}

func (m *MetronomeUsageColumnFilterValuesItem) GetDateTime() time.Time {
	if m == nil {
		return (time.Time{})
	}
	return m.DateTime
}

func (m *MetronomeUsageColumnFilterValuesItem) GetBoolean() bool {
	if m == nil {
		return false
	}
	return m.Boolean
}

func (m *MetronomeUsageColumnFilterValuesItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		m.typ = "Integer"
		m.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		m.typ = "Double"
		m.Double = valueDouble
		return nil
	}
	var valueDateTime time.Time
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		m.typ = "DateTime"
		m.DateTime = valueDateTime
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		m.typ = "Boolean"
		m.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MetronomeUsageColumnFilterValuesItem) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "Integer" || m.Integer != 0 {
		return json.Marshal(m.Integer)
	}
	if m.typ == "Double" || m.Double != 0 {
		return json.Marshal(m.Double)
	}
	if m.typ == "DateTime" || m.DateTime != (time.Time{}) {
		return json.Marshal(m.DateTime)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return json.Marshal(m.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MetronomeUsageColumnFilterValuesItemVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitDateTime(time.Time) error
	VisitBoolean(bool) error
}

func (m *MetronomeUsageColumnFilterValuesItem) Accept(visitor MetronomeUsageColumnFilterValuesItemVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "Integer" || m.Integer != 0 {
		return visitor.VisitInteger(m.Integer)
	}
	if m.typ == "Double" || m.Double != 0 {
		return visitor.VisitDouble(m.Double)
	}
	if m.typ == "DateTime" || m.DateTime != (time.Time{}) {
		return visitor.VisitDateTime(m.DateTime)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return visitor.VisitBoolean(m.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

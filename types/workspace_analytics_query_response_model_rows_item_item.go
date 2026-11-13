package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type WorkspaceAnalyticsQueryResponseModelRowsItemItem struct {
	String   string
	Integer  int
	Double   float64
	Boolean  bool
	DateTime time.Time
	typ      string
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) GetString() string {
	if w == nil {
		return ""
	}
	return w.String
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) GetInteger() int {
	if w == nil {
		return 0
	}
	return w.Integer
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) GetDouble() float64 {
	if w == nil {
		return 0
	}
	return w.Double
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) GetBoolean() bool {
	if w == nil {
		return false
	}
	return w.Boolean
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) GetDateTime() time.Time {
	if w == nil {
		return (time.Time{})
	}
	return w.DateTime
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		w.typ = "String"
		w.String = valueString
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		w.typ = "Integer"
		w.Integer = valueInteger
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		w.typ = "Double"
		w.Double = valueDouble
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		w.typ = "Boolean"
		w.Boolean = valueBoolean
		return nil
	}
	var valueDateTime time.Time
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		w.typ = "DateTime"
		w.DateTime = valueDateTime
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, w)
}

func (w WorkspaceAnalyticsQueryResponseModelRowsItemItem) MarshalJSON() ([]byte, error) {
	if w.typ == "String" || w.String != "" {
		return json.Marshal(w.String)
	}
	if w.typ == "Integer" || w.Integer != 0 {
		return json.Marshal(w.Integer)
	}
	if w.typ == "Double" || w.Double != 0 {
		return json.Marshal(w.Double)
	}
	if w.typ == "Boolean" || w.Boolean != false {
		return json.Marshal(w.Boolean)
	}
	if w.typ == "DateTime" || w.DateTime != (time.Time{}) {
		return json.Marshal(w.DateTime)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", w)
}

type WorkspaceAnalyticsQueryResponseModelRowsItemItemVisitor interface {
	VisitString(string) error
	VisitInteger(int) error
	VisitDouble(float64) error
	VisitBoolean(bool) error
	VisitDateTime(time.Time) error
}

func (w *WorkspaceAnalyticsQueryResponseModelRowsItemItem) Accept(visitor WorkspaceAnalyticsQueryResponseModelRowsItemItemVisitor) error {
	if w.typ == "String" || w.String != "" {
		return visitor.VisitString(w.String)
	}
	if w.typ == "Integer" || w.Integer != 0 {
		return visitor.VisitInteger(w.Integer)
	}
	if w.typ == "Double" || w.Double != 0 {
		return visitor.VisitDouble(w.Double)
	}
	if w.typ == "Boolean" || w.Boolean != false {
		return visitor.VisitBoolean(w.Boolean)
	}
	if w.typ == "DateTime" || w.DateTime != (time.Time{}) {
		return visitor.VisitDateTime(w.DateTime)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", w)
}

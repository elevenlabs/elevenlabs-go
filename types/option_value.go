package types

import (
	"encoding/json"
	"fmt"
)

type OptionValue struct {
	Double  float64
	String  string
	Boolean bool
	typ     string
}

func (o *OptionValue) GetDouble() float64 {
	if o == nil {
		return 0
	}
	return o.Double
}

func (o *OptionValue) GetString() string {
	if o == nil {
		return ""
	}
	return o.String
}

func (o *OptionValue) GetBoolean() bool {
	if o == nil {
		return false
	}
	return o.Boolean
}

func (o *OptionValue) UnmarshalJSON(data []byte) error {
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		o.typ = "Double"
		o.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		o.typ = "String"
		o.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		o.typ = "Boolean"
		o.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OptionValue) MarshalJSON() ([]byte, error) {
	if o.typ == "Double" || o.Double != 0 {
		return json.Marshal(o.Double)
	}
	if o.typ == "String" || o.String != "" {
		return json.Marshal(o.String)
	}
	if o.typ == "Boolean" || o.Boolean != false {
		return json.Marshal(o.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OptionValueVisitor interface {
	VisitDouble(float64) error
	VisitString(string) error
	VisitBoolean(bool) error
}

func (o *OptionValue) Accept(visitor OptionValueVisitor) error {
	if o.typ == "Double" || o.Double != 0 {
		return visitor.VisitDouble(o.Double)
	}
	if o.typ == "String" || o.String != "" {
		return visitor.VisitString(o.String)
	}
	if o.typ == "Boolean" || o.Boolean != false {
		return visitor.VisitBoolean(o.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}

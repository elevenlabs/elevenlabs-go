package types

import (
	"encoding/json"
	"fmt"
)

type OTelAttributeValueValue struct {
	String  string
	Boolean bool
	Double  float64
	typ     string
}

func (o *OTelAttributeValueValue) GetString() string {
	if o == nil {
		return ""
	}
	return o.String
}

func (o *OTelAttributeValueValue) GetBoolean() bool {
	if o == nil {
		return false
	}
	return o.Boolean
}

func (o *OTelAttributeValueValue) GetDouble() float64 {
	if o == nil {
		return 0
	}
	return o.Double
}

func (o *OTelAttributeValueValue) UnmarshalJSON(data []byte) error {
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
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		o.typ = "Double"
		o.Double = valueDouble
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OTelAttributeValueValue) MarshalJSON() ([]byte, error) {
	if o.typ == "String" || o.String != "" {
		return json.Marshal(o.String)
	}
	if o.typ == "Boolean" || o.Boolean != false {
		return json.Marshal(o.Boolean)
	}
	if o.typ == "Double" || o.Double != 0 {
		return json.Marshal(o.Double)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OTelAttributeValueValueVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
	VisitDouble(float64) error
}

func (o *OTelAttributeValueValue) Accept(visitor OTelAttributeValueValueVisitor) error {
	if o.typ == "String" || o.String != "" {
		return visitor.VisitString(o.String)
	}
	if o.typ == "Boolean" || o.Boolean != false {
		return visitor.VisitBoolean(o.Boolean)
	}
	if o.typ == "Double" || o.Double != 0 {
		return visitor.VisitDouble(o.Double)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}

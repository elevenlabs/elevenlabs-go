package types

import (
	"encoding/json"
	"fmt"
)

// Value of this literal.
type AstNumberNodeInputValue struct {
	Double  float64
	Integer int
	typ     string
}

func (a *AstNumberNodeInputValue) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *AstNumberNodeInputValue) GetInteger() int {
	if a == nil {
		return 0
	}
	return a.Integer
}

func (a *AstNumberNodeInputValue) UnmarshalJSON(data []byte) error {
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		a.typ = "Integer"
		a.Integer = valueInteger
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AstNumberNodeInputValue) MarshalJSON() ([]byte, error) {
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "Integer" || a.Integer != 0 {
		return json.Marshal(a.Integer)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AstNumberNodeInputValueVisitor interface {
	VisitDouble(float64) error
	VisitInteger(int) error
}

func (a *AstNumberNodeInputValue) Accept(visitor AstNumberNodeInputValueVisitor) error {
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "Integer" || a.Integer != 0 {
		return visitor.VisitInteger(a.Integer)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

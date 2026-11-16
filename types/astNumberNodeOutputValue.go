package types

import (
	"encoding/json"
	"fmt"
)

// Value of this literal.
type AstNumberNodeOutputValue struct {
	Double  float64
	Integer int
	typ     string
}

func (astNumberNodeOutputValue *AstNumberNodeOutputValue) GetDouble() float64 {
	if astNumberNodeOutputValue == nil {
		return 0
	}
	return astNumberNodeOutputValue.Double
}

func (astNumberNodeOutputValue *AstNumberNodeOutputValue) GetInteger() int {
	if astNumberNodeOutputValue == nil {
		return 0
	}
	return astNumberNodeOutputValue.Integer
}

func (astNumberNodeOutputValue *AstNumberNodeOutputValue) UnmarshalJSON(data []byte) error {
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		astNumberNodeOutputValue.typ = "Double"
		astNumberNodeOutputValue.Double = valueDouble
		return nil
	}
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		astNumberNodeOutputValue.typ = "Integer"
		astNumberNodeOutputValue.Integer = valueInteger
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, astNumberNodeOutputValue)
}

func (astNumberNodeOutputValue AstNumberNodeOutputValue) MarshalJSON() ([]byte, error) {
	if astNumberNodeOutputValue.typ == "Double" || astNumberNodeOutputValue.Double != 0 {
		return json.Marshal(astNumberNodeOutputValue.Double)
	}
	if astNumberNodeOutputValue.typ == "Integer" || astNumberNodeOutputValue.Integer != 0 {
		return json.Marshal(astNumberNodeOutputValue.Integer)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", astNumberNodeOutputValue)
}

type AstNumberNodeOutputValueVisitor interface {
	VisitDouble(float64) error
	VisitInteger(int) error
}

func (astNumberNodeOutputValue *AstNumberNodeOutputValue) Accept(visitor AstNumberNodeOutputValueVisitor) error {
	if astNumberNodeOutputValue.typ == "Double" || astNumberNodeOutputValue.Double != 0 {
		return visitor.VisitDouble(astNumberNodeOutputValue.Double)
	}
	if astNumberNodeOutputValue.typ == "Integer" || astNumberNodeOutputValue.Integer != 0 {
		return visitor.VisitInteger(astNumberNodeOutputValue.Integer)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", astNumberNodeOutputValue)
}

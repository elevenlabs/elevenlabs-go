package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigInputMaxResponseOutputTokens struct {
	Integer          int
	InfStringLiteral string
	typ              string
}

func (o *OpenAiSessionConfigInputMaxResponseOutputTokens) GetInteger() int {
	if o == nil {
		return 0
	}
	return o.Integer
}

func (o *OpenAiSessionConfigInputMaxResponseOutputTokens) GetInfStringLiteral() string {
	if o == nil {
		return ""
	}
	return o.InfStringLiteral
}

func (o *OpenAiSessionConfigInputMaxResponseOutputTokens) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		o.typ = "Integer"
		o.Integer = valueInteger
		return nil
	}
	var valueInfStringLiteral string
	if err := json.Unmarshal(data, &valueInfStringLiteral); err == nil {
		o.typ = "InfStringLiteral"
		o.InfStringLiteral = valueInfStringLiteral
		if o.InfStringLiteral != "inf" {
			return fmt.Errorf("unexpected value for literal on type %T; expected inf got %v", o, o.InfStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OpenAiSessionConfigInputMaxResponseOutputTokens) MarshalJSON() ([]byte, error) {
	if o.typ == "Integer" || o.Integer != 0 {
		return json.Marshal(o.Integer)
	}
	if o.typ == "InfStringLiteral" || o.InfStringLiteral != "" {
		return json.Marshal(o.InfStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OpenAiSessionConfigInputMaxResponseOutputTokensVisitor interface {
	VisitInteger(int) error
	VisitInfStringLiteral(string) error
}

func (o *OpenAiSessionConfigInputMaxResponseOutputTokens) Accept(visitor OpenAiSessionConfigInputMaxResponseOutputTokensVisitor) error {
	if o.typ == "Integer" || o.Integer != 0 {
		return visitor.VisitInteger(o.Integer)
	}
	if o.typ == "InfStringLiteral" || o.InfStringLiteral != "" {
		return visitor.VisitInfStringLiteral(o.InfStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}

func NewOpenAiSessionConfigInputMaxResponseOutputTokensWithInfStringLiteral() *OpenAiSessionConfigInputMaxResponseOutputTokens {
	return &OpenAiSessionConfigInputMaxResponseOutputTokens{typ: "InfStringLiteral", InfStringLiteral: "inf"}
}

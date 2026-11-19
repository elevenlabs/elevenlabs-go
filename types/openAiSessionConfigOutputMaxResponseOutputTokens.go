package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigOutputMaxResponseOutputTokens struct {
	Integer          int
	InfStringLiteral string
	typ              string
}

func (openAiSessionConfigOutputMaxResponseOutputTokens *OpenAiSessionConfigOutputMaxResponseOutputTokens) GetInteger() int {
	if openAiSessionConfigOutputMaxResponseOutputTokens == nil {
		return 0
	}
	return openAiSessionConfigOutputMaxResponseOutputTokens.Integer
}

func (openAiSessionConfigOutputMaxResponseOutputTokens *OpenAiSessionConfigOutputMaxResponseOutputTokens) GetInfStringLiteral() string {
	if openAiSessionConfigOutputMaxResponseOutputTokens == nil {
		return ""
	}
	return openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral
}

func (openAiSessionConfigOutputMaxResponseOutputTokens *OpenAiSessionConfigOutputMaxResponseOutputTokens) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		openAiSessionConfigOutputMaxResponseOutputTokens.typ = "Integer"
		openAiSessionConfigOutputMaxResponseOutputTokens.Integer = valueInteger
		return nil
	}
	var valueInfStringLiteral string
	if err := json.Unmarshal(data, &valueInfStringLiteral); err == nil {
		openAiSessionConfigOutputMaxResponseOutputTokens.typ = "InfStringLiteral"
		openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral = valueInfStringLiteral
		if openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral != "inf" {
			return fmt.Errorf("unexpected value for literal on type %T; expected inf got %v", openAiSessionConfigOutputMaxResponseOutputTokens, openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, openAiSessionConfigOutputMaxResponseOutputTokens)
}

func (openAiSessionConfigOutputMaxResponseOutputTokens OpenAiSessionConfigOutputMaxResponseOutputTokens) MarshalJSON() ([]byte, error) {
	if openAiSessionConfigOutputMaxResponseOutputTokens.typ == "Integer" || openAiSessionConfigOutputMaxResponseOutputTokens.Integer != 0 {
		return json.Marshal(openAiSessionConfigOutputMaxResponseOutputTokens.Integer)
	}
	if openAiSessionConfigOutputMaxResponseOutputTokens.typ == "InfStringLiteral" || openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral != "" {
		return json.Marshal(openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", openAiSessionConfigOutputMaxResponseOutputTokens)
}

type OpenAiSessionConfigOutputMaxResponseOutputTokensVisitor interface {
	VisitInteger(int) error
	VisitInfStringLiteral(string) error
}

func (openAiSessionConfigOutputMaxResponseOutputTokens *OpenAiSessionConfigOutputMaxResponseOutputTokens) Accept(visitor OpenAiSessionConfigOutputMaxResponseOutputTokensVisitor) error {
	if openAiSessionConfigOutputMaxResponseOutputTokens.typ == "Integer" || openAiSessionConfigOutputMaxResponseOutputTokens.Integer != 0 {
		return visitor.VisitInteger(openAiSessionConfigOutputMaxResponseOutputTokens.Integer)
	}
	if openAiSessionConfigOutputMaxResponseOutputTokens.typ == "InfStringLiteral" || openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral != "" {
		return visitor.VisitInfStringLiteral(openAiSessionConfigOutputMaxResponseOutputTokens.InfStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", openAiSessionConfigOutputMaxResponseOutputTokens)
}

func NewOpenAiSessionConfigOutputMaxResponseOutputTokensWithInfStringLiteral() *OpenAiSessionConfigOutputMaxResponseOutputTokens {
	return &OpenAiSessionConfigOutputMaxResponseOutputTokens{typ: "InfStringLiteral", InfStringLiteral: "inf"}
}

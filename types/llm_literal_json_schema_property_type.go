package types

import (
	"encoding/json"
	"fmt"
)

type LlmLiteralJsonSchemaPropertyType struct {
	LlmLiteralJsonSchemaPropertyTypeZero LlmLiteralJsonSchemaPropertyTypeZero
	StringList                           []string
	typ                                  string
}

func (l *LlmLiteralJsonSchemaPropertyType) GetLlmLiteralJsonSchemaPropertyTypeZero() LlmLiteralJsonSchemaPropertyTypeZero {
	if l == nil {
		return ""
	}
	return l.LlmLiteralJsonSchemaPropertyTypeZero
}

func (l *LlmLiteralJsonSchemaPropertyType) GetStringList() []string {
	if l == nil {
		return nil
	}
	return l.StringList
}

func (l *LlmLiteralJsonSchemaPropertyType) UnmarshalJSON(data []byte) error {
	var valueLlmLiteralJsonSchemaPropertyTypeZero LlmLiteralJsonSchemaPropertyTypeZero
	if err := json.Unmarshal(data, &valueLlmLiteralJsonSchemaPropertyTypeZero); err == nil {
		l.typ = "LlmLiteralJsonSchemaPropertyTypeZero"
		l.LlmLiteralJsonSchemaPropertyTypeZero = valueLlmLiteralJsonSchemaPropertyTypeZero
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		l.typ = "StringList"
		l.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LlmLiteralJsonSchemaPropertyType) MarshalJSON() ([]byte, error) {
	if l.typ == "LlmLiteralJsonSchemaPropertyTypeZero" || l.LlmLiteralJsonSchemaPropertyTypeZero != "" {
		return json.Marshal(l.LlmLiteralJsonSchemaPropertyTypeZero)
	}
	if l.typ == "StringList" || l.StringList != nil {
		return json.Marshal(l.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LlmLiteralJsonSchemaPropertyTypeVisitor interface {
	VisitLlmLiteralJsonSchemaPropertyTypeZero(LlmLiteralJsonSchemaPropertyTypeZero) error
	VisitStringList([]string) error
}

func (l *LlmLiteralJsonSchemaPropertyType) Accept(visitor LlmLiteralJsonSchemaPropertyTypeVisitor) error {
	if l.typ == "LlmLiteralJsonSchemaPropertyTypeZero" || l.LlmLiteralJsonSchemaPropertyTypeZero != "" {
		return visitor.VisitLlmLiteralJsonSchemaPropertyTypeZero(l.LlmLiteralJsonSchemaPropertyTypeZero)
	}
	if l.typ == "StringList" || l.StringList != nil {
		return visitor.VisitStringList(l.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}

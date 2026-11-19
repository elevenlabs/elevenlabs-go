package types

import (
	"encoding/json"
	"fmt"
)

type LiteralJsonSchemaPropertyType struct {
	LiteralJsonSchemaPropertyTypeZero LiteralJsonSchemaPropertyTypeZero
	StringList                        []string
	typ                               string
}

func (l *LiteralJsonSchemaPropertyType) GetLiteralJsonSchemaPropertyTypeZero() LiteralJsonSchemaPropertyTypeZero {
	if l == nil {
		return ""
	}
	return l.LiteralJsonSchemaPropertyTypeZero
}

func (l *LiteralJsonSchemaPropertyType) GetStringList() []string {
	if l == nil {
		return nil
	}
	return l.StringList
}

func (l *LiteralJsonSchemaPropertyType) UnmarshalJSON(data []byte) error {
	var valueLiteralJsonSchemaPropertyTypeZero LiteralJsonSchemaPropertyTypeZero
	if err := json.Unmarshal(data, &valueLiteralJsonSchemaPropertyTypeZero); err == nil {
		l.typ = "LiteralJsonSchemaPropertyTypeZero"
		l.LiteralJsonSchemaPropertyTypeZero = valueLiteralJsonSchemaPropertyTypeZero
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

func (l LiteralJsonSchemaPropertyType) MarshalJSON() ([]byte, error) {
	if l.typ == "LiteralJsonSchemaPropertyTypeZero" || l.LiteralJsonSchemaPropertyTypeZero != "" {
		return json.Marshal(l.LiteralJsonSchemaPropertyTypeZero)
	}
	if l.typ == "StringList" || l.StringList != nil {
		return json.Marshal(l.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LiteralJsonSchemaPropertyTypeVisitor interface {
	VisitLiteralJsonSchemaPropertyTypeZero(LiteralJsonSchemaPropertyTypeZero) error
	VisitStringList([]string) error
}

func (l *LiteralJsonSchemaPropertyType) Accept(visitor LiteralJsonSchemaPropertyTypeVisitor) error {
	if l.typ == "LiteralJsonSchemaPropertyTypeZero" || l.LiteralJsonSchemaPropertyTypeZero != "" {
		return visitor.VisitLiteralJsonSchemaPropertyTypeZero(l.LiteralJsonSchemaPropertyTypeZero)
	}
	if l.typ == "StringList" || l.StringList != nil {
		return visitor.VisitStringList(l.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}

package types

import (
	"fmt"
)

type LiteralJsonSchemaPropertyTypeZero string

const (
	LiteralJsonSchemaPropertyTypeZeroBoolean LiteralJsonSchemaPropertyTypeZero = "boolean"
	LiteralJsonSchemaPropertyTypeZeroString  LiteralJsonSchemaPropertyTypeZero = "string"
	LiteralJsonSchemaPropertyTypeZeroInteger LiteralJsonSchemaPropertyTypeZero = "integer"
	LiteralJsonSchemaPropertyTypeZeroNumber  LiteralJsonSchemaPropertyTypeZero = "number"
)

func NewLiteralJsonSchemaPropertyTypeZeroFromString(s string) (LiteralJsonSchemaPropertyTypeZero, error) {
	switch s {
	case "boolean":
		return LiteralJsonSchemaPropertyTypeZeroBoolean, nil
	case "string":
		return LiteralJsonSchemaPropertyTypeZeroString, nil
	case "integer":
		return LiteralJsonSchemaPropertyTypeZeroInteger, nil
	case "number":
		return LiteralJsonSchemaPropertyTypeZeroNumber, nil
	}
	var t LiteralJsonSchemaPropertyTypeZero
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LiteralJsonSchemaPropertyTypeZero) Ptr() *LiteralJsonSchemaPropertyTypeZero {
	return &e
}

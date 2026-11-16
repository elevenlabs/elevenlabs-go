package types

import (
	"fmt"
)

type LlmLiteralJsonSchemaPropertyTypeZero string

const (
	LlmLiteralJsonSchemaPropertyTypeZeroBoolean LlmLiteralJsonSchemaPropertyTypeZero = "boolean"
	LlmLiteralJsonSchemaPropertyTypeZeroString  LlmLiteralJsonSchemaPropertyTypeZero = "string"
	LlmLiteralJsonSchemaPropertyTypeZeroInteger LlmLiteralJsonSchemaPropertyTypeZero = "integer"
	LlmLiteralJsonSchemaPropertyTypeZeroNumber  LlmLiteralJsonSchemaPropertyTypeZero = "number"
)

func NewLlmLiteralJsonSchemaPropertyTypeZeroFromString(s string) (LlmLiteralJsonSchemaPropertyTypeZero, error) {
	switch s {
	case "boolean":
		return LlmLiteralJsonSchemaPropertyTypeZeroBoolean, nil
	case "string":
		return LlmLiteralJsonSchemaPropertyTypeZeroString, nil
	case "integer":
		return LlmLiteralJsonSchemaPropertyTypeZeroInteger, nil
	case "number":
		return LlmLiteralJsonSchemaPropertyTypeZeroNumber, nil
	}
	var t LlmLiteralJsonSchemaPropertyTypeZero
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LlmLiteralJsonSchemaPropertyTypeZero) Ptr() *LlmLiteralJsonSchemaPropertyTypeZero {
	return &e
}

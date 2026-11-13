package types

import (
	"fmt"
)

type ObjectJsonSchemaPropertyInputPropertyKind string

const (
	ObjectJsonSchemaPropertyInputPropertyKindArray  ObjectJsonSchemaPropertyInputPropertyKind = "array"
	ObjectJsonSchemaPropertyInputPropertyKindObject ObjectJsonSchemaPropertyInputPropertyKind = "object"
)

func NewObjectJsonSchemaPropertyInputPropertyKindFromString(s string) (ObjectJsonSchemaPropertyInputPropertyKind, error) {
	switch s {
	case "array":
		return ObjectJsonSchemaPropertyInputPropertyKindArray, nil
	case "object":
		return ObjectJsonSchemaPropertyInputPropertyKindObject, nil
	}
	var t ObjectJsonSchemaPropertyInputPropertyKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ObjectJsonSchemaPropertyInputPropertyKind) Ptr() *ObjectJsonSchemaPropertyInputPropertyKind {
	return &e
}

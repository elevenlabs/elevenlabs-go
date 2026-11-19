package types

import (
	"fmt"
)

type DataCollectionValueSchemaTypeZero string

const (
	DataCollectionValueSchemaTypeZeroBoolean DataCollectionValueSchemaTypeZero = "boolean"
	DataCollectionValueSchemaTypeZeroString  DataCollectionValueSchemaTypeZero = "string"
	DataCollectionValueSchemaTypeZeroInteger DataCollectionValueSchemaTypeZero = "integer"
	DataCollectionValueSchemaTypeZeroNumber  DataCollectionValueSchemaTypeZero = "number"
)

func NewDataCollectionValueSchemaTypeZeroFromString(s string) (DataCollectionValueSchemaTypeZero, error) {
	switch s {
	case "boolean":
		return DataCollectionValueSchemaTypeZeroBoolean, nil
	case "string":
		return DataCollectionValueSchemaTypeZeroString, nil
	case "integer":
		return DataCollectionValueSchemaTypeZeroInteger, nil
	case "number":
		return DataCollectionValueSchemaTypeZeroNumber, nil
	}
	var t DataCollectionValueSchemaTypeZero
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DataCollectionValueSchemaTypeZero) Ptr() *DataCollectionValueSchemaTypeZero {
	return &e
}

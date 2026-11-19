package types

import (
	"encoding/json"
	"fmt"
)

type DataCollectionValueSchemaType struct {
	DataCollectionValueSchemaTypeZero DataCollectionValueSchemaTypeZero
	StringList                        []string
	typ                               string
}

func (d *DataCollectionValueSchemaType) GetDataCollectionValueSchemaTypeZero() DataCollectionValueSchemaTypeZero {
	if d == nil {
		return ""
	}
	return d.DataCollectionValueSchemaTypeZero
}

func (d *DataCollectionValueSchemaType) GetStringList() []string {
	if d == nil {
		return nil
	}
	return d.StringList
}

func (d *DataCollectionValueSchemaType) UnmarshalJSON(data []byte) error {
	var valueDataCollectionValueSchemaTypeZero DataCollectionValueSchemaTypeZero
	if err := json.Unmarshal(data, &valueDataCollectionValueSchemaTypeZero); err == nil {
		d.typ = "DataCollectionValueSchemaTypeZero"
		d.DataCollectionValueSchemaTypeZero = valueDataCollectionValueSchemaTypeZero
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		d.typ = "StringList"
		d.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DataCollectionValueSchemaType) MarshalJSON() ([]byte, error) {
	if d.typ == "DataCollectionValueSchemaTypeZero" || d.DataCollectionValueSchemaTypeZero != "" {
		return json.Marshal(d.DataCollectionValueSchemaTypeZero)
	}
	if d.typ == "StringList" || d.StringList != nil {
		return json.Marshal(d.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", d)
}

type DataCollectionValueSchemaTypeVisitor interface {
	VisitDataCollectionValueSchemaTypeZero(DataCollectionValueSchemaTypeZero) error
	VisitStringList([]string) error
}

func (d *DataCollectionValueSchemaType) Accept(visitor DataCollectionValueSchemaTypeVisitor) error {
	if d.typ == "DataCollectionValueSchemaTypeZero" || d.DataCollectionValueSchemaTypeZero != "" {
		return visitor.VisitDataCollectionValueSchemaTypeZero(d.DataCollectionValueSchemaTypeZero)
	}
	if d.typ == "StringList" || d.StringList != nil {
		return visitor.VisitStringList(d.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", d)
}

package types

import (
	"encoding/json"
	"fmt"
)

type EnvironmentVariableResponseValues struct {
	StringUnknownMap map[string]string
	typ              string
}

func (e *EnvironmentVariableResponseValues) GetStringUnknownMap() map[string]string {
	if e == nil {
		return nil
	}
	return e.StringUnknownMap
}

func (e *EnvironmentVariableResponseValues) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]string
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		e.typ = "StringUnknownMap"
		e.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, e)
}

func (e EnvironmentVariableResponseValues) MarshalJSON() ([]byte, error) {
	if e.typ == "StringUnknownMap" || e.StringUnknownMap != nil {
		return json.Marshal(e.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EnvironmentVariableResponseValuesVisitor interface {
	VisitStringUnknownMap(map[string]string) error
}

func (e *EnvironmentVariableResponseValues) Accept(visitor EnvironmentVariableResponseValuesVisitor) error {
	if e.typ == "StringUnknownMap" || e.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(e.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", e)
}

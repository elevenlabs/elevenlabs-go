package types

import (
	"encoding/json"
	"fmt"
)

// Labels for the voice. Keys can be language, accent, gender, or age.
type UpdateVoicesRequestLabels struct {
	StringUnknownMap map[string]string
	String           string
	typ              string
}

func (u *UpdateVoicesRequestLabels) GetStringUnknownMap() map[string]string {
	if u == nil {
		return nil
	}
	return u.StringUnknownMap
}

func (u *UpdateVoicesRequestLabels) GetString() string {
	if u == nil {
		return ""
	}
	return u.String
}

func (u *UpdateVoicesRequestLabels) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]string
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		u.typ = "StringUnknownMap"
		u.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		u.typ = "String"
		u.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UpdateVoicesRequestLabels) MarshalJSON() ([]byte, error) {
	if u.typ == "StringUnknownMap" || u.StringUnknownMap != nil {
		return json.Marshal(u.StringUnknownMap)
	}
	if u.typ == "String" || u.String != "" {
		return json.Marshal(u.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateVoicesRequestLabelsVisitor interface {
	VisitStringUnknownMap(map[string]string) error
	VisitString(string) error
}

func (u *UpdateVoicesRequestLabels) Accept(visitor UpdateVoicesRequestLabelsVisitor) error {
	if u.typ == "StringUnknownMap" || u.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(u.StringUnknownMap)
	}
	if u.typ == "String" || u.String != "" {
		return visitor.VisitString(u.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}
